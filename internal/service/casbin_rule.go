package service

import (
	"golang.org/x/crypto/bcrypt"
	"xiongmao/internal/model"
	"xiongmao/pkg/db"
	"xiongmao/pkg/respond"
)

// RuleAdd 角色添加
func RuleAdd(req model.AdminAddParam) {

	tx := db.Gorm.Begin()

	pwd := req.Pwd

	if pwd == "" {
		panic(respond.Error("密码不能为空"))
	}

	name := req.Name

	var admin admin

	r := db.Gorm.Where("name = ?", name).First(&admin)

	if r.RowsAffected != 0 {
		panic(respond.Error("帐号已存在"))
	}

	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		panic(respond.Error("加密密码失败"))
	}

	pwd = string(pwdBytes)

	admin.Name = name
	admin.Pwd = pwd

	r = tx.Create(&admin)

	if r.Error != nil {
		panic(respond.Error("保存失败"))
	}

}
