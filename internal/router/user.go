// Package router 用户路由
package router

import (
	"github.com/gofiber/fiber/v2"
	"panda/internal/handler"
)

func user(app *fiber.App) {

	userRouter := app.Group("user")
	{
		userRouter.Post("login", handler.UserLogin) // 用户登录
		userRouter.Post("add", handler.UserAdd)     // 用户添加
	}

}
