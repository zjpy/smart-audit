package contract

// 用于定义单个服务器规则
type RuleItem struct {
	Logic     LogicOperator
	Condition ConditionalOperator
	Params    []string
}

type ServiceRules struct {
	Rules []RuleItem
}

func ServiceRulesFromArgs(args []string) (*ServiceRules, error) {
	// todo complete me
	return nil, nil
}
