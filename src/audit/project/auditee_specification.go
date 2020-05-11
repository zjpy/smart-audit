package project

import (
	"audit/orgnization"
	"audit/rules"
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
