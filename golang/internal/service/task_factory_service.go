package service

import (
	"machine-coding-basic-jira/internal/model"
	"time"
)

type TaskFactory interface {
	CreateTask(title string, assignee string, dueDate time.Time) (model.ITask, error)
	//UpdateTaskAssignee(task model.ITask, assignee string) (model.ITask, error)
	// UpdateTaskStatus(task model.ITask, status model.TaskStatus) (model.ITask, error)
	//TasksAssignedToUser(assignee string) ([]*model.Task, error)
}
