package service

import (
	"fmt"
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
	FeatureTaskFactory  // embedding factory structs
	StoryTaskFactory
	BugTaskFactory
}

type TaskPlannerService interface {
	GetTaskFactoryFromTaskType(model.TaskType) TaskFactory
	CreateTask(title string, assignee string, dueDate time.Time, taskType model.TaskType) (model.ITask, error)
	UpdateAssignee(task model.ITask) (model.ITask, error)
	DisplayTasksAssignedToUser(assignee string)
}

func (t *TaskPlannerServiceImpl) GetTaskFactoryFromTaskType(taskType model.TaskType) TaskFactory {
	// returns specific task factory as per task type
	// returning address to the structs since methods of TaskFactory interface are implemented by pointers to these structs
	// ie, check method signatures:
	// func (f *FeatureTaskFactory) CreateTask(title string, assignee string, dueDate time.Time) (model.ITask, error)
	switch taskType {
	case model.TASK_TYPE_FEATURE:
		return &t.FeatureTaskFactory
	case model.TASK_TYPE_BUG:
		return &t.BugTaskFactory
	case model.TASK_TYPE_STORY:
		return &t.StoryTaskFactory
	}
	return nil
}

func (t *TaskPlannerServiceImpl) CreateTask(title string, assignee string, dueDate time.Time, taskType model.TaskType) (model.ITask, error) {
	taskFactory := t.GetTaskFactoryFromTaskType(taskType)
	task, err := taskFactory.CreateTask(title, assignee, dueDate)
	if err != nil {
		return nil, err
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

func (t *TaskPlannerServiceImpl) UpdateAssignee(task model.ITask) (model.ITask, error) {
	fmt.Println("Enter new assignee")
	var assignee string
	fmt.Scanln(&assignee)
	taskFactory := t.GetTaskFactoryFromTaskType(task.GetType())
	task, err := taskFactory.UpdateTaskAssignee(task, assignee)
	// handle error properly
	if err != nil {
		return nil, err
	}
	if taskGroup, ok := t.TasksGroupedByUsers[task.GetAssignee()]; ok {
		// Key exists, append to values
		t.TasksGroupedByUsers[task.GetAssignee()] = append(taskGroup, &task)
	} else {
		// Key does not exist, create new key-value pair
		t.TasksGroupedByUsers[task.GetAssignee()] = []*model.ITask{&task}
	}
	return task, nil
}

func (t *TaskPlannerServiceImpl) DisplayTasksAssignedToUser(assignee string) {
	if taskGroup, ok := t.TasksGroupedByUsers[assignee]; ok {
		for _, task := range taskGroup {
			// Ideally, should have  DisplayTask as part of ITask interface and other task factories should implement it
			log.Println("task: ", task)
		}
	} else {
		// Key does not exist
		log.Println("No tasks assigned to given user: ", assignee)
	}
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
