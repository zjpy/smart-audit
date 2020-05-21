package project

import (
	"core/common"
	"core/orgnization"
	rules2 "core/rules"
	"encoding/json"
	"errors"
)

// 审计当事人规范结构，主要用于定义审计当事人与规则的关联
type AuditeeSpecification struct {
	// 规范ID，前12byte为审计当事人ID+项目ID+规则ID，多余字节可支持之后扩展
	ID common.Uint256

	// 审计当事人的相关信息
	Auditee orgnization.Auditee

	// 该规范所属的业务
	Project Project

	// 审计当事人需要遵守的规则
	Rule rules2.ValidationRelationship
}

func (a *AuditeeSpecification) Validate() error {
	// todo complete me
	return nil
}

func (a *AuditeeSpecification) Key() string {
	return a.ID.String()
}

func (a *AuditeeSpecification) Value() ([]byte, error) {
	value, err := json.Marshal(a)
	if err != nil {
		return nil, errors.New("failed to marshal auditee specification")
	}

	return value, nil
}
