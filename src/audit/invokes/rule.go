package invokes

import (
	"audit/record"
	"audit/rules"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegisterRuleMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	rule, err := rules.FromStrings(args, stub)
	if err != nil {
		return shim.Error(fmt.Sprint("解析规则失败，详细信息：", err))
	}

	if err = rule.Validate(); err != nil {
		return shim.Error(fmt.Sprintf("规则%s数据验证失败，详细信息：%s", rule.Key(), err))
	}

	if err = record.StoreItem(rule, stub); err != nil {
		return shim.Error(fmt.Sprintf("规则%s存储失败，详细信息：%s", rule.Key(), err))
	}
	if err = record.StoreCount(rule, stub); err != nil {
		return shim.Error(fmt.Sprintf("规则%s相应的索引值存储失败，详细信息：%s",
			rule.Key(), err))
	}
	return shim.Success(nil)
}
