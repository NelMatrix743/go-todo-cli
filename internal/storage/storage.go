package storage

import "github.com/NelMatrix743/go-todo-cli/internal/task"

type StorageInteface interface {
	CreateTask(task.Task) (task.Task, error)
	GetAllTask() ([]task.Task, error)
	GetTaskByID(int) (task.Task, error)
	UpdateTask(task.Task) error
	DeleteTask(task.Task) error
}
