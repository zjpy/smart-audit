package invokes

import (
	"audit/rules"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegisterRuleMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	err := rules.RegisterRules(args, stub)
	if err != nil {
		return shim.Error(fmt.Sprint("注册规则失败，详细信息：", err))
	}

	return shim.Success(nil)
}
