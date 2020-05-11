package rules

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

// 用于定义一种规则模式的结构
type Rule struct {
	// 规则名称
	Name string

	// 规则唯一标识
	ID uint32

	// 具体规则定义
	Expression string
}

func (r *Rule) Validate(stub shim.ChaincodeStubInterface) error {
	// todo complete me
	return nil
}

func (r *Rule) Key() string {
	return strconv.FormatUint(uint64(r.ID), 10)
}

func (r *Rule) Value() []byte {
	// todo 除ID之外的序列化工作
	return nil
}

func FromStrings(args []string) (*Rule, error) {
	// todo 根据参数构造Rule
	return nil, nil
}
