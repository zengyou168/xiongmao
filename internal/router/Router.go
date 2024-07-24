// Package router 路由
package router

import "github.com/gofiber/fiber/v2"

func Routers() *fiber.App {

	app := fiber.New()

	user(app)

	return app
}
