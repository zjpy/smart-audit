package project

import (
	"audit/orgnization"
	"audit/rules"
	"github.com/hyperledger/fabric/core/chaincode/shim"
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

func (a *AuditeeSpecification) Validate(stub shim.ChaincodeStubInterface) error {
	// todo complete me
	return nil
}

func (a *AuditeeSpecification) Key() string {
	return string(a.ID)
}

func (a *AuditeeSpecification) Value() []byte {
	// todo 这里不需要序列化所有的值，只是将Auditee, Project, Rule的Key值序列化即可
	return nil
}
