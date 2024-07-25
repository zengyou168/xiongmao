package respond

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

// Response 定义一个标准响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const ()

// Ok 函数返回一个成功响应
func Ok(c *fiber.Ctx) error {
	return respond(c, 0, "", nil)
}

// Ok 函数返回一个成功响应
func OkData(c *fiber.Ctx, data interface{}) error {
	return respond(c, 0, "", data)
}

// 创建一个辅助函数来生成响应
func respond(c *fiber.Ctx, code int, msg string, data interface{}) error {
	return c.JSON(Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 定义自定义错误类型
type APIError struct {
	Code int
	Msg  string
}

// 实现 Error 接口
func (e *APIError) Error() string {
	return e.Msg
}

// 错误处理中间件
func ErrorHandler(c *fiber.Ctx, err error) error {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return respond(c, apiErr.Code, apiErr.Msg, nil)
	}
	// 默认错误处理
	return respond(c, 500, err.Error(), nil)
}
