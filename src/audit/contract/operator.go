package contract

type LogicOperator string
type ConditionalOperator string

// 逻辑关系表达相关操作符
const (
	// 空规则，仅用于单个时
	NONE LogicOperator = ""

	// 并且
	AND LogicOperator = "&&"

	// 或
	OR LogicOperator = "||"

	// 非
	NOT LogicOperator = "!"
)

// 条件表达式相关操作符
const (
	// 大于
	GT ConditionalOperator = ">"

	// 大于等于
	GE ConditionalOperator = ">="

	// 小于
	LT ConditionalOperator = "<"

	// 小于等于
	LE ConditionalOperator = "<="

	// 等于
	EQ ConditionalOperator = "=="

	// 在…范围内
	IN ConditionalOperator = "in"
)
