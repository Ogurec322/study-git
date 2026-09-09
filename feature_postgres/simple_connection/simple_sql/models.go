package simple_sql

import "time"

type TaskModel struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAT   time.Time
	CompletedAt *time.Time
}
