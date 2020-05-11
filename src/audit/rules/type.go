package rules

type RuleType int16

const (
	None RuleType = iota
	Time
	Location
	FaceRecognize
	ObjectRecognize
	Custom
)
