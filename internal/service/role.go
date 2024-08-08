package service

import (
	"xiongmao/internal/model"
	"xiongmao/pkg/db"
	"xiongmao/pkg/respond"
)

// RoleAdd 角色添加
func RoleAdd(req model.RoleAddParam) {

	role := model.Role{
		Name:     req.Name,
		Code:     req.Code,
		Describe: req.Describe,
	}

	r := db.Gorm.Create(&role)

	if r.Error != nil {
		panic(respond.Error("保存失败"))
	}
}

// RoleEdit 角色编辑
func RoleEdit(req model.RoleEditParam, tenantId string) {

	ID := req.ID

	if ID == "" {
		panic(respond.Error("ID不能为空，编辑失败"))
	}

	if ID == "1" {
		panic(respond.Error("系统默认角色，编辑失败"))
	}

	role := model.Role{
		Name:     req.Name,
		Code:     req.Code,
		Describe: req.Describe,
		///TenantId: router.UserJwtVar.Method(),
	}

	r := db.Gorm.Where("id = ?", ID).Updates(&role)

	if r.Error != nil {
		panic(respond.Error("保存失败"))
	}
}
