package task

import "time"

type Task struct {
	Description string
	ID          int
	Status      Status
	createdAt   time.Time
}
