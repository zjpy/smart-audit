package orgnization

import (
	"strconv"
)

// 一个组织中成员的基本结构
type Member struct {
	// 成员名称
	Name string

	// 成员唯一标识
	ID uint32

	// 成员公钥
	PK []byte
}

func (m *Member) Key() string {
	return strconv.FormatUint(uint64(m.ID), 10)
}

func (m *Member) Value() []byte {
	// todo 除Key之外的信息序列化
	return nil
}
