package project

import (
	"audit/record"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
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
	value, err := json.Marshal(p)
	if err != nil {
		return nil, errors.New("failed to marshal project")
	}

	return value, nil
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
