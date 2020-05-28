package contract

import (
	"core/contract"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ContextImpl struct {
	stub shim.ChaincodeStubInterface
}

func (c *ContextImpl) GetFunctionName() string {
	fn, _ := c.stub.GetFunctionAndParameters()
	return fn
}

func (c *ContextImpl) GetArgs() []string {
	_, args := c.stub.GetFunctionAndParameters()
	return args
}

func (c *ContextImpl) PutState(key string, value []byte) error {
	return c.stub.PutState(key, value)
}

func (c *ContextImpl) GetState(key string) ([]byte, error) {
	return c.stub.GetState(key)
}

func (c *ContextImpl) DeleteState(key string) error {
	return c.stub.DelState(key)
}

func (c *ContextImpl) GetStateByRange(startKey, endKey string) (contract.Iterator, error) {
	raw, err := c.stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	return NewIterator(raw), nil
}

func (c *ContextImpl) InvokeContract(name, function string, args []string) contract.Response {
	fabricArgs := [][]byte{
		[]byte(function),
	}
	for _, v := range args {
		fabricArgs = append(fabricArgs, []byte(v))
	}
	// 这里channel参数为空则默认会发送到当前合约所在channel上
	rtn := c.stub.InvokeChaincode(name, fabricArgs, "")
	response := contract.Response{
		Payload: rtn.Payload,
	}
	if rtn.Status != shim.OK {
		response.Err = errors.New(rtn.Message)
	}
	return response
}

func NewContext(stub shim.ChaincodeStubInterface) *ContextImpl {
	return &ContextImpl{
		stub: stub,
	}
}
