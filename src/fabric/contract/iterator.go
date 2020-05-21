package contract

import "github.com/hyperledger/fabric/core/chaincode/shim"

type IteratorImpl struct {
	raw shim.StateQueryIteratorInterface
}

func (i *IteratorImpl) HasNext() bool {
	return i.raw.HasNext()
}

func (i *IteratorImpl) Next() (key string, value []byte, err error) {
	return i.Next()
}

func (i *IteratorImpl) Close() error {
	return i.raw.Close()
}

func NewIterator(raw shim.StateQueryIteratorInterface) *IteratorImpl {
	return &IteratorImpl{raw: raw}
}
