package usecases

import (
	"hw2/domain"
)

type Task interface {
	PostTask(name domain.TaskName) (domain.TaskId, error)
	GetStatus(id domain.TaskId) (domain.TaskStatus, error)
	GetResult(id domain.TaskId) (domain.TaskResult, error)
}
