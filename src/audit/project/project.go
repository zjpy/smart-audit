package project

import "strconv"

// 用于定义一个审计业务的结构
type Project struct {
	// 业务名称
	Name string

	// 业务唯一标识
	ID uint32

	// 业务相关描述
	Description string
}

func (p *Project) Key() string {
	return strconv.FormatUint(uint64(p.ID), 10)
}

func (p *Project) Value() []byte {
	// todo complete me
	return nil
}

func FromStrings(args []string) (*Project, error) {
	// todo complete me
	return nil, nil
}
