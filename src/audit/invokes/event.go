package invokes

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func AddEventMain(args []string, stub shim.ChaincodeStubInterface) peer.Response {
	// todo compete me
	return shim.Success(nil)
}
