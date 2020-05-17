package orgnization

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	maintainerPrefix   = "maintainer-"

	MaintainerCountKey = "maintainer-count"
)

// 合约运维人员
type Maintainer struct {
	*Member
}

func (m *Maintainer) Key() string {
	return maintainerPrefix + m.Member.Key()
}

func (m *Maintainer) CountKey() string {
	return MaintainerCountKey
}

func (m *Maintainer) GetCount() uint32 {
	return m.ID + 1
}

func MaintainerFromString(args []string, index uint32,
	stub shim.ChaincodeStubInterface) (mt *Maintainer, err error) {
	var mem *Member
	if mem, err = MemberFromString(args, stub); err != nil {
		return
	}

	mem.ID = index
	mt = &Maintainer{
		Member: mem,
	}
	return
}
