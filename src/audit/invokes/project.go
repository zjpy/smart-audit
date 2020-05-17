package invokes

import (
	"audit/project"
	"audit/record"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// 注册项目，返回项目ID
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

// 根据项目ID获取项目信息
func GetProjectMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	if len(args) == 0 {
		return shim.Error("查询失败，需要提供项目ID")
	}
	projectID, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return shim.Error(fmt.Sprintf("解析项目ID出错，详细信息：%s", err.Error()))
	}
	projectKey := project.Project{ID: uint32(projectID)}.Key()
	ruleBuf, err := stub.GetState(projectKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取项目信息出错，详细信息：%s", err.Error()))
	}
	// todo 将结果反序列化，转为json再显示
	return peer.Response{Payload: ruleBuf}
}
