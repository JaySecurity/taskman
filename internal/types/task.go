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
	priority := string(t.Priority)
	return db.Task{
		ID:        t.ID,
		Name:      t.Name,
		Project:   t.Project,
		Priority:  &priority,
		Notes:     t.Notes,
		DueDate:   t.DueDate,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func TaskFromDB(task db.Task) (*Task, error) {
	return &Task{
		ID:        task.ID,
		Name:      task.Name,
		Project:   task.Project,
		Client:    task.Client,
		Priority:  Priority(*task.Priority),
		Notes:     task.Notes,
		DueDate:   task.DueDate,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}
