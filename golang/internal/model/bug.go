package model

type Severity int

const (
	SEVERITY_P0 Severity = iota
	SEVERITY_P1
	SEVERITY_P2
)

type Bug struct {
	Task
	Severity Severity
	Status   TaskStatus
}

func (b Bug) CreateTask(task *Task, taskType TaskType) (*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (b Bug) UpdateTaskStatus(task *Task, status TaskStatus) (*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (b Bug) UpdateTaskAssignee(task *Task, assignee string) (*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (b Bug) TasksAssignedToUser(assignee string) ([]*Task, error) {
	//TODO implement me
	panic("implement me")
}
