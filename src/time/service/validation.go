package service

import "audit/contract"

type TimeValidation struct {
}

func (t *TimeValidation) Validate(id contract.ServiceRuleID, args []string) error {
	rules, err := contract.ServiceRulesFromArgs(args)
	if err != nil {
		return err
	}

	return t.serviceValidate(id, rules)
}

func (t *TimeValidation) serviceValidate(id contract.ServiceRuleID,
	rules *contract.ServiceRules) error {
	// fixme 实际商用时实现时间预言机，然后在这里调用预言机服务
	return nil
}
