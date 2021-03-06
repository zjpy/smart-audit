package main

import (
	"core/invokes"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/driver"
	"xchain/contract"
)

// XuperChain平台上的审计业务智能合约实现
type SmartAudit struct {
}

// 审计合约初始化
func (s *SmartAudit) Initialize(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	res := invokes.InitMain(context)
	return contract.Response(res)
}

// 对应 core/invokes包下的RegisterRule方法定义
func (s *SmartAudit) RegisterRules(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.RegisterRulesMain(args, context))
}

// 对应 core/invokes包下的RegisterProject方法定义
func (s *SmartAudit) RegisterProject(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.RegisterProjectMain(args, context))
}

// 对应 core/invokes包下的RegisterAuditee方法定义
func (s *SmartAudit) RegisterAuditee(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.RegisterAuditeeMain(args, context))
}

// 对应 core/invokes包下的AddEvent方法定义
func (s *SmartAudit) AddEvent(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.AddEventMain(args, context))
}

// 对应 core/invokes包下的GetAuditee方法定义
func (s *SmartAudit) GetAuditee(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.GetAuditeeMain(args, context))
}

// 对应 core/invokes包下的GetRule方法定义
func (s *SmartAudit) GetRules(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.GetRulesMain(args, context))
}

// 对应 core/invokes包下的GetProject方法定义
func (s *SmartAudit) GetProject(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.GetProjectMain(args, context))
}

// 对应 core/invokes包下的GetMaintainers方法定义
func (s *SmartAudit) GetMaintainers(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	return contract.Response(invokes.GetMaintainersMain(context))
}

// 对应 core/invokes包下的QueryEvents方法定义
func (s *SmartAudit) QueryEvents(ctx code.Context) code.Response {
	context := contract.NewContext(ctx)
	args := context.GetArgs()
	return contract.Response(invokes.QueryEventsMain(args, context))
}

// 审计合约主程序入口
func main() {
	driver.Serve(new(SmartAudit))
}
