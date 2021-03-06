package main

import (
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/driver"
	"oracles/identify/invokes"
	"xchain/contract"
)

// 用于处理与物体识别预言机交互的智能合约
type IdentifyService struct {
}

// 物体识别合约初始化
func (s *IdentifyService) Initialize(ctx code.Context) code.Response {
	// 初始化物体识别预言机服务相关信息……
	return code.OK(nil)
}

// 对应 core/contract包下的RegisterFunctionName方法定义
func (s *IdentifyService) Register(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.RegisterMain(args, context))
}

// 对应 core/contract包下的ValidationFunctionName方法定义
func (s *IdentifyService) Validation(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.ValidateMain(args))
}

// 物体识别主程序入口
func main() {
	driver.Serve(new(IdentifyService))
}
