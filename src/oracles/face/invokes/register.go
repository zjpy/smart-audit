package invokes

import (
	"core/contract"
	"oracles/face/service/dummy"
	_ "oracles/identify/service"
	"strconv"
)

var registration contract.Registration = initRegistration()

func RegisterMain(args []string, context contract.Context) *contract.Response {
	id, err := registration.Register(args)
	if err != nil {
		return contract.Error("规则注册错误，详细信息：" + err.Error())
	}

	return &contract.Response{
		Payload: []byte(strconv.FormatUint(uint64(id), 32)),
	}
}

func initRegistration() contract.Registration {
	// fixme 在真实商用环境下替换为完成好的service.FaceRegistration
	//return &service.FaceRegistration{}
	return &dummy.FaceRegistration{}
}