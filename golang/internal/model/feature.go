package model

type Impact int

const (
	IMPACT_LOW Impact = iota
	IMPACT_MODERATE
	IMPACT_HIGH
)

type Feature struct {
	Task           // Embedding Task struct
	FeatureSummary string
	Impact         Impact
	Status         TaskStatus
}
