package ram_storage

import (
	"hw2/domain"
	"hw2/repository"
)

type TaskLine struct {
	TaskName domain.TaskName
	Status   domain.TaskStatus
	Result   domain.TaskResult
}

type Task struct {
	data map[domain.TaskId]TaskLine
}

func taskToLine(task domain.Task) *TaskLine {
	return &TaskLine{TaskName: task.TaskName, Status: task.Status, Result: task.Result}
}

func lineToTask(taskLine TaskLine) domain.Task {
	return domain.Task{TaskName: taskLine.TaskName, Status: taskLine.Status, Result: taskLine.Result}
}

func NewTask() *Task {
	return &Task{
		data: make(map[domain.TaskId]TaskLine),
	}
}

func (s *Task) Post(task domain.Task) error {
	if _, exists := s.data[task.Id]; exists {
		return repository.AlreadyExists
	}
	s.data[task.Id] = *taskToLine(task)
	return nil
}

func (s *Task) Put(task domain.Task) error {
	s.data[task.Id] = *taskToLine(task)
	return nil
}

func (s *Task) Get(id domain.TaskId) (domain.Task, error) {
	line, exists := s.data[id]
	if !exists {
		return domain.Task{}, repository.NotFound
	}
	return lineToTask(line), nil
}
