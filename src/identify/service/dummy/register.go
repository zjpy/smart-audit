package dummy

import "audit/contract"

type EntityIdentifyRegistration struct {
}

func (f *EntityIdentifyRegistration) Register(args []string) (contract.ServiceRuleID, error) {
	// 由于假设是固定的验证规则，所以这里不需要额外工作
	return 0, nil
}
