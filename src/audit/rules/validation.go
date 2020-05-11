package rules

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ValidationItem struct {
	// 条件操作符
	Condition ConditionalOperator

	// 规则类型
	Type RuleType

	// 验证值
	ActualValue string
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
	// todo complete me
	return nil
}

func (i *ValidationItem) Validate(stub shim.ChaincodeStubInterface) error {
	switch i.Type {
	case None:
		return nil
	case Time:
		return nil
	case Location:
		return nil
	case FaceRecognize:
		return nil
	case ObjectRecognize:
		return nil
	default:
		return fmt.Errorf("编码为%d的类型尚未支持", i.Type)
	}
}
