package rules

import (
	"audit/contract"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

type ValidationValue struct {
	// 规则类型
	Type RuleType

	// 验证值
	ActualValues string

	// 规则表达式在预言机中对应的ID
	ID contract.ServiceRuleID
}

func ValidateRules(ruleID uint32, expressions []string,
	stub shim.ChaincodeStubInterface) error {

	// todo 这里通过ruleID获取相应的ValidationRelationship
	relation := &ValidationRelationship{}

	items, err := parseRuleValues(expressions)
	if err != nil {
		return err
	}

	if err = setRuleIds(relation, items); err != nil {
		return err
	}

	for _, v := range items {
		if err := v.Validate(stub); err != nil {
			return err
		}
	}
	return nil
}

func setRuleIds(relation *ValidationRelationship, items []*ValidationValue) error {
	if len(items) != len(relation.Rules) {
		return errors.New("给定的规则值数与需要验证的规则数不匹配")
	}

	var ok bool
	for _, v := range items {
		v.ID, ok = relation.Rules[v.Type]
		if !ok {
			return fmt.Errorf("规则类型%s不在需要验证的列表中", v.Type.ContractName())
		}
	}
	return nil
}

func parseRuleValues(expressions []string) ([]*ValidationValue, error) {
	var results []*ValidationValue
	for i := 0; i+1 < len(expressions); i += 2 {
		t, err := getRuleType(expressions[i])
		if err != nil {
			return nil, err
		}
		results = append(results, &ValidationValue{
			Type:         t,
			ActualValues: expressions[i+1],
		})
	}
	return results, nil
}

func (i *ValidationValue) Validate(stub shim.ChaincodeStubInterface) error {
	switch i.Type {
	case None:
		return nil
	case Time, Location, FaceRecognize, ObjectRecognize:
		return i.validateFromContract(i.Type.ContractName(), stub)
	default:
		return fmt.Errorf("编码为%d的类型尚未支持", i.Type)
	}
}

func (i *ValidationValue) validateFromContract(contractName string,
	stub shim.ChaincodeStubInterface) error {
	args := [][]byte{
		[]byte(contract.ValidationFunctionName),
		[]byte(strconv.FormatUint(uint64(i.ID), 32)),
		[]byte(i.ActualValues),
	}

	// 这里channel参数为空则默认会发送到当前合约所在channel上
	if rtn := stub.InvokeChaincode(contractName, args,
		""); rtn.Status != shim.OK {
		return errors.New(rtn.Message)
	}
	return nil
}
