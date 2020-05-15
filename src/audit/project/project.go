package project

import (
	"audit/common"
	"audit/record"
	"bytes"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

const (
	projectPrefix   = "project-"
	projectCountKey = "project-count"
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

func (p *Project) CountKey() string {
	return projectCountKey
}

func (p *Project) GetCount() uint32 {
	return p.ID + 1
}

func (p *Project) Validate() error {
	// todo complete me
	return nil
}

func (p *Project) Key() string {
	return projectPrefix + strconv.FormatUint(uint64(p.ID), 10)
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

func FromStrings(args []string, stub shim.ChaincodeStubInterface) (*Project, error) {
	if len(args) < 2 {
		return nil, errors.New("参数不足")
	}
	count, err := record.GetRecordCount(projectCountKey, stub)
	if err != nil {
		return nil, err
	}

	return &Project{
		ID:          count,
		Name:        args[0],
		Description: args[1],
	}, nil
}
