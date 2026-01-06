package task

type Status string

const (
	Completed  Status = "completed"
	InProgress Status = "in-progress"
	OnHold     Status = "on-hold"
)
