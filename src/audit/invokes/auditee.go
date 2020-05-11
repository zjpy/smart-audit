package invokes

import (
	"audit/orgnization"
	"audit/record"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegisterAuditeeMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	auditee, err := orgnization.AuditeeFromString(args)
	if err != nil {
		return shim.Error(fmt.Sprint("解析审计当事人失败，详细信息：", err))
	}

	if err = auditee.Validate(stub); err != nil {
		return shim.Error(fmt.Sprintf("审计当事人%s数据验证失败，详细信息：", err))
	}

	if err = record.Store(auditee, stub); err != nil {
		return shim.Error(fmt.Sprintf("审计当事人%s存储失败，详细信息：%s", auditee.Key(), err))
	}
	return shim.Success(nil)
}
