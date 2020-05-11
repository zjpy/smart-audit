package invokes

import (
	"audit/project"
	"audit/record"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func AddEventMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	registration, err := project.RegistrationFromString(args, stub)
	if err != nil {
		return shim.Error(fmt.Sprint("合规事件登录失败，详细信息：", err))
	}

	if err = record.Store(registration, stub); err != nil {
		return shim.Error(fmt.Sprintf("合规事件%s存储失败，详细信息：%s",
			registration.Key(), err))
	}
	// 存储审计当事人规范对象所对应的存储条数
	if err = stub.PutState(project.GetRegistrationCountKey(registration.ID),
		[]byte(strconv.FormatUint(uint64(registration.Index+1), 10))); err != nil {
		return shim.Error(fmt.Sprintf("合规事件%s响应的索引值存储失败，详细信息：%s",
			registration.Key(), err))
	}
	return shim.Success(nil)
}
