package store

import (
	"context"
	"database/sql"
	"fmt"

	"taskman/internal/types"
	db "taskman/pkg/db/sqlite"
)

type TaskStore struct {
	ctx      *context.Context
	database *sql.DB
	queries  *db.Queries
}

func NewTaskStore(ctx *context.Context, database *sql.DB) *TaskStore {
	queries := db.New(database)
	return &TaskStore{
		ctx:      ctx,
		database: database,
		queries:  queries,
	}
}

func (s *TaskStore) AddTask(ctx context.Context, task types.Task) error {
	// Check for project and if task has a project lookup / add it
	// Create task and return
	return nil
}

func (s *TaskStore) ModifyTask(ctx context.Context, task db.Task) error {
	return nil
}

func (s *TaskStore) RemoveTask(ctx context.Context, taskId int64) error {
	return nil
}

func (s *TaskStore) GetAllTasks() ([]types.Task, error) {
	ctx := context.Background()
	_, err := s.queries.GetAllTasks(ctx)
	if err != nil {
		fmt.Printf("Error fetching tasks: %v", err)
		return nil, err
	}
	fmt.Println("Test")
	return nil, nil
}
