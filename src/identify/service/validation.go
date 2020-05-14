package service

import "audit/contract"

// 调用预言机服务的人脸比对接口返回的结果
type EntityIdentifyResult struct {
	// 人脸比对结果，0表示请求成功，非0会有响应的错误码意义
	Result int `json:"result"`

	// 相似度评分，在0~1之间取值，越大表示越相似
	Score float32 `json:"score"`

	// 会话序号
	Sequence string `json:"sequence"`
}

type EntityIdentifyValidation struct {
}

func (f *EntityIdentifyValidation) Validate(id contract.ServiceRuleID, args []string) error {
	rules, err := contract.ServiceRulesFromArgs(args)
	if err != nil {
		return err
	}

	return f.serviceValidate(id, rules)
}

func (f *EntityIdentifyValidation) serviceValidate(id contract.ServiceRuleID,
	rules *contract.ServiceRules) error {
	// fixme 实际商用时实现时间预言机，然后在这里调用预言机服务
	return nil
}
