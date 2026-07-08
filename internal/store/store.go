package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"taskman/internal/config"
	"taskman/internal/types"
)

type TaskStoreInterface interface {
	AddTask(ctx context.Context, task types.Task) error
	ModifyTask(ctx context.Context, task types.Task) error
	RemoveTask(ctx context.Context, taskId int64) error
	GetAllTasks() ([]types.Task, error)
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
	fmt.Println("Database Connected")
	// db.Conn(context.Background())

	taskStore := NewTaskStore(ctx, dbase)

	return &Store{
		ctx:       ctx,
		cfg:       cfg,
		db:        dbase,
		TaskStore: taskStore,
	}
}
