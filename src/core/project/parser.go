package project

import (
	"core/common"
	"core/contract"
	"core/orgnization"
	"core/record"
	"core/rules"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// 从输入参数获取项目信息
func FromStrings(args []string, context contract.Context) (*Project, error) {
	if len(args) < 4 {
		return nil, errors.New("参数不足")
	}
	count, err := record.GetRecordCount(CountKey, context)
	if err != nil {
		return nil, err
	}

	auditRulesMap := make(map[string]string, 0)
	for i := 2; i+1 < len(args); i += 2 {
		auditeeID, err := strconv.ParseUint(args[i], 10, 32)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("解析审计当事人ID出错，详细信息：%s", err.Error()))
		}
		auditeeKey, err := getAuditeeKeyFromID(uint32(auditeeID), context)

		ruleID, err := strconv.ParseUint(args[i+1], 10, 32)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("解析规则ID出错，详细信息：%s", err.Error()))
		}
		rule := rules.ValidationRelationship{ID: uint32(ruleID)}
		ruleBytes, err := context.GetState(rule.Key())

		auditRulesMap[auditeeKey] = string(ruleBytes)
	}

	return &Project{
		ID:              count,
		Name:            args[0],
		Description:     args[1],
		AuditeeRulesMap: auditRulesMap,
	}, nil
}

// 从输入参数获取审计事件信息
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
	ruleID, err := getRelationID(uint32(auditeeID), uint32(projectID), context)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应规则ID出错，详细信息：%s", err.Error()))
	}

	// 获取审计当事人ID、项目ID、规则ID构建审计事件ID
	eventID := GenerateEventID(uint32(auditeeID), uint32(projectID), ruleID)
	registration.ID = eventID

	// 获取时间戳
	timeStamp, err := strconv.ParseUint(args[2], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应当事人ID出错，详细信息：%s", err.Error()))
	}
	registration.Timestamp = timeStamp

	// 先只构建只有Key信息的审计当事人、项目、规则，如果需要补全则需调用合约根据ID查询
	registration.Auditee = orgnization.Auditee{
		Member: &orgnization.Member{ID: uint32(auditeeID)}}
	registration.Project = Project{ID: uint32(projectID)}
	registration.Rule = rules.ValidationRelationship{
		Rules: make(map[rules.RuleType]contract.ServiceRuleID, 0),
		ID:    uint32(ruleID)}

	// 构建用于规则验证的参数
	registration.Params = make([]string, 0)
	for i := 3; i < len(args); i++ {
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
func GetEventID(args []string, context contract.Context) (*common.Uint256, error) {

	auditeeID, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应当事人ID出错，详细信息：%s", err.Error()))
	}

	projectID, err := strconv.ParseUint(args[1], 10, 32)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应项目ID出错，详细信息：%s", err.Error()))
	}

	ruleID, err := getRelationID(uint32(auditeeID), uint32(projectID), context)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("解析审计事件对应规则ID出错，详细信息：%s", err.Error()))
	}

	u32AuditeeID := common.Uint32ToBytes(uint32(auditeeID))
	u32ProjectID := common.Uint32ToBytes(uint32(projectID))
	u32RuleID := common.Uint32ToBytes(ruleID)

	ids := append(append(append(
		[]byte{}, u32AuditeeID[:]...), u32ProjectID[:]...), u32RuleID[:]...)
	var eventID common.Uint256
	copy(eventID[:], ids)

	return &eventID, nil
}

// 根据传入的参数获取审计事件ID
func GenerateEventID(auditeeID, projectID, ruleID uint32) common.Uint256 {
	u32AuditeeID := common.Uint32ToBytes(auditeeID)
	u32ProjectID := common.Uint32ToBytes(projectID)
	u32RuleID := common.Uint32ToBytes(ruleID)

	ids := append(append(append(
		[]byte{}, u32AuditeeID[:]...), u32ProjectID[:]...), u32RuleID[:]...)
	var eventID common.Uint256
	copy(eventID[:], ids)
	return eventID
}

// 获取当前存储的审计事件计数
func GetRegistrationCountKey(specID common.Uint256) string {
	return specID.String() + RegistrationCountKeySuffix
}

// 通过审计当事人ID和项目ID获取当事人在该项目下对应的规则关系ID
func getRelationID(auditeeID, projectID uint32, context contract.Context) (uint32, error) {
	pj, err := getProjectObject(projectID, context)
	if err != nil {
		return 0, err
	}

	auditeeKey, err := getAuditeeKeyFromID(auditeeID, context)
	if err != nil {
		return 0, err
	}

	ruleContent, ok := pj.AuditeeRulesMap[auditeeKey]
	if !ok {
		return 0, errors.New("未找到规则")
	}

	relation := &rules.ValidationRelationship{}
	if err := json.Unmarshal([]byte(ruleContent), relation); err != nil {
		return 0, err
	}
	return relation.ID, nil
}

// 通过项目ID获取项目对象
func getProjectObject(projectID uint32, context contract.Context) (*Project, error) {
	pj := &Project{ID: projectID}
	ruleBuf, err := context.GetState(pj.Key())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(ruleBuf, pj); err != nil {
		return nil, errors.New("项目解析失败，详细信息：" + err.Error())
	}
	return pj, nil
}

// 通过审计当事人ID生成相应的ID值
func getAuditeeKeyFromID(auditeeID uint32, context contract.Context) (string, error) {
	auditee := orgnization.Auditee{Member: &orgnization.Member{ID: uint32(auditeeID)}}
	auditBytes, err := context.GetState(auditee.Key())
	if err != nil {
		return "", errors.New(fmt.Sprintf("不存在审计当事人ID对应的审计当事人"))
	}
	return string(auditBytes), nil
}
