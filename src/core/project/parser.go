package project

import (
	"core/common"
	"core/contract"
	"core/orgnization"
	"core/record"
	rules2 "core/rules"
	"errors"
	"fmt"
	"strconv"
)

func FromStrings(args []string, context contract.Context) (*Project, error) {
	if len(args) < 2 {
		return nil, errors.New("参数不足")
	}
	count, err := record.GetRecordCount(CountKey, context)
	if err != nil {
		return nil, err
	}

	return &Project{
		ID:          count,
		Name:        args[0],
		Description: args[1],
	}, nil
}

func RegistrationFromString(args []string,
	context contract.Context) (*Registration, error) {
	if len(args) < 4 {
		return nil, errors.New("初始化成员参数不足")
	}
	// 开始审计事件创建
	registration := &Registration{}

	// 获取审计当事人ID
	auditeeID, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应当事人ID出错，详细信息：%s", err.Error()))
	}

	// 获取项目ID
	projectID, err := strconv.ParseUint(args[1], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应项目ID出错，详细信息：%s", err.Error()))
	}

	// 获取规则ID
	ruleID, err := strconv.ParseUint(args[2], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应规则ID出错，详细信息：%s", err.Error()))
	}

	// 获取审计当事人ID、项目ID、规则ID构建审计事件ID
	eventID := GetEventID(uint32(auditeeID), uint32(projectID), uint32(ruleID))
	registration.ID = eventID

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
	registration.Rule = rules2.ValidationRelationship{
		Rules: make(map[rules2.RuleType]contract.ServiceRuleID, 0),
		ID:    uint32(ruleID)}

	// 构建用于规则验证的参数
	registration.Params = make([]string, 0)
	for i := 4; i < len(args); i++ {
		registration.Params = append(registration.Params, args[i])
	}

	// 获取第几次录入信息
	index, err := record.GetRecordCount(GetRegistrationCountKey(registration.ID), context)
	if err != nil {
		return nil, err
	}
	registration.Index = index
	return registration, nil
}

// 根据传入的参数获取审计事件ID
func GetEventID(auditeeID, projectID, ruleID uint32) common.Uint256 {

	u32AuditeeID := common.Uint32ToBytes(auditeeID)
	u32ProjectID := common.Uint32ToBytes(projectID)
	u32RuleID := common.Uint32ToBytes(ruleID)

	ids := append(append(append(
		[]byte{}, u32AuditeeID[:]...), u32ProjectID[:]...), u32RuleID[:]...)
	var eventID common.Uint256
	copy(eventID[:], ids)
	return eventID
}

func GetRegistrationCountKey(specID common.Uint256) string {
	return specID.String() + RegistrationCountKeySuffix
}
