package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"taskman/internal/config"
	"taskman/internal/service"
	"taskman/internal/store"
)

func main() {
	ctx := context.Background()
	cfg := config.Init()
	dbase, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		panic(err)
	}

	ds := store.NewStore(ctx, cfg, dbase)
	tasksrv := service.NewTaskService(ds)
	tasks, err := tasksrv.GetAllTasks()
	if err != nil {
		log.Fatalf("Error fetching tasks: %v", err)
	}

	for _, task := range tasks {
		if task.Client == nil {
			fmt.Printf(
				"%d: %s  Project: %s  Priority: %s\n",
				task.ID,
				task.Name,
				*task.Project,
				task.Priority,
			)
		} else {
			fmt.Printf(
				"%d: %s Client: %s   Project: %s  Priority: %s\n",
				task.ID,
				task.Name,
				*task.Client,
				*task.Project,
				task.Priority,
			)
		}
	}
}
