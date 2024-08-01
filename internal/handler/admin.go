package handler

import (
	"github.com/gofiber/fiber/v2"
	"panda/internal/model"
	"panda/internal/service"
	"panda/pkg/respond"
)

// AdminLogin 管理员登录
func AdminLogin(c *fiber.Ctx) error {

	var req model.AdminLoginParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	return respond.OkData(c, service.Login(req))
}

// AdminAdd 管理员添加
func AdminAdd(c *fiber.Ctx) error {

	var req model.AdminAddParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	service.Add(req)

	return respond.Ok(c)
}
