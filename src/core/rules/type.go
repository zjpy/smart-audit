package rules

type RuleType int16

const (
	None RuleType = iota
	Time
	Location
	FaceRecognize
	ObjectRecognize
)

func (t RuleType) ContractName() string {
	switch t {
	case Time:
		return "Time"
	case Location:
		return "Location"
	case FaceRecognize:
		return "FaceRecognize"
	case ObjectRecognize:
		return "ObjectRecognize"
	default:
		return "Unknown"
	}
}
