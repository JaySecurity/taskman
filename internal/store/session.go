package store

import (
	"context"
	"database/sql"
	"time"

	"taskman/internal/types"
	db "taskman/pkg/db/sqlite"
)

type SessionStore struct {
	ctx     context.Context
	dbase   *sql.DB
	queries *db.Queries
}

func SessionFromSQLite(session db.Session) *types.Session {
	return &types.Session{
		ID:     session.ID,
		TaskID: session.TaskID,
		Start:  session.Start,
		Stop:   session.Stop,
	}
}

func NewSessionStore(ctx context.Context, database *sql.DB) *SessionStore {
	queries := db.New(database)
	return &SessionStore{
		ctx:     ctx,
		dbase:   database,
		queries: queries,
	}
}

func (s *SessionStore) GetSession(sessionId int64) (*types.Session, error) {
	session, err := s.queries.GetSession(s.ctx, sessionId)
	if err != nil {
		return nil, err
	}
	return SessionFromSQLite(session), nil
}

func (s *SessionStore) GetTaskSessions(taskId int64) ([]types.Session, error) {
	result, err := s.queries.GetTaskSessions(s.ctx, taskId)
	if err != nil {
		return nil, err
	}
	var data []types.Session
	for _, session := range result {
		data = append(data, *SessionFromSQLite(session))
	}
	return data, nil
}

func (s *SessionStore) Start(task *types.Task) (*types.Session, error) {
	if task.CurrentSession != 0 {
		session, err := s.GetSession(task.CurrentSession)
		if err != nil {
			return nil, ErrSessionNotFound
		}
		return session, ErrSessionInProgress
	}
	startTime := time.Now()

	params := db.StartParams{
		TaskID: task.ID,
		Start:  &startTime,
	}
	data, err := s.queries.Start(s.ctx, params)
	if err != nil {
		return nil, ErrSessionFailed
	}
	session := SessionFromSQLite(data)
	taskParams := db.ModifyTaskParams{
		ID:             task.ID,
		CurrentSession: &session.ID,
	}

	_, err = s.queries.ModifyTask(s.ctx, taskParams)
	if err != nil {
		return nil, ErrSessionFailed
	}
	return session, nil
}

func (s *SessionStore) Stop(sessionId int64) error {
	session, err := s.GetSession(sessionId)
	if err != nil {
		return ErrSessionNotFound
	}
	task, err := s.queries.GetTask(s.ctx, session.TaskID)
	if err != nil {
		return ErrSessionFailed
	}
	if *task.CurrentSession == 0 {
		return ErrSessionFailed
	}
	stopTime := time.Now()

	params := db.StopParams{
		ID:   sessionId,
		Stop: &stopTime,
	}
	data, err := s.queries.Stop(s.ctx, params)
	if err != nil {
		return ErrSessionFailed
	}
	session = SessionFromSQLite(data)
	taskParams := db.ModifyTaskParams{
		ID:             task.ID,
		CurrentSession: &ClearSession,
	}

	_, err = s.queries.ModifyTask(s.ctx, taskParams)
	if err != nil {
		return ErrSessionFailed
	}
	return nil
}

func (s *SessionStore) Modify(session *types.Session) (*types.Session, error) {
	data, err := s.queries.ModifySession(s.ctx, db.ModifySessionParams{
		ID:    session.ID,
		Start: session.Start,
		Stop:  session.Stop,
	})
	if err != nil {
		return nil, err
	}
	return SessionFromSQLite(data), nil
}

func (s *SessionStore) Delete(sessionId int64) error {
	err := s.queries.DeleteSession(s.ctx, sessionId)
	if err != nil {
		return err
	}
	return nil
}
