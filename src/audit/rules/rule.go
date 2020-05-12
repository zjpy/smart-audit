package rules

import (
	"audit/common"
	"bytes"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
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

func (r *Rule) Value() ([]byte, error) {
	w := new(bytes.Buffer)
	if err := common.WriteVarString(w, r.Name); err != nil {
		return nil, errors.New("failed to serialize rule name")
	}
	if err := common.WriteVarString(w, r.Expression); err != nil {
		return nil, errors.New("failed to serialize member expression")
	}
	return w.Bytes(), nil
}

func FromStrings(args []string) (*Rule, error) {
	// todo 根据参数构造Rule
	return nil, nil
}
