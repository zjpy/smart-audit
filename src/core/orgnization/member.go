package orgnization

import (
	"core/contract"
	"encoding/json"
	"errors"
	"strconv"
)

// 一个组织中成员的基本结构
type Member struct {
	// 成员名称
	Name string

	// 成员唯一标识
	ID uint32
}

func (m *Member) Validate() error {
	// todo complete me
	return nil
}

func (m *Member) Key() string {
	return strconv.FormatUint(uint64(m.ID), 10)
}

func (m *Member) Value() ([]byte, error) {
	value, err := json.Marshal(m)
	if err != nil {
		return nil, errors.New("failed to marshal member")
	}

	return value, nil
}

func MemberFromString(args []string,
	context contract.Context) (mem *Member, err error) {
	if len(args) < 1 {
		err = errors.New("初始化成员参数不足")
		return
	}

	mem = &Member{
		Name: args[0],
	}
	return
}