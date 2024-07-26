package handler

import (
	"github.com/gofiber/fiber/v2"
	"panda/pkg/respond"
)

func UserLogin(c *fiber.Ctx) error {

	// 定义请求结构
	/*var req model.User

	  // 解析请求体中的 JSON 数据
	  if err := c.BodyParser(&req); err != nil {
	  	// 解析失败，返回错误响应
	  	return respond.Error("解析请求体中的 JSON 数据失败")
	  }

	  users, _ := service.CreateUser(req)*/

	return respond.OkData(c, c.Locals("user"))

}
