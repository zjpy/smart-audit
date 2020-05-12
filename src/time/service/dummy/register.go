package dummy

import "audit/contract"

type TimeRegistration struct {
}

func (t *TimeRegistration) Register(id contract.ServiceRuleID, args []string) error {
	// 由于假设是固定的验证规则，所以这里不需要额外工作
	return nil
}

