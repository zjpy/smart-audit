package contract

import "github.com/xuperchain/xuperchain/core/contractsdk/go/code"

type IteratorImpl struct {
	it code.Iterator
}

func (i *IteratorImpl) HasNext() bool {
	return i.it.Next()
}

func (i *IteratorImpl) Next() (key string, value []byte, err error) {
	return string(i.it.Key()), i.it.Value(), i.it.Error()
}

func (i *IteratorImpl) Close() error {
	i.it.Close()
	return i.it.Error()
}

func NewIterator(it code.Iterator) *IteratorImpl {
	return &IteratorImpl{it: it}
}
