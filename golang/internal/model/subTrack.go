package model

type AllowedStatusSubTrack int

const (
	ALLOWED_STATUS_SUBTRACK_OPEN AllowedStatusSubTrack = iota
	ALLOWED_STATUS_SUBTRACK_IN_PROGRESS
	ALLOWED_STATUS_SUBTRACK_COMPLETED
)

type SubTrack struct {
	Title         string
	AllowedStatus AllowedStatusSubTrack
	ParentTask    *Task
}
