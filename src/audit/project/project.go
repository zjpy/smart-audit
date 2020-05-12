package project

import (
	"audit/common"
	"bytes"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// 用于定义一个审计业务的结构
type Project struct {
	// 业务名称
	Name string

	// 业务唯一标识
	ID uint32

	// 业务相关描述
	Description string
}

func (p *Project) Validate(stub shim.ChaincodeStubInterface) error {
	// todo complete me
	return nil
}

func (p *Project) Key() string {
	return strconv.FormatUint(uint64(p.ID), 10)
}

func (p *Project) Value() ([]byte, error) {
	w := new(bytes.Buffer)
	if err := common.WriteVarString(w, p.Name); err != nil {
		return nil, errors.New("failed to serialize project name")
	}
	if err := common.WriteVarString(w, p.Description); err != nil {
		return nil, errors.New("failed to serialize project description")
	}
	return w.Bytes(), nil
}

func FromStrings(args []string) (*Project, error) {
	// todo complete me
	return nil, nil
}
