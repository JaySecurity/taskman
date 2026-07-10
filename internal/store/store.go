package store

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"taskman/internal/config"
	"taskman/internal/types"
	db "taskman/pkg/db/sqlite"
)

type TaskStoreInterface interface {
	AddTask(task types.Task) (*types.Task, error)
	ModifyTask(task db.Task) (*types.Task, error)
	RemoveTask(taskId int64) error
	GetAllTasks() ([]types.Task, error)
	GetTask(taskId int64) (*types.Task, error)
}

type SessionStoreInterface interface {
	Add()
	Modify()
	Delete()
}

type Store struct {
	ctx *context.Context
	db  *sql.DB
	cfg *config.Config

	TaskStore    TaskStoreInterface
	SessionStore SessionStoreInterface
}

func NewStore(ctx *context.Context, cfg *config.Config) *Store {
	dbase, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	taskStore := NewTaskStore(ctx, dbase)

	return &Store{
		ctx:       ctx,
		cfg:       cfg,
		db:        dbase,
		TaskStore: taskStore,
	}
}
