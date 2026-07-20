package service

import (
	"taskman/internal/store"
	"taskman/internal/types"
)

type TaskService struct {
	store *store.Store
}

func NewTaskService(store *store.Store) *TaskService {
	return &TaskService{
		store: store,
	}
}

func (s *TaskService) GetAllTasks() ([]types.Task, error) {
	return s.store.TaskStore.GetAllTasks()
}
