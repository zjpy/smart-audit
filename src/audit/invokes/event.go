package invokes

import (
	"audit/common"
	"audit/project"
	"audit/record"
	"audit/rules"
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func AddEventMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	registration, err := project.RegistrationFromString(args, stub)
	if err != nil {
		return shim.Error(fmt.Sprint("合规事件登录失败，详细信息：", err))
	}

	if err = verify(registration, stub); err != nil {
		return shim.Error(fmt.Sprint("合规事件数据验证失败，详细信息：", err))
	}

	if err = record.StoreItem(registration, stub); err != nil {
		return shim.Error(fmt.Sprintf("合规事件%s存储失败，详细信息：%s",
			registration.Key(), err))
	}
	// 存储审计当事人规范对象所对应的存储条数
	if err = record.StoreCount(registration, stub); err != nil {
		return shim.Error(fmt.Sprintf("合规事件%s相应的索引值存储失败，详细信息：%s",
			registration.Key(), err))
	}
	return shim.Success([]byte("OK"))
}

func verify(registration *project.Registration, stub shim.ChaincodeStubInterface) error {
	if err := registration.Validate(); err != nil {
		return fmt.Errorf("合规事件%s数据验证失败，详细信息：%s", registration.ID, err)
	}

	if err := rules.ValidateRules(registration.Rule.ID, registration.Params,
		stub); err != nil {
		return fmt.Errorf("合规事件%s规则验证失败，详细信息：%s", registration.ID, err)
	}

	return nil
}

// 根据审计当事人ID、项目ID以及规则ID，获取所有审计当事人的审计事件
func QueryEventMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	if len(args) < 3 {
		return shim.Error("查询失败，需要提查询供审计事件对应的当事人ID、项目ID以及规则ID")
	}

	auditeeID, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return shim.Error(fmt.Sprintf("解析审计事件对应当事人ID出错，详细信息：%s", err.Error()))
	}
	projectID, err := strconv.ParseUint(args[1], 10, 32)
	if err != nil {
		return shim.Error(fmt.Sprintf("解析审计事件对应项目ID出错，详细信息：%s", err.Error()))
	}
	ruleID, err := strconv.ParseUint(args[2], 10, 32)
	if err != nil {
		return shim.Error(fmt.Sprintf("解析审计事件对应规则ID出错，详细信息：%s", err.Error()))
	}
	u32AuditeeID := common.Uint32ToBytes(uint32(auditeeID))
	u32ProjectID := common.Uint32ToBytes(uint32(projectID))
	u32RuleID := common.Uint32ToBytes(uint32(ruleID))
	eventID, err := common.Uint256FromBytes(append(append(append(
		[]byte{}, u32AuditeeID[:]...), u32ProjectID[:]...), u32RuleID[:]...))
	if err != nil {
		return shim.Error(fmt.Sprintf("构建审计事件ID出错，详细信息：%s", err.Error()))
	}

	// 获取第几次录入信息
	index, err := record.GetRecordCount(project.GetRegistrationCountKey(*eventID), stub)
	if err != nil {
		return shim.Error(fmt.Sprintf("构建审计事件ID出错，详细信息：%s", err.Error()))
	}

	// 获取开始及结束Key，查询满足条件的所有记录
	reg := project.Registration{AuditeeSpecification: project.AuditeeSpecification{
		ID: *eventID}}
	startKey := reg.Key()
	reg.Index = index
	endKey := reg.Key()
	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取审计事件信息出错，详细信息：%s", err.Error()))
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false
	buffer.WriteString(`{"result":[`)

	// todo 将结果反序列化，转为json再显示
	for resultsIterator.HasNext() {
		//获取迭代器中的每一个值
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Fail")
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		//将查询结果放入Buffer中
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString(`]}`)
	fmt.Printf("Query result: %s", buffer.String())
	return shim.Success(buffer.Bytes())
}
