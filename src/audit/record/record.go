package record

import "github.com/hyperledger/fabric/core/chaincode/shim"

// 用于抽象一条记录的信息
type Item interface {
	Key() string
	Value() []byte
}

// 与智能合约记录相关的操作定义
func Store(item Item, stub shim.ChaincodeStubInterface) error {
	return stub.PutState(item.Key(), item.Value())
}
