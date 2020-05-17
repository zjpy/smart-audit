package invokes

import (
	"audit/contract"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	_ "identify/service"
	"identify/service/dummy"
	"strconv"
)

const (
	RegisterFunctionName = "register"
)

var registration contract.Registration = initRegistration()

func RegisterMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	id, err := registration.Register(args)
	if err != nil {
		return shim.Error("规则注册错误，详细信息：" + err.Error())
	}

	return shim.Success([]byte(strconv.FormatUint(uint64(id), 32)))
}

func initRegistration() contract.Registration {
	// fixme 在真实商用环境下替换为完成好的service.EntityIdentifyRegistration
	//return &service.EntityIdentifyRegistration{}
	return &dummy.EntityIdentifyRegistration{}
}
