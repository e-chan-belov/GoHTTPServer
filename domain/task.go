package domain

type TaskId string
type TaskName string
type TaskStatus string
type TaskResult string

type Task struct {
	Id       TaskId     `json:"id"`
	TaskName TaskName   `json:"task"`
	Status   TaskStatus `json:"status"`
	Result   TaskResult `json:"result"`
}
