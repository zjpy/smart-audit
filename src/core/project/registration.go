package project

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	RegistrationCountKeySuffix = "_count"
)

type RegistrationCount map[string]uint32

// 用于记录一次信息登录
type Registration struct {
	// 审计当事人规范对象
	AuditeeSpecification

	// 登录发生的时间
	Timestamp uint64

	// 登录涉及的参数
	Params []string

	// 用于标记隶属于一个业务下该审计当事人第几次录入
	Index uint32
}

func (r *Registration) CountKey() string {
	return r.ID.String() + RegistrationCountKeySuffix
}

func (r *Registration) GetCount() uint32 {
	return r.Index + 1
}

func (r *Registration) Validate() error {
	// todo complete me
	return nil
}

func (r *Registration) Key() string {
	return r.ID.String() + strconv.FormatUint(uint64(r.Index), 10)
}

func (r *Registration) Value() ([]byte, error) {
	value, err := json.Marshal(r)
	if err != nil {
		return nil, errors.New("failed to marshal auditee specification")
	}

	return value, nil
}
