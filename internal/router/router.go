// Package router 路由
package router

import (
	"github.com/gofiber/fiber/v2"
	"panda/config"
	"panda/pkg/respond"
	"strconv"
)

func Init() {

	app := fiber.New(fiber.Config{
		ErrorHandler: respond.ErrorHandler, // 设置全局错误处理函数
	})

	user(app)

	err := app.Listen(":" + strconv.Itoa(config.ServerVar.Port))

	if err != nil {
		panic(err)
	}
}
