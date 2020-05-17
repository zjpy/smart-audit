package invokes

import (
	"audit/common"
	"audit/orgnization"
	"audit/record"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// 注册审计当事人，反回值为审计当时人ID
func RegisterAuditeeMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	auditee, err := orgnization.AuditeeFromString(args, stub)
	if err != nil {
		return shim.Error(fmt.Sprint("解析审计当事人失败，详细信息：", err))
	}

	if err = auditee.Validate(); err != nil {
		return shim.Error(fmt.Sprintf("审计当事人%s数据验证失败，详细信息：", err))
	}

	if err = record.StoreItem(auditee, stub); err != nil {
		return shim.Error(fmt.Sprintf("审计当事人%s存储失败，详细信息：%s", auditee.Key(), err))
	}
	if err = record.StoreCount(auditee, stub); err != nil {
		return shim.Error(fmt.Sprintf("审计当事人%s相应的索引值存储失败，详细信息：%s",
			auditee.Key(), err))
	}

	return shim.Success(common.Uint32ToBytes(auditee.ID))
}

// 根据审计当事人ID获取当事人信息
func GetAuditeeMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	if len(args) == 0 {
		return shim.Error("查询失败，需要提供审计当事人ID")
	}
	auditeeID, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return shim.Error(fmt.Sprintf("解析审计当事人ID出错，详细信息：%s", err.Error()))
	}

	auditeeKey := orgnization.Auditee{
		Member: &orgnization.Member{ID: uint32(auditeeID)}}.Key()
	ruleBuf, err := stub.GetState(auditeeKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取审计当事人信息出错，详细信息：%s", err.Error()))
	}
	// todo 将结果反序列化，转为json再显示
	return peer.Response{Payload: ruleBuf}
}
