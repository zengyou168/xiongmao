package handler

import (
	"github.com/gofiber/fiber/v2"
	"panda/internal/model"
	"panda/internal/service"
	"panda/pkg/respond"
)

func UserLogin(c *fiber.Ctx) error {

	// 定义请求结构
	var req model.UserLoginParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	// 生成盐值并哈希密码
	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)

	return respond.OkData(c, service.Login(req))

}
