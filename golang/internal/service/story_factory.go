package service

import (
	"fmt"
	"machine-coding-basic-jira/internal/model"
	"time"
)

type StoryTaskFactory struct{}

func (s *StoryTaskFactory) CreateTask(title string, assignee string, dueDate time.Time) (model.ITask, error) {
	var summary string
	fmt.Println("Enter Story Summary")
	fmt.Scanln(&summary)

	return &model.Story{
		Task: model.Task{
			Title:    title,
			Assignee: assignee,
			DueDate:  dueDate,
			Type:     model.TASK_TYPE_STORY,
		},
		StorySummary: summary,
		Status:       model.TASK_STATUS_OPEN,
	}, nil
}

func (s *StoryTaskFactory) UpdateTaskAssignee(task model.ITask, assignee string) (model.ITask, error) {
	task.SetAssignee(assignee)
	return task, nil
}
