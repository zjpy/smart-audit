package contract

import (
	"core/contract"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
)

type ContextImpl struct {
	ctx code.Context
}

func (c *ContextImpl) GetFunctionName() string {
	return c.ctx.Caller()
}

func (c *ContextImpl) GetArgs() []string {
	// todo implement me
	return nil
}

func (c *ContextImpl) PutState(key string, value []byte) error {
	return c.ctx.PutObject([]byte(key), value)
}

func (c *ContextImpl) GetState(key string) ([]byte, error) {
	return c.ctx.GetObject([]byte(key))
}

func (c *ContextImpl) DeleteState(key string) error {
	return c.ctx.DeleteObject([]byte(key))
}

func (c *ContextImpl) GetStateByRange(startKey, endKey string) (contract.Iterator, error) {
	// todo 填入正确的limit
	return NewIterator(c.ctx.NewIterator([]byte(startKey), nil)), nil
}

func (c *ContextImpl) InvokeContract(name, function string, args [][]byte) contract.Response {
	// todo 填入正确的args
	res, err := c.ctx.Call("wasm", name, function, nil)
	if err != nil {
		return contract.Response{Err: err}
	}
	return contract.Response{
		Payload: res.Body,
	}
}

func NewContext(ctx code.Context) *ContextImpl {
	return &ContextImpl{ctx: ctx}
}
