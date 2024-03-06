package service

import (
	"fmt"
	"machine-coding-basic-jira/internal/model"
	"time"
)

type BugTaskFactory struct{}

func (b *BugTaskFactory) CreateTask(title string, assignee string, dueDate time.Time) (model.ITask, error) {
	// can have validators return error
	fmt.Println("Enter Severity: (0 for P0, 1 for P1, 2 for P2")
	var severityInput int
	_, err := fmt.Scanln(&severityInput)

	if err != nil || severityInput < 0 || severityInput > 2 {
		fmt.Println("Invalid input for severityInput value.")
		return nil, err
	}
	return &model.Bug{
		Task: model.Task{
			Title:    title,
			Assignee: assignee,
			DueDate:  dueDate,
			Type:     model.TASK_TYPE_BUG,
		},
		Severity: model.Severity(severityInput),
		Status:   model.TASK_STATUS_OPEN,
	}, nil
}

func (b *BugTaskFactory) UpdateTaskAssignee(task model.ITask, assignee string) (model.ITask, error) {
	task.SetAssignee(assignee)
	return task, nil
}
