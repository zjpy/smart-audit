package invokes

import (
	"audit/contract"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	_ "time/service"
	"time/service/dummy"
)

var validation contract.Validation = initValidation()

func ValidateTimeMain(args []string) peer.Response {
	if len(args) == 0 {
		return shim.Error("缺少规则ID")
	}

	value, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return shim.Error("解析规则ID出错，详细信息：" + err.Error())
	}
	if err = validation.Validate(contract.ServiceRuleID(value), args[1:]); err != nil {
		return shim.Error("验证错误，详细信息：" + err.Error())
	}

	return shim.Success(nil)
}

func initValidation() contract.Validation {
	// fixme 在真实商用环境下替换为完成好的service.TimeValidation
	//return &service.TimeValidation{}
	return &dummy.TimeValidation{}
}
