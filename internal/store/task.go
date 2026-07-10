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

func (s *TaskStore) AddTask(task types.Task) (*types.Task, error) {
	err := s.queries.AddProject(*s.ctx, task.Project)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	priority := string(task.Priority)
	data, err := s.queries.CreateTask(*s.ctx, db.CreateTaskParams{
		Name:     task.Name,
		Project:  task.Project,
		Priority: &priority,
	})
	if err != nil {
		return nil, err
	}
	newtask, err := types.TaskFromDB(data)
	if err != nil {
		return nil, err
	}
	return newtask, nil
}

func (s *TaskStore) ModifyTask(task db.Task) (*types.Task, error) {
	params := db.ModifyTaskParams{
		ID:       task.ID,
		Name:     task.Name,
		Client:   task.Client,
		Project:  task.Project,
		Priority: task.Priority,
		DueDate:  task.DueDate,
	}
	updatedTask, err := s.queries.ModifyTask(*s.ctx, params)
	if err != nil {
		return nil, err
	}
	result, err := types.TaskFromDB(updatedTask)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TaskStore) RemoveTask(taskId int64) error {
	err := s.queries.DeleteTask(*s.ctx, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskStore) GetAllTasks() ([]types.Task, error) {
	tasks, err := s.queries.GetAllTasks(*s.ctx)
	if err != nil {
		fmt.Printf("Error fetching tasks: %v", err)
		return nil, err
	}
	var result []types.Task
	for _, task := range tasks {
		t, err := types.TaskFromDB(task)
		if err != nil {
			return nil, err
		}
		result = append(result, *t)
	}
	return result, nil
}

func (s *TaskStore) GetTask(taskId int64) (*types.Task, error) {
	task, err := s.queries.GetTask(*s.ctx, taskId)
	if err != nil {
		return nil, err
	}
	result, err := types.TaskFromDB(task)
	if err != nil {
		return nil, err
	}
	return result, nil
}
