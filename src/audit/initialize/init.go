package initialize

import (
	"audit/orgnization"
	"audit/record"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strings"
)

const (
	maintainerSplitter = ","
)

func InitMain(stub shim.ChaincodeStubInterface) peer.Response {
	maintainers, err := parseMaintainers(stub.GetStringArgs(), stub)
	if err != nil {
		shim.Error(fmt.Sprint("解析合约运维人员出错，详细信息："))
	}

	for _, v := range maintainers {
		if err := record.StoreItem(v, stub); err != nil {
			return shim.Error(fmt.Sprintf("运维人员'%s'信息存储出错，详细信息：%s",
				v.Name, err.Error()))
		}
	}
	return shim.Success(nil)
}

func parseMaintainers(args []string, stub shim.ChaincodeStubInterface) (
	[]*orgnization.Maintainer, error) {
	var result []*orgnization.Maintainer
	for i, v := range args {
		m, err := parseSingleMaintainer(v, uint32(i), stub)
		if err != nil {
			return nil, err
		}

		result = append(result, m)
	}
	return result, nil
}

func parseSingleMaintainer(arg string, index uint32, stub shim.ChaincodeStubInterface) (
	*orgnization.Maintainer, error) {
	subArgs := strings.Split(arg, maintainerSplitter)
	return orgnization.MaintainerFromString(subArgs, index, stub)
}
