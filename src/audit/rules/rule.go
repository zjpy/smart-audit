package rules

// 用于定义一种规则模式的结构
type Rule struct {
	// 规则名称
	Name       string

	// 规则唯一标识
	ID         uint32

	// 具体规则定义
	Expression []string
}
