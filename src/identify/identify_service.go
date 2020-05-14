package main

import (
	"audit/contract"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"identify/invokes"
)

type IdentifyService struct {
}

func (s *IdentifyService) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// 初始化物体识别预言机服务相关信息……
	return shim.Success(nil)
}

func (s *IdentifyService) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	switch fn {
	case invokes.RegisterFunctionName:
		return invokes.RegisterMain(args, stub)
	case contract.ValidationFunctionName:
		return invokes.ValidateTimeMain(args)
	default:
		return shim.Error(fmt.Sprintf("找不到名为%s的方法，调用失败", fn))
	}
}

func main() {
	if err := shim.Start(new(IdentifyService)); err != nil {
		fmt.Printf("智能合约启动出错，详细信息：%s", err)
	}
}
