package invokes

import (
	"audit/orgnization"
	"bytes"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// 查询所有的审计运维人员信息
func QueryMaintainersMain(stub shim.ChaincodeStubInterface) peer.Response {
	countBuf, err := stub.GetState(orgnization.MaintainerCountKey)
	if err != nil {
		return shim.Error(fmt.Sprintf("获取审计当事人信息出错，详细信息：%s", err.Error()))
	}
	count, err := strconv.ParseUint(string(countBuf), 10, 32)

	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false
	buffer.WriteString(`{"result":[`)

	for i := uint64(0); i < count; i++ {
		mKey := orgnization.Maintainer{
			Member: &orgnization.Member{ID: uint32(i)}}.Key()
		mBuf, err := stub.GetState(mKey)
		if err != nil {
			return shim.Error(fmt.Sprintf(
				"获取审计维护人员信息出错，详细信息：%s", err.Error()))
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(mBuf))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString(`]}`)
	fmt.Printf("Query result: %s", buffer.String())

	// todo 将结果反序列化，转为json再显示
	return peer.Response{Payload: buffer.Bytes()}
}
