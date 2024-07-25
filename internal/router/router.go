// Package router 路由
package router

import (
    "github.com/gofiber/fiber/v2"
    "panda/config"
    "strconv"
)

func Init() {

    app := fiber.New()

    user(app)

    err := app.Listen(":" + strconv.Itoa(config.ServerVar.Port))

    if err != nil {
        panic(err)
    }
}
