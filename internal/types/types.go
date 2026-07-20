package types

import "time"

type Priority string

const (
	PriorityLow    Priority = "L"
	PriorityMedium Priority = "M"
	PriorityHigh   Priority = "H"
)

type Task struct {
	ID             int64      `json:"id"`
	Name           string     `json:"name"`
	Project        *string    `json:"project"`
	Client         *string    `json:"client"`
	Priority       Priority   `json:"priority"`
	Status         *string    `json:"status"`
	Notes          *string    `json:"notes"`
	DueDate        *time.Time `json:"due_date"`
	CurrentSession int64      `json:"current_session"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type Session struct {
	ID     int64      `json:"id"`
	TaskID int64      `json:"task_id"`
	Start  *time.Time `json:"start"`
	Stop   *time.Time `json:"stop"`
}

type Project struct {
	Name *string  `json:"name"`
	Rate *float64 `json:"rate"`
}
