package exception

import "fmt"

// ErrorCode 自定义错误结构体
type ErrorCode struct {
	Code string
	Msg  string
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("code: %s, msg: %s", e.Code, e.Msg)
}
