package handler

import (
	"github.com/gofiber/fiber/v2"
	"panda/pkg/respond"
)

func Add(c *fiber.Ctx) error {
	return respond.OkData(c, "addrrrrrrrrrrrrrrrrrrr")
}
