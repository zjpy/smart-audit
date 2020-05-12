package project

import (
	"audit/common"
	"bytes"
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	countKeySuffix = "_count"
)

type RegistrationCount map[string]uint32

// 用于记录一次信息登录
type Registration struct {
	// 审计当事人规范对象
	AuditeeSpecification

	// 登录发生的时间
	Timestamp uint32

	// 登录涉及的参数
	Params map[string]string

	// 用于标记隶属于一个业务下该审计当事人第几次登录
	Index uint32
}

func (r *Registration) Validate(stub shim.ChaincodeStubInterface) error {
	// todo complete me
	return nil
}

func (r *Registration) Key() string {
	return string(r.ID) + strconv.FormatUint(uint64(r.Index), 10)
}

func (r *Registration) Value() ([]byte, error) {
	w := new(bytes.Buffer)
	auditeeValue, err := r.AuditeeSpecification.Value()
	if err != nil {
		return nil, err
	}
	// 此处只序列化Value就行？（多序列化了一个字节）
	if err := common.WriteVarBytes(w, auditeeValue); err != nil {
		return nil, errors.New("failed to serialize registration auditee value")
	}
	if err := common.WriteUint32(w, r.Timestamp); err != nil {
		return nil, errors.New("failed to serialize registration timestamp")
	}
	if err := common.WriteVarUint(w, uint64(len(r.Params))); err != nil {
		return nil, errors.New("failed to serialize length of registration params")
	}
	for k, v := range r.Params {
		if err := common.WriteVarString(w, k); err != nil {
			return nil, errors.New("failed to serialize params")
		}
		if err := common.WriteVarString(w, v); err != nil {
			return nil, errors.New("failed to serialize params")
		}
	}
	return w.Bytes(), nil
}

func RegistrationFromString(args []string, stub shim.ChaincodeStubInterface) (*Registration, error) {
	// todo 通过参数解析
	registration := &Registration{}

	index, err := getRegistrationCount(registration.ID, stub)
	if err != nil {
		return nil, err
	}

	registration.Index = index
	return registration, nil
}

func GetRegistrationCountKey(specID Uint256) string {
	return string(specID) + countKeySuffix
}

func getRegistrationCount(specID Uint256, stub shim.ChaincodeStubInterface) (
	uint32, error) {
	countBuf, err := stub.GetState(GetRegistrationCountKey(specID))
	if err != nil {
		return 0, err
	}

	// 未找到则从0开始
	if countBuf == nil {
		return 0, nil
	}

	count, err := strconv.ParseUint(string(countBuf), 10, 32)
	return uint32(count), err
}
