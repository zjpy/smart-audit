package project

import (
	"audit/common"
	"audit/orgnization"
	"audit/record"
	"audit/rules"
	"bytes"
	"errors"
	"fmt"
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
	Timestamp uint64

	// 登录涉及的参数
	Params []string

	// 用于标记隶属于一个业务下该审计当事人第几次录入
	Index uint32

	stub shim.ChaincodeStubInterface
}

func (r *Registration) CountKey() string {
	return r.ID.String() + countKeySuffix
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
	w := new(bytes.Buffer)
	auditeeValue, err := r.AuditeeSpecification.Value()
	if err != nil {
		return nil, err
	}
	// 此处只序列化Value就行？（多序列化了一个字节）
	if err := common.WriteVarBytes(w, auditeeValue); err != nil {
		return nil, errors.New("failed to serialize registration auditee value")
	}
	if err := common.WriteUint64(w, r.Timestamp); err != nil {
		return nil, errors.New("failed to serialize registration timestamp")
	}
	if err := common.WriteVarUint(w, uint64(len(r.Params))); err != nil {
		return nil, errors.New("failed to serialize length of registration params")
	}
	for _, v := range r.Params {
		if err := common.WriteVarString(w, v); err != nil {
			return nil, errors.New("failed to serialize params")
		}
	}
	return w.Bytes(), nil
}

func RegistrationFromString(args []string,
	stub shim.ChaincodeStubInterface) (*Registration, error) {
	if len(args) < 4 {
		return nil, errors.New("初始化成员参数不足")
	}
	// 开始审计事件创建
	registration := &Registration{
		stub: stub,
	}
	// 获取审计当事人ID、项目ID、规则ID构建审计事件ID
	auditeeID, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应当事人ID出错，详细信息：%s", err.Error()))
	}
	projectID, err := strconv.ParseUint(args[1], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析项目ID出错，详细信息：%s", err.Error()))
	}
	ruleID, err := strconv.ParseUint(args[2], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析规则ID出错，详细信息：%s", err.Error()))
	}
	u32AuditeeID := common.Uint32ToBytes(uint32(auditeeID))
	u32ProjectID := common.Uint32ToBytes(uint32(projectID))
	u32RuleID := common.Uint32ToBytes(uint32(ruleID))
	eventID, err := common.Uint256FromBytes(append(append(append(
		[]byte{}, u32AuditeeID[:]...), u32ProjectID[:]...), u32RuleID[:]...))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("构建审计事件ID出错，详细信息：%s", err.Error()))
	}
	registration.ID = *eventID

	// 获取时间戳
	timeStamp, err := strconv.ParseUint(args[3], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应当事人ID出错，详细信息：%s", err.Error()))
	}
	registration.Timestamp = timeStamp

	// todo 先只构建只有Key信息的审计当事人、项目、规则，如果需要补全则需调用合约根据ID查询
	registration.Auditee = orgnization.Auditee{
		Member: &orgnization.Member{ID: uint32(auditeeID)}}
	registration.Project = Project{ID: uint32(projectID)}
	registration.Rule = rules.ValidationRelationship{ID: uint32(ruleID)}

	// 构建用于规则验证的参数，第一个参数为规则ID
	registration.Params = []string{strconv.Itoa(int(registration.Rule.ID))}
	for i := 4; i < len(args); i++ {
		registration.Params = append(registration.Params, args[i])
	}

	// 获取第几次录入信息
	index, err := record.GetRecordCount(GetRegistrationCountKey(registration.ID), stub)
	if err != nil {
		return nil, err
	}
	registration.Index = index
	return registration, nil
}

func GetRegistrationCountKey(specID common.Uint256) string {
	return specID.String() + countKeySuffix
}
