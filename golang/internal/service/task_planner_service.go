package service

import (
	"log"
	"machine-coding-basic-jira/internal/model"
	"sync"
	"time"
)

var allowedStatus = map[model.TaskType][]model.TaskStatus{
	model.TASK_TYPE_FEATURE: {model.TASK_STATUS_OPEN, model.TASK_STATUS_IN_PROGRESS, model.TASK_STATUS_TESTING, model.TASK_STATUS_DEPLOYED},
	model.TASK_TYPE_BUG:     {model.TASK_STATUS_OPEN, model.TASK_STATUS_IN_PROGRESS, model.TASK_STATUS_FIXED},
	model.TASK_TYPE_STORY:   {model.TASK_STATUS_OPEN, model.TASK_STATUS_IN_PROGRESS, model.TASK_STATUS_COMPLETED},
}

type TaskPlannerServiceImpl struct {
	Tasks               []*model.ITask
	TasksGroupedByUsers map[string][]*model.ITask
	FeatureTaskFactory
	StoryTaskFactory
	BugTaskFactory
}

type TaskPlannerService interface {
	CreateTask(title string, assignee string, dueDate time.Time, taskType model.TaskType) (model.ITask, error)
}

func (t *TaskPlannerServiceImpl) CreateTask(title string, assignee string, dueDate time.Time, taskType model.TaskType) (model.ITask, error) {
	var task model.ITask
	var err error
	switch taskType {
	case model.TASK_TYPE_FEATURE:
		task, err = t.FeatureTaskFactory.CreateTask(title, assignee, dueDate)
		if err != nil {
			return nil, err
		}
	case model.TASK_TYPE_BUG:
		task, err = t.BugTaskFactory.CreateTask(title, assignee, dueDate)
		if err != nil {
			return nil, err
		}
	case model.TASK_TYPE_STORY:
		task, err = t.StoryTaskFactory.CreateTask(title, assignee, dueDate)
		if err != nil {
			return nil, err
		}
	}
	t.Tasks = append(t.Tasks, &task)
	if taskGroup, ok := t.TasksGroupedByUsers[task.GetAssignee()]; ok {
		// Key exists, append to values
		t.TasksGroupedByUsers[task.GetAssignee()] = append(taskGroup, &task)
	} else {
		// Key does not exist, create new key-value pair
		t.TasksGroupedByUsers[task.GetAssignee()] = []*model.ITask{&task}
	}

	return task, nil
}

var taskPlannerServiceInstance TaskPlannerService
var taskPlannerOnce sync.Once

func GetTaskPlannerService() TaskPlannerService {
	log.Println("initialising GetTaskPlannerService")
	taskPlannerOnce.Do(func() {
		taskPlannerServiceInstance = &TaskPlannerServiceImpl{
			Tasks:               []*model.ITask{},
			TasksGroupedByUsers: map[string][]*model.ITask{},
			FeatureTaskFactory:  FeatureTaskFactory{},
			StoryTaskFactory:    StoryTaskFactory{},
			BugTaskFactory:      BugTaskFactory{},
		}
	})
	return taskPlannerServiceInstance
}
