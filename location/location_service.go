package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type LocationService struct {
}

func (s *LocationService) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *LocationService) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(LocationService)); err != nil {
		fmt.Printf("智能合约启动出错，详细信息：%s", err)
	}
}
