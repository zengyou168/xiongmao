package handler

import (
	"github.com/gofiber/fiber/v2"
	"xiongmao/internal/model"
	"xiongmao/internal/service"
	"xiongmao/pkg/respond"
)

// AdminLogin 管理员登录
func AdminLogin(c *fiber.Ctx) error {

	var req model.AdminLoginParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	return respond.OkData(c, service.AdminLogin(req))
}

// AdminAdd 管理员添加
func AdminAdd(c *fiber.Ctx) error {

	var req model.AdminAddParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	service.AdminAdd(req)

	return respond.Ok(c)
}

// RuleAdd 角色添加
func RuleAdd(c *fiber.Ctx) error {
	return respond.Ok(c)
}
