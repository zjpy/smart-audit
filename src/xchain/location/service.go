package main

import (
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/driver"
	"oracles/location/invokes"
	"xchain/contract"
)

// 用于处理与定位服务预言机交互的智能合约
type LocationService struct {
}

// 定位服务合约初始化
func (s *LocationService) Initialize(ctx code.Context) code.Response {
	// 初始化定位预言机服务相关信息……
	return code.OK(nil)
}

// 对应 core/contract包下的RegisterFunctionName方法定义
func (s *LocationService) Register(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.RegisterMain(args, context))
}

// 对应 core/contract包下的ValidationFunctionName方法定义
func (s *LocationService) Validation(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.ValidateMain(args))
}

// 定位服务主程序入口
func main() {
	driver.Serve(new(LocationService))
}
