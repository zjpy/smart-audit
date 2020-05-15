package record

import "github.com/hyperledger/fabric/core/chaincode/shim"

// 用于抽象一条记录的信息
type Item interface {
	// Key值
	Key() string

	// 序列化存储除Key值外的所有数据
	Value() ([]byte, error)

	// 合法性检查，主要为数据本身的合法性相关
	Validate() error
}

// 存储一条记录，将接口中的Key、Value值一起存储到区块链中
func StoreItem(item Item, stub shim.ChaincodeStubInterface) error {
	value, err := item.Value()
	if err != nil {
		return err
	}
	return stub.PutState(item.Key(), value)
}
