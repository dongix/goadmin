package web

// UserQuery 用户查询参数
type UserQuery struct {
	ID int32 `uri:"id"`
}

// SmsCode 验证码
type SmsCode struct {
	Mobile string `uri:"mobile"`
}
