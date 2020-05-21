package rules

import (
	"bytes"
	"core/common"
	"core/contract"
	"fmt"
)

type ValidationExpression struct {
	// 规则类型
	Type RuleType

	// 具体验证规则
	Expression string
}

func RegisterRules(expression []string, context contract.Context) (uint32, error) {
	op, expressions, err := Parse(expression)
	if err != nil {
		return 0, err
	}

	relation := &ValidationRelationship{
		Operator: op,
	}
	for _, v := range expressions {
		ruleID, err := v.registerRule(context)
		if err != nil {
			return 0, err
		}

		relation.Rules[v.Type] = ruleID
	}

	return registerValidationRelationship(relation, context)
}

func (r *ValidationExpression) registerRule(
	context contract.Context) (contract.ServiceRuleID, error) {
	switch r.Type {
	case Time, Location, FaceRecognize, ObjectRecognize:
		return r.registerFromContract(r.Type.ContractName(), context)
	default:
		return 0, fmt.Errorf("编码为%d的类型尚未支持", r.Type)
	}
}

func (r *ValidationExpression) registerFromContract(contractName string,
	context contract.Context) (contract.ServiceRuleID, error) {
	args := [][]byte{
		[]byte(r.Expression),
	}

	rtn := context.InvokeContract(contractName, contract.RegisterFunctionName, args)
	if rtn.Err != nil {
		return 0, rtn.Err
	}

	buf := bytes.Buffer{}
	buf.Write(rtn.Payload)
	id, err := common.ReadUint32(&buf)
	if err != nil {
		return 0, err
	}

	return contract.ServiceRuleID(id), nil
}
