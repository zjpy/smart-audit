package rules

import (
	"core/contract"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type ValidationValue struct {
	// 规则类型
	Type RuleType

	// 验证值
	ActualValues string

	// 规则表达式在预言机中对应的ID
	ID contract.ServiceRuleID
}

func ValidateRules(relationID uint32, expressions []string,
	context contract.Context) error {

	relation := &ValidationRelationship{
		Rules: make(map[RuleType]contract.ServiceRuleID, 0),
		ID:    relationID}
	value, err := context.GetState(relation.Key())
	if err != nil {
		return errors.New("规则ID对应的规则不存在，详细信息：" + err.Error())
	}

	if err := json.Unmarshal(value, relation); err != nil {
		return errors.New("规则解析失败，详细信息：" + err.Error())
	}

	items, err := parseRuleValues(expressions)
	if err != nil {
		return err
	}

	if err = setRuleIds(relation, items); err != nil {
		return err
	}

	for _, v := range items {
		if err := v.Validate(context); err != nil {
			return err
		}
	}
	return nil
}

func setRuleIds(relation *ValidationRelationship, items []*ValidationValue) error {
	if len(items) != len(relation.Rules) {
		return errors.New("给定的规则值数与需要验证的规则数不匹配")
	}

	var ok bool
	for _, v := range items {
		v.ID, ok = relation.Rules[v.Type]
		if !ok {
			return fmt.Errorf("规则类型%s不在需要验证的列表中", v.Type.ContractName())
		}
	}
	return nil
}

func parseRuleValues(expressions []string) ([]*ValidationValue, error) {
	var results []*ValidationValue
	for i := 0; i+1 < len(expressions); i += 2 {
		t, err := getRuleType(expressions[i])
		if err != nil {
			return nil, err
		}
		results = append(results, &ValidationValue{
			Type:         t,
			ActualValues: expressions[i+1],
		})
	}
	return results, nil
}

func (i *ValidationValue) Validate(context contract.Context) error {
	switch i.Type {
	case None:
		return nil
	case Time, Location, FaceRecognize, ObjectRecognize:
		return i.validateFromContract(i.Type.ContractName(), context)
	default:
		return fmt.Errorf("编码为%d的类型尚未支持", i.Type)
	}
}

func (i *ValidationValue) validateFromContract(contractName string,
	context contract.Context) error {
	args := []string{
		strconv.FormatUint(uint64(i.ID), 32),
		i.ActualValues,
	}

	if rtn := context.InvokeContract(contractName,
		contract.ValidationFunctionName, args); rtn.Err != nil {
		return rtn.Err
	}
	return nil
}
