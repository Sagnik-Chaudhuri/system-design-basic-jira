package service

import (
	"fmt"
	"machine-coding-basic-jira/internal/model"
	"time"
)

type FeatureTaskFactory struct{}

func (f *FeatureTaskFactory) CreateTask(title string, assignee string, dueDate time.Time) (model.ITask, error) {
	// can have validators return error
	var summary string
	var impactInput int
	fmt.Println("Enter Feature Summary")
	fmt.Scanln(&summary)
	fmt.Println("Enter impact (0 for Low, 1 for Moderate, 2 for High):")
	_, err := fmt.Scanln(&impactInput)

	if err != nil || impactInput < 0 || impactInput > 2 {
		fmt.Println("Invalid input or impact value.")
		return nil, err
	}
	return &model.Feature{
		Task: model.Task{
			Title:    title,
			Assignee: assignee,
			DueDate:  dueDate,
			Type:     model.TASK_TYPE_FEATURE,
		},
		FeatureSummary: summary,
		Impact:         model.Impact(impactInput),
		Status:         model.TASK_STATUS_OPEN,
	}, nil
}

func (f *FeatureTaskFactory) UpdateTaskAssignee(task model.ITask, assignee string) (model.ITask, error) {
	task.SetAssignee(assignee)
	return task, nil
}
