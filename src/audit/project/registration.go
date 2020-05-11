package project

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
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

func (r *Registration) Key() string {
	return string(r.ID) + strconv.FormatUint(uint64(r.Index), 10)
}

func (r *Registration) Value() []byte {
	// todo complete me
	return nil
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
