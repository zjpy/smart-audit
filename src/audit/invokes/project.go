package invokes

import (
	"audit/project"
	"audit/record"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegisterProjectMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	p, err := project.FromStrings(args, stub)
	if err != nil {
		return shim.Error(fmt.Sprint("解析审计业务失败，详细信息：", err))
	}

	if err = p.Validate(); err != nil {
		return shim.Error(fmt.Sprintf("审计业务%s数据验证失败，详细信息：%s", p.Key(), err))
	}

	if err = record.StoreItem(p, stub); err != nil {
		return shim.Error(fmt.Sprintf("审计业务%s存储失败，详细信息：%s", p.Key(), err))
	}
	if err = record.StoreCount(p, stub); err != nil {
		return shim.Error(fmt.Sprintf("审计业务%s相应的索引值存储失败，详细信息：%s",
			p.Key(), err))
	}
	return shim.Success(nil)
}
