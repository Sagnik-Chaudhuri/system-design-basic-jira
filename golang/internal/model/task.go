package model

import "time"

type TaskType int

const (
	TASK_TYPE_FEATURE TaskType = iota
	TASK_TYPE_BUG
	TASK_TYPE_STORY
)

type TaskStatus int

const (
	TASK_STATUS_OPEN TaskStatus = iota
	TASK_STATUS_IN_PROGRESS
	TASK_STATUS_TESTING
	TASK_STATUS_DEPLOYED
	TASK_STATUS_FIXED
	TASK_STATUS_COMPLETED
)

type Task struct {
	Title    string
	Assignee string
	Type     TaskType
	DueDate  time.Time
}

type ITask interface {
	GetTitle() string
	GetAssignee() string
	GetType() TaskType
	GetDueDate() time.Time
}

func (t *Task) GetTitle() string {
	return t.Title
}
func (t *Task) GetAssignee() string {
	return t.Assignee
}
func (t *Task) GetType() TaskType {
	return t.Type
}
func (t *Task) GetDueDate() time.Time {
	return t.DueDate
}
