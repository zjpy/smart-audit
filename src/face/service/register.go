package service

import (
	"audit/contract"
	"errors"
)

type FaceRegistration struct {
}

func (f *FaceRegistration) Register(args []string) (contract.ServiceRuleID, error) {
	if len(args) < 1 {
		return 0, errors.New("注册规则所需参数不足")
	}

	return f.serviceRegister(args[0])
}

func (f *FaceRegistration) serviceRegister(
	expression string) (contract.ServiceRuleID, error) {
	// fixme 实际商用时实现人脸识别预言机，然后在这里调用预言机服务
	return 0, nil
}
