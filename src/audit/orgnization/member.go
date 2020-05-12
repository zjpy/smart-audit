package orgnization

import (
	"audit/common"
	"bytes"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
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

func (m *Member) Validate(stub shim.ChaincodeStubInterface) error {
	// todo complete me
	return nil
}

func (m *Member) Key() string {
	return strconv.FormatUint(uint64(m.ID), 10)
}

func (m *Member) Value() ([]byte, error) {
	w := new(bytes.Buffer)
	if err := common.WriteVarString(w, m.Name); err != nil {
		return nil, errors.New("failed to serialize member name")
	}
	if err := common.WriteVarBytes(w, m.PK); err != nil {
		return nil, errors.New("failed to serialize member PK")
	}
	return w.Bytes(), nil
}
