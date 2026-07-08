package types

import (
	"time"

	db "taskman/pkg/db/sqlite"
)

type Task struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Project   *string    `json:"project"`
	Client    *string    `json:"client"`
	Priority  Priority   `json:"priority"`
	Notes     *string    `json:"notes"`
	DueDate   *time.Time `json:"due_date"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (t *Task) ToDB() db.Task {
	return db.Task{}
}

func TaskFromDB(db.Task) (Task, error) {
	return Task{}, nil
}
