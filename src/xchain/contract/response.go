package contract

import (
	"core/contract"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
)

func Response(res *contract.Response) code.Response {
	if res.Err != nil {
		return code.Error(res.Err)
	} else {
		return code.OK(res.Payload)
	}
}
