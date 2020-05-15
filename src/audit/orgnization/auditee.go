package orgnization

import (
	"audit/record"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	auditeePrefix   = "auditee-"
	auditeeCountKey = "auditee-count"
)

// 审计人员
type Auditee struct {
	*Member
}

func (a *Auditee) Key() string {
	return auditeePrefix + a.Member.Key()
}

func (a *Auditee) CountKey() string {
	return auditeeCountKey
}

func (a *Auditee) GetCount() uint32 {
	return a.ID + 1
}

func AuditeeFromString(args []string,
	stub shim.ChaincodeStubInterface) (auditee *Auditee, err error) {
	var mem *Member
	if mem, err = MemberFromString(args, stub); err != nil {
		return
	}

	auditee = &Auditee{
		Member: mem,
	}
	auditee.ID, err = record.GetRecordCount(auditeeCountKey, stub)
	return
}
