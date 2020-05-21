package contract

import (
	"core/contract"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func Response(res *contract.Response) peer.Response {
	if res.Err != nil {
		return shim.Error(res.Err.Error())
	} else {
		return shim.Success(res.Payload)
	}
}
