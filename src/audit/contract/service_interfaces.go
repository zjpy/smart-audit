package contract

const (
	ValidationFunctionName = "validation"
	RegisterFunctionName = "register"
)

// 用于定义服务中的规则ID
type ServiceRuleID uint32

type Validation interface {
	Validate(id ServiceRuleID, args []string) error
}

type Registration interface {
	Register(id ServiceRuleID, args []string) error
}
