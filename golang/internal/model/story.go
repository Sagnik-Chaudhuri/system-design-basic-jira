package model

type Story struct {
	Task
	StorySummary string
	Status       TaskStatus
	SubTracks    []*SubTrack
}
