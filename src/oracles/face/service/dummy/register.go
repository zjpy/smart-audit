package dummy

import "core/contract"

type FaceRegistration struct {
}

func (f *FaceRegistration) Register(args []string) (contract.ServiceRuleID, error) {
	// 由于假设是固定的验证规则，所以这里不需要额外工作
	return 0, nil
}
