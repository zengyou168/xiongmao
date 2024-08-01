package respond

import (
	"github.com/gofiber/fiber/v2"
)

// Response 定义一个标准响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// Ok 函数返回一个成功响应
func Ok(c *fiber.Ctx) error {
	return respond(c, 0, "", nil)
}

// OkData Ok 函数返回一个成功响应
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
