package main

import (
	"context"
	"fmt"

	"taskman/internal/config"
	"taskman/internal/store"
)

func main() {
	ctx := context.Background()
	cfg := config.Init()
	store := store.NewStore(&ctx, cfg)
	data, err := store.TaskStore.GetAllTasks()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
