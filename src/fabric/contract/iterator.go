package contract

import "github.com/hyperledger/fabric/core/chaincode/shim"

// 将StateQueryIteratorInterface封装一层，提供统一的查询方法以支持不同链的调用
type IteratorImpl struct {
	raw shim.StateQueryIteratorInterface
}

// 判断是否还有更多的查询结果
func (i *IteratorImpl) HasNext() bool {
	return i.raw.HasNext()
}

// 获取下一个查询结果
func (i *IteratorImpl) Next() (key string, value []byte, err error) {
	return i.Next()
}

// 关闭迭代器
func (i *IteratorImpl) Close() error {
	return i.raw.Close()
}

// 生成IteratorImpl实例
func NewIterator(raw shim.StateQueryIteratorInterface) *IteratorImpl {
	return &IteratorImpl{raw: raw}
}
