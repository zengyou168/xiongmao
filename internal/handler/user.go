package handler

import (
	"github.com/gofiber/fiber/v2"
	"panda/internal/model"
	"panda/internal/service"
	"panda/pkg/respond"
)

// UserLogin 用户登录
func UserLogin(c *fiber.Ctx) error {

	var req model.UserLoginParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	return respond.OkData(c, service.Login(req))
}

// UserAdd 用户添加
func UserAdd(c *fiber.Ctx) error {

	var req model.UserAddParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	service.Add(req)

	return respond.Ok(c)
}
