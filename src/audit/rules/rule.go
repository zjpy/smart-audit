package rules

import (
	"audit/common"
	"audit/record"
	"bytes"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	rulePrefix   = "rule-"
	ruleCountKey = "rule-count"
)

// 用于定义一种规则模式的结构
type Rule struct {
	// 规则名称
	Name string

	// 规则唯一标识
	ID uint32

	// 具体规则定义
	Expression string

	stub shim.ChaincodeStubInterface
}

func (r *Rule) CountKey() string {
	return ruleCountKey
}

func (r *Rule) GetCount() uint32 {
	return r.ID + 1
}

func (r *Rule) Validate() error {
	// todo complete me
	return nil
}

func (r *Rule) Key() string {
	return rulePrefix + strconv.FormatUint(uint64(r.ID), 10)
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

func FromStrings(args []string, stub shim.ChaincodeStubInterface) (*Rule, error) {
	if len(args) < 2 {
		return nil, errors.New("解析规则参数不足")
	}
	count, err := record.GetRecordCount(ruleCountKey, stub)
	if err != nil {
		return nil, err
	}

	return &Rule{
		Name:       args[0],
		Expression: args[1],
		ID:         count,
		stub:       stub,
	}, nil
}
