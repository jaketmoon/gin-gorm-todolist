package casbin

// AddBasedPolicies 添加基础策略
func (e *MyEnforce) AddBasedPolicies() {
	e.AddPolicy("user", "users", "read")
	e.AddPolicy("user", "users", "write")
}

// LinkUserWithPolicy 关联用户与策略
func (e *MyEnforce) LinkUserWithPolicy(name string) error {
	_, err := e.AddGroupingPolicy(name, "user")
	return err
}

// UnLinkUserWithPolicy 解除用户与策略的关联
func (e *MyEnforce) UnLinkUserWithPolicy(name string) error {
	_, err := e.RemoveGroupingPolicy(name, "user")
	return err
}

// CheckUserPolicyForRead 检查是否有操作权限
func (e *MyEnforce) CheckUserPolicyForRead(name, data, action string) bool {
	ok, _ := e.Enforce(name, data, action)
	return ok
}
