package invokes

import (
	"audit/common"
	"audit/rules"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// 注册规则，返回规则ID
func RegisterRuleMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	ruleID, err := rules.RegisterRules(args, stub)
	if err != nil {
		return shim.Error(fmt.Sprint("注册规则失败，详细信息：", err))
	}

	return shim.Success(common.Uint32ToBytes(ruleID))
}

// 根据规则ID获取规则信息
func GetRuleMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	if len(args) == 0 {
		return shim.Error("查询失败，需要提供规则ID")
	}
	ruleID, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return shim.Error(fmt.Sprintf("解析规则ID出错，详细信息：%s", err.Error()))
	}

	ruleKey := rules.ValidationRelationship{ID: uint32(ruleID)}.Key()
	ruleBuf, err := stub.GetState(ruleKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取规则出错，详细信息：%s", err.Error()))
	}

	// todo 将结果反序列化，转为json再显示
	return peer.Response{Payload: ruleBuf}
}
