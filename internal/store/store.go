package store

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"taskman/internal/config"
	"taskman/internal/types"
	db "taskman/pkg/db/sqlite"
)

var (
	ErrNotFound          = errors.New("record not found")
	ErrInvalidArgument   = errors.New("invalid argument")
	ErrSessionInProgress = errors.New("session already in progress")
	ErrSessionNotFound   = errors.New("session not found")
	ErrSessionFailed     = errors.New("session failed")

	ClearSession int64 = 0
)

type ProjectStoreInterface interface {
	Add(project types.Project) (*types.Project, error)
	Modify(projectId *string, project types.Project) (*types.Project, error)
	Delete(projectId string) error
	GetProject(projectId string) (*types.Project, error)
}

type TaskStoreInterface interface {
	AddTask(task types.Task) (*types.Task, error)
	ModifyTask(task db.Task) (*types.Task, error)
	RemoveTask(taskId int64) error
	GetAllTasks() ([]types.Task, error)
	GetTask(taskId int64) (*types.Task, error)
}

type SessionStoreInterface interface {
	GetSession(sessionId int64) (*types.Session, error)
	GetTaskSessions(taskId int64) ([]types.Session, error)
	Start(task *types.Task) (*types.Session, error)
	Stop(sessionId int64) error
	Modify(session *types.Session) (*types.Session, error)
	Delete(sessionId int64) error
}

type Store struct {
	ctx context.Context
	db  *sql.DB
	cfg *config.Config

	TaskStore    TaskStoreInterface
	SessionStore SessionStoreInterface
	ProjectStore ProjectStoreInterface
}

func NewStore(ctx context.Context, cfg *config.Config, dbase *sql.DB) *Store {
	taskStore := NewTaskStore(ctx, dbase)
	projectStore := NewProjectStore(ctx, dbase)
	sessionStore := NewSessionStore(ctx, dbase)

	return &Store{
		ctx:          ctx,
		cfg:          cfg,
		db:           dbase,
		TaskStore:    taskStore,
		ProjectStore: projectStore,
		SessionStore: sessionStore,
	}
}
