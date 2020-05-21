package project

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	Prefix   = "project-"
	CountKey = "project-count"
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
	return CountKey
}

func (p *Project) GetCount() uint32 {
	return p.ID + 1
}

func (p *Project) Validate() error {
	// todo complete me
	return nil
}

func (p *Project) Key() string {
	return Prefix + strconv.FormatUint(uint64(p.ID), 10)
}

func (p *Project) Value() ([]byte, error) {
	value, err := json.Marshal(p)
	if err != nil {
		return nil, errors.New("failed to marshal project")
	}

	return value, nil
}
