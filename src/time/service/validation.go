package service

import (
	"audit/contract"
	"errors"
)

type TimeValidation struct {
}

func (t *TimeValidation) Validate(id contract.ServiceRuleID, args []string) error {
	if len(args) < 1 {
		return errors.New("验证规则所需参数不足")
	}

	return t.serviceValidate(id, args[0])
}

func (t *TimeValidation) serviceValidate(id contract.ServiceRuleID,
	valueExpression string) error {
	// fixme 实际商用时实现时间预言机，然后在这里调用预言机服务
	return nil
}
