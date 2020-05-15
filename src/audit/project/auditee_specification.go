package project

import (
	"audit/common"
	"audit/orgnization"
	"audit/rules"
	"bytes"
	"errors"
)

// todo 在fabric中找到对应的类型
type Uint256 []byte

// 审计当事人规范结构，主要用于定义审计当事人与规则的关联
type AuditeeSpecification struct {
	// 规范ID
	ID Uint256

	// 审计当事人的相关信息
	Auditee orgnization.Auditee

	// 该规范所属的业务
	Project Project

	// 审计当事人需要遵守的规则
	Rule rules.Rule
}

func (a *AuditeeSpecification) Validate() error {
	// todo complete me
	return nil
}

func (a *AuditeeSpecification) Key() string {
	return string(a.ID)
}

func (a *AuditeeSpecification) Value() ([]byte, error) {
	// 这里不需要序列化所有的值，只是将Auditee, Project, Rule的Key值序列化即可
	w := new(bytes.Buffer)
	if err := common.WriteVarString(w, a.Auditee.Key()); err != nil {
		return nil, errors.New("failed to serialize key of Auditee")
	}
	if err := common.WriteVarString(w, a.Project.Key()); err != nil {
		return nil, errors.New("failed to serialize key of Project")
	}
	if err := common.WriteVarString(w, a.Rule.Key()); err != nil {
		return nil, errors.New("failed to serialize key of Auditee")
	}
	return w.Bytes(), nil
}
