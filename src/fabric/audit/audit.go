package main

import (
	"core/invokes"
	"fabric/contract"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SmartAudit struct {
}

func (s *SmartAudit) Init(stub shim.ChaincodeStubInterface) peer.Response {
	context := contract.NewContext(stub)
	res := invokes.InitMain(context)
	return contract.Response(res)
}

func (s *SmartAudit) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	context := contract.NewContext(stub)
	args := context.GetArgs()

	switch context.GetFunctionName() {
	case invokes.RegisterRule:
		return contract.Response(invokes.RegisterRuleMain(args, context))
	case invokes.RegisterAuditee:
		return contract.Response(invokes.RegisterAuditeeMain(args, context))
	case invokes.RegisterProject:
		return contract.Response(invokes.RegisterProjectMain(args, context))
	case invokes.AddEvent:
		return contract.Response(invokes.AddEventMain(args, context))
	case invokes.GetAuditee:
		return contract.Response(invokes.GetAuditeeMain(args, context))
	case invokes.GetRule:
		return contract.Response(invokes.GetRuleMain(args, context))
	case invokes.GetProject:
		return contract.Response(invokes.GetProjectMain(args, context))
	case invokes.GetMaintainers:
		return contract.Response(invokes.GetMaintainersMain(context))
	case invokes.QueryEvents:
		return contract.Response(invokes.QueryEventsMain(args, context))
	default:
		return shim.Error(fmt.Sprintf("找不到名为%s的方法，调用失败",
			context.GetFunctionName()))
	}
}

func main() {
	if err := shim.Start(new(SmartAudit)); err != nil {
		fmt.Printf("智能合约启动出错，详细信息：%s", err)
	}
}
