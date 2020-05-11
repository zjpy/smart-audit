package project

// 用于记录一次信息登录
type Registration struct {
	// 登录事件相关信息
	EventInfo AuditeeSpecification

	// 登录发生的时间
	Timestamp uint32

	// 登录涉及的参数
	Params map[string]string
}
