package handler

import (
	"machine-coding-basic-jira/internal/model"
	"machine-coding-basic-jira/internal/service"
	"time"
)

type TaskHandler struct {
	taskPlannerService service.TaskPlannerService
}

func GetTaskHandler() *TaskHandler {
	return &TaskHandler{
		taskPlannerService: service.GetTaskPlannerService(),
	}
}

func (t *TaskHandler) CreateTask(title string, assignee string, dueDate time.Time, taskType model.TaskType) (model.ITask, error) {
	return t.taskPlannerService.CreateTask(title, assignee, dueDate, taskType)
}
