package main

import (
	"audit/initialize"
	"audit/invokes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SmartAudit struct {
}

func (s *SmartAudit) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return initialize.InitMain(stub)
}

func (s *SmartAudit) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	switch fn {
	case invokes.RegisterRule:
		return invokes.RegisterRuleMain(args, stub)
	case invokes.RegisterAuditee:
		return invokes.RegisterAuditeeMain(args, stub)
	case invokes.RegisterProject:
		return invokes.RegisterProjectMain(args, stub)
	case invokes.AddEvent:
		return invokes.AddEventMain(args, stub)
	case invokes.GetAuditee:
		return invokes.GetAuditeeMain(args, stub)
	case invokes.GetRule:
		return invokes.GetRuleMain(args, stub)
	case invokes.GetProject:
		return invokes.GetProjectMain(args, stub)
	case invokes.GetMaintainers:
		return invokes.QueryMaintainersMain(stub)
	case invokes.QueryEvents:
		invokes.QueryEventMain(args, stub)
	default:
		return shim.Error(fmt.Sprintf("找不到名为%s的方法，调用失败", fn))
	}

	return shim.Error(fmt.Sprintf("调用失败"))
}

func main() {
	if err := shim.Start(new(SmartAudit)); err != nil {
		fmt.Printf("智能合约启动出错，详细信息：%s", err)
	}
}
