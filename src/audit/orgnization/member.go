package orgnization

// 一个组织中成员的基本结构
type Member struct {
	// 成员名称
	Name string

	// 成员唯一标识
	ID uint32

	// 成员公钥
	PK []byte
}