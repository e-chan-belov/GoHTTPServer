package service

import (
	"hw2/domain"
	"hw2/repository"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	repo *repository.Task
}

func NewTask(repo *repository.Task) *Task {
	return &Task{
		repo: repo,
	}
}

func (s *Task) PostTask(name domain.TaskName) (domain.TaskId, error) {
	task := domain.Task{Id: domain.TaskId(uuid.New().String()), TaskName: name, Status: "not_finished", Result: ""}
	if err := (*s.repo).Post(task); err != nil {
		return "0", err
	}
	go func() {
		time.Sleep(time.Second * 30) // compiler needs some time...
		task = domain.Task{Id: task.Id, TaskName: name, Status: "finished", Result: "ABC"}
		if err := (*s.repo).Put(task); err != nil {
			return
		}
	}()
	return task.Id, nil
}

func (s *Task) GetStatus(id domain.TaskId) (domain.TaskStatus, error) {
	value, err := (*s.repo).Get(id)
	return value.Status, err
}

func (s *Task) GetResult(id domain.TaskId) (domain.TaskResult, error) {
	value, err := (*s.repo).Get(id)
	return value.Result, err
}
