package orgnization

import (
	"core/contract"
)

const (
	MaintainerPrefix = "maintainer-"

	MaintainerCountKey = "maintainer-count"
)

// 合约运维人员
type Maintainer struct {
	*Member
}

func (m *Maintainer) Key() string {
	return MaintainerPrefix + m.Member.Key()
}

func (m *Maintainer) CountKey() string {
	return MaintainerCountKey
}

func (m *Maintainer) GetCount() uint32 {
	return m.ID + 1
}

func MaintainerFromString(args []string, index uint32,
	context contract.Context) (mt *Maintainer, err error) {
	var mem *Member
	if mem, err = MemberFromString(args, context); err != nil {
		return
	}

	mem.ID = index
	mt = &Maintainer{
		Member: mem,
	}
	return
}
