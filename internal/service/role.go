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
