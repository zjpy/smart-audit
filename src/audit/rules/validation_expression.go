package rules

import (
	"audit/common"
	"audit/contract"
	"bytes"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ValidationExpression struct {
	// 规则类型
	Type RuleType

	// 具体验证规则
	Expression string
}

func RegisterRules(expression []string, stub shim.ChaincodeStubInterface) error {
	op, expressions, err := Parse(expression)
	if err != nil {
		return err
	}

	relation := &ValidationRelationship{
		Operator: op,
	}
	for _, v := range expressions {
		ruleID, err := v.registerRule(stub)
		if err != nil {
			return err
		}

		relation.Rules[ruleID] = v.Type
	}

	return registerValidationRelationship(relation, stub)
}

func (r *ValidationExpression) registerRule(stub shim.ChaincodeStubInterface) (uint32, error) {
	switch r.Type {
	case Time, Location, FaceRecognize, ObjectRecognize:
		return r.registerFromContract(r.Type.ContractName(), stub)
	default:
		return 0, fmt.Errorf("编码为%d的类型尚未支持", r.Type)
	}
}

func (r *ValidationExpression) registerFromContract(contractName string,
	stub shim.ChaincodeStubInterface) (uint32, error) {
	args := [][]byte{
		[]byte(contract.RegisterFunctionName),
		[]byte(r.Expression),
	}

	// 这里channel参数为空则默认会发送到当前合约所在channel上
	rtn := stub.InvokeChaincode(contractName, args, "")
	if rtn.Status != shim.OK {
		return 0, errors.New(rtn.Message)
	}

	buf := bytes.Buffer{}
	buf.Write(rtn.Payload)
	return common.ReadUint32(&buf)
}
