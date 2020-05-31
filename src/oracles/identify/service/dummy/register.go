package dummy

import "core/contract"

type EntityIdentifyRegistration struct {
}

// 模拟注册物体识别规则
func (f *EntityIdentifyRegistration) Register(args []string) (contract.ServiceRuleID, error) {
	// 由于假设是固定的验证规则，所以这里不需要额外工作
	return 0, nil
}
