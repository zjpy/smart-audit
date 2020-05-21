package orgnization

import (
	"core/contract"
	"core/record"
)

const (
	AuditeePrefix   = "auditee-"
	AuditeeCountKey = "auditee-count"
)

// 审计人员
type Auditee struct {
	*Member
}

func (a *Auditee) Key() string {
	return AuditeePrefix + a.Member.Key()
}

func (a *Auditee) CountKey() string {
	return AuditeeCountKey
}

func (a *Auditee) GetCount() uint32 {
	return a.ID + 1
}

func AuditeeFromString(args []string,
	context contract.Context) (auditee *Auditee, err error) {
	var mem *Member
	if mem, err = MemberFromString(args, context); err != nil {
		return
	}

	auditee = &Auditee{
		Member: mem,
	}
	auditee.ID, err = record.GetRecordCount(AuditeeCountKey, context)
	return
}
