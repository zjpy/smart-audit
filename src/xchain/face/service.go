package main

import (
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/driver"
	"oracles/face/invokes"
	"xchain/contract"
)

// 用于处理与人脸识别预言机交互的智能合约
type FaceService struct {
}

// 人脸识别合约初始化
func (s *FaceService) Initialize(ctx code.Context) code.Response {
	// 初始化人脸识别预言机服务相关信息……
	return code.OK(nil)
}

// 对应 core/contract包下的RegisterFunctionName方法定义
func (s *FaceService) Register(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.RegisterMain(args, context))
}

// 对应 core/contract包下的ValidationFunctionName方法定义
func (s *FaceService) Validation(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.ValidateMain(args))
}

// 人脸识别主程序入口
func main() {
	driver.Serve(new(FaceService))
}
