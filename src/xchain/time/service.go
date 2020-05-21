package main

import (
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/driver"
	"oracles/time/invokes"
	"xchain/contract"
)

type TimeService struct {
}

func (s *TimeService) Initialize(ctx code.Context) code.Response {
	// 初始化时间预言机服务相关信息……
	return code.OK(nil)
}

// 对应 core/contract包下的RegisterFunctionName方法定义
func (s *TimeService) Register(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.RegisterMain(args, context))
}

// 对应 core/contract包下的ValidationFunctionName方法定义
func (s *TimeService) Validation(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.ValidateMain(args))
}

func main() {
	driver.Serve(new(TimeService))
}
