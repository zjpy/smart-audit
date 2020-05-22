package contract

import (
	"core/contract"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"sort"
	"strconv"
)

type ContextImpl struct {
	ctx code.Context
}

func (c *ContextImpl) GetFunctionName() string {
	return c.ctx.Caller()
}

func (c *ContextImpl) GetArgs() (rtn []string) {
	return mapToArray(c.ctx.Args())
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
	return NewIterator(
		c.ctx.NewIterator([]byte(startKey), []byte(endKey))), nil
}

func (c *ContextImpl) InvokeContract(name, function string, args []string) contract.Response {
	res, err := c.ctx.Call("wasm", name, function, arrayToMap(args))
	if err != nil {
		return contract.Response{Err: err}
	}
	return contract.Response{
		Payload: res.Body,
	}
}

func arrayToMap(args []string) map[string][]byte {
	argMap := make(map[string][]byte)
	for i, v := range args {
		argMap[strconv.FormatInt(int64(i), 32)] = []byte(v)
	}
	return argMap
}

func mapToArray(argsMap map[string][]byte) (rtn []string) {
	var keys []string
	for k := range argsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v := range keys {
		rtn = append(rtn, string(argsMap[v]))
	}
	return
}

func NewContext(ctx code.Context) *ContextImpl {
	return &ContextImpl{ctx: ctx}
}
