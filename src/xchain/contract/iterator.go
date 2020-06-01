package contract

import "github.com/xuperchain/xuperchain/core/contractsdk/go/code"

// 通过封装XuperChain内部的迭代器，提供统一的查询方法以支持不同链的调用
type IteratorImpl struct {
	it code.Iterator
}

// 判断是否还有更多的查询结果
func (i *IteratorImpl) HasNext() bool {
	return i.it.Next()
}

// 获取下一个查询结果
func (i *IteratorImpl) Next() (key string, value []byte, err error) {
	return string(i.it.Key()), i.it.Value(), i.it.Error()
}

// 关闭迭代器
func (i *IteratorImpl) Close() error {
	i.it.Close()
	return i.it.Error()
}

// 生成IteratorImpl实例
func NewIterator(it code.Iterator) *IteratorImpl {
	return &IteratorImpl{it: it}
}
