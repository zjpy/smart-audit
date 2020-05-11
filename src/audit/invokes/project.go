package invokes

import (
	"audit/project"
	"audit/record"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegisterProjectMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	p, err := project.FromStrings(args)
	if err != nil {
		return shim.Error(fmt.Sprint("解析审计业务失败，详细信息：", err))
	}

	if err = p.Validate(stub); err != nil {
		return shim.Error(fmt.Sprintf("审计业务%s数据验证失败，详细信息：%s", p.Key(), err))
	}

	if err = record.Store(p, stub); err != nil {
		return shim.Error(fmt.Sprintf("审计业务%s存储失败，详细信息：%s", p.Key(), err))
	}
	return shim.Success(nil)
}
