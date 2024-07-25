package respond

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

// ErrorData 定义自定义错误类型
type ErrorData struct {
	Code int
	Msg  string
}

func Error(msg string) *ErrorData {
	return &ErrorData{
		Code: 500, // 设置默认错误码
		Msg:  msg,
	}
}

func ErrorCode(code int, msg string) *ErrorData {
	return &ErrorData{
		Code: code,
		Msg:  msg,
	}
}

// 实现 error 接口的方法
func (e *ErrorData) Error() string {
	return e.Msg
}

// ErrorHandler 错误处理中间件
func ErrorHandler(c *fiber.Ctx, err error) error {

	var data *ErrorData

	if errors.As(err, &data) {
		return respond(c, data.Code, data.Msg, nil)
	}

	// 默认错误处理
	return respond(c, 500, err.Error(), nil)
}