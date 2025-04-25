package repository

import (
	"hw2/domain"
)

type Task interface {
	Post(task domain.Task) error
	Put(task domain.Task) error
	Get(id domain.TaskId) (domain.Task, error)
}
