package model

import "time"

type AllowedStatusSprint int

const (
	ALLOWED_STATUS_SPRINT_OPEN AllowedStatusSprint = iota
	ALLOWED_STATUS_SPRINT_IN_PROGRESS
	ALLOWED_STATUS_SPRINT_COMPLETED
)

type Sprint struct {
	Tasks         []*Task
	Name          string
	StartDate     time.Time
	EndDate       time.Time
	AllowedStatus AllowedStatusSprint
}
