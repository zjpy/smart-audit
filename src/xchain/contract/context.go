package contract

import (
	"core/contract"
	"github.com/xuperchain/xuperchain/core/common"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"sort"
	"strconv"
)

// 将XuperChain平台上下文封装，并实现统一的数据库使用等方法接口，以支持不同平台的调用
type ContextImpl struct {
	ctx code.Context
}

// 获取输入参数中的方法名，参数中第一个字段为方法名
func (c *ContextImpl) GetFunctionName() string {
	return c.ctx.Caller()
}

// 获取输入参数列表
func (c *ContextImpl) GetArgs() (rtn []string) {
	return mapToArray(c.ctx.Args())
}

// 存储数据
func (c *ContextImpl) PutState(key string, value []byte) error {
	return c.ctx.PutObject([]byte(key), value)
}

// 根据Key值查询之前存储的value值
func (c *ContextImpl) GetState(key string) ([]byte, error) {
	rtn, err := c.ctx.GetObject([]byte(key))
	if err == common.ErrKVNotFound {
		return nil, nil
	}
	return rtn, err
}

// 删除Key值对应记录
func (c *ContextImpl) DeleteState(key string) error {
	return c.ctx.DeleteObject([]byte(key))
}

// 根据起始Key值及结束Key值，查询所有满足条件的值，返回迭代器
func (c *ContextImpl) GetStateByRange(startKey, endKey string) (contract.Iterator, error) {
	return NewIterator(
		c.ctx.NewIterator([]byte(startKey), []byte(endKey))), nil
}

// 根据合约名、方法名及输入参数调用合约
func (c *ContextImpl) InvokeContract(name, function string, args []string) contract.Response {
	res, err := c.ctx.Call("wasm", name, function, arrayToMap(args))
	if err != nil {
		return contract.Response{Err: err}
	}
	return contract.Response{
		Payload: res.Body,
	}
}

// 工具方法，用于将一个字符数组转换成字典
func arrayToMap(args []string) map[string][]byte {
	argMap := make(map[string][]byte)
	for i, v := range args {
		argMap[strconv.FormatInt(int64(i), 32)] = []byte(v)
	}
	return argMap
}

// 工具方法，用于将一个字典对象转换成数组
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

// 生成ContextImpl实例
func NewContext(ctx code.Context) *ContextImpl {
	return &ContextImpl{ctx: ctx}
}
