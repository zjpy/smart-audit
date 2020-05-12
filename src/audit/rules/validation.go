package rules

import (
	"audit/contract"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ValidationItem struct {
	// 条件操作符
	Condition contract.ConditionalOperator

	// 规则类型
	Type RuleType

	// 验证值
	ActualValues []string
}

func ValidateRules(expression string, stub shim.ChaincodeStubInterface) error {
	items := parserRules(expression)
	for _, v := range items {
		if err := v.Validate(stub); err != nil {
			return err
		}
	}
	return nil
}

func parserRules(expression string) []ValidationItem {
	// todo 这里会解析出来一个验证的多个参数，其中第一个参数是ServiceRuleID，用于定义和验证服务中的规则对应关系
	return nil
}

func (i *ValidationItem) Validate(stub shim.ChaincodeStubInterface) error {
	switch i.Type {
	case None:
		return nil
	case Time, Location, FaceRecognize, ObjectRecognize:
		return i.validateFromContract(i.Type.ContractName(), stub)
	default:
		return fmt.Errorf("编码为%d的类型尚未支持", i.Type)
	}
}

func (i *ValidationItem) validateFromContract(contractName string,
	stub shim.ChaincodeStubInterface) error {
	args := [][]byte{[]byte(contract.ValidationFunctionName)}
	for _, v := range i.ActualValues {
		args = append(args, []byte(v))
	}

	if rtn := stub.InvokeChaincode(contractName, args,
		""); rtn.Status != shim.OK {
		return errors.New(rtn.Message)
	}
	return nil
}
