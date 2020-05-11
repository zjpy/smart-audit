package rules

type RuleType int16

const (
	Node RuleType = iota
	Time
	Location
	FaceRecognize
	ObjectRecognize
	Custom
)
