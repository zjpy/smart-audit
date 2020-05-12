package record

import "github.com/hyperledger/fabric/core/chaincode/shim"

// 用于抽象一条记录的信息
type Item interface {
	// Key值
	Key() string

	// 序列化存储除Key值外的所有数据
	Value() ([]byte, error)

	// 合法性检查，主要为数据本身的合法性相关
	Validate(stub shim.ChaincodeStubInterface) error
}

// 与智能合约记录相关的操作定义
func Store(item Item, stub shim.ChaincodeStubInterface) error {
	value, err := item.Value()
	if err != nil {
		return err
	}
	return stub.PutState(item.Key(), value)
}
