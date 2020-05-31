package main

import (
	core "core/contract"
	"fabric/contract"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"log"
	"oracles/time/invokes"
)

type TimeService struct {
}

// 时间服务合约初始化
func (s *TimeService) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// 初始化时间预言机服务相关信息……
	return shim.Success(nil)
}

// 时间服务合约方法调用
func (s *TimeService) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	context := contract.NewContext(stub)
	args := context.GetArgs()

	switch context.GetFunctionName() {
	// 注册时间规则
	case core.RegisterFunctionName:
		return contract.Response(invokes.RegisterMain(args, context))
	// 时间验证
	case core.ValidationFunctionName:
		return contract.Response(invokes.ValidateMain(args))
	default:
		return shim.Error(fmt.Sprintf("找不到名为%s的方法，调用失败",
			context.GetFunctionName()))
	}
}

// 时间服务合约主程序入口
func main() {
	if err := shim.Start(new(TimeService)); err != nil {
		log.Printf("智能合约启动出错，详细信息：%s", err)
	}
}
