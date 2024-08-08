package handler

import (
	"github.com/gofiber/fiber/v2"
	Jwt "github.com/golang-jwt/jwt/v5"
	"xiongmao/config"
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

// RoleAdd 角色添加
func RoleAdd(c *fiber.Ctx) error {

	var req model.RoleAddParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	//   tenantId := custom.TenantIdStr(c.Locals("user").(Jwt.MapClaims)["tenantId"].(string))
	tenantId := config.UserJwt["tenantId"].(string)

	panic(tenantId)

	//log.SugarVar.Error("ddddddddddddddddddddd：", tenantId)

	service.RoleAdd(req)

	return respond.Ok(c)
}

// RoleEdit 角色编辑
func RoleEdit(c *fiber.Ctx) error {

	var req model.RoleEditParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error("参数错误"))
	}

	tenantId := c.Locals("user").(Jwt.MapClaims)["tenantId"].(string)

	service.RoleEdit(req, tenantId)

	return respond.Ok(c)
}
