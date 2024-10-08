package service

import (
	Jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
	"xiongmao/config"
	"xiongmao/internal/model"
	"xiongmao/pkg/db"
	"xiongmao/pkg/respond"
)

// AdminLogin 管理员登录
func AdminLogin(req model.AdminLoginParam) model.AdminLoginVO {

	admin := model.Admin{
		Name: req.Name,
	}

	r := db.Gorm.Where("name = ?", req.Name).First(&admin)

	if r.RowsAffected == 0 {
		panic(respond.Error("帐号或密码错误"))
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Pwd), []byte(req.Pwd))

	if err != nil {
		panic(respond.Error("帐号或密码错误"))
	}

	id := admin.ID
	name := admin.Name

	token := Jwt.NewWithClaims(Jwt.SigningMethodHS256, Jwt.MapClaims{
		"id":       id,
		"name":     name,
		"tenantId": "1",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenStr, err := token.SignedString(config.JwtKeyVar)

	if err != nil {
		panic(respond.Error("生成Token失败"))
	}

	adminLoginVO := model.AdminLoginVO{
		ID:    id,
		Name:  name,
		Token: tokenStr,
	}

	return adminLoginVO
}

// AdminAdd 管理员添加
func AdminAdd(req model.AdminAddParam) {

	tx := db.Gorm.Begin()

	pwd := req.Pwd

	if pwd == "" {
		panic(respond.Error("密码不能为空"))
	}

	name := req.Name

	admin := model.Admin{
		Name: req.Name,
	}

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
		tx.Rollback()
		panic(respond.Error("保存失败"))
	}

	tx.Commit()
}

type ModelVO struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

// ToDTO converts a Model to a ModelDTO with formatted dates
/*func (param admin) ModelVO11() ModelVO {
	return ModelVO{
		ID:        param.ID,
		Name:      param.Name,
		CreatedAt: param.CreatedAt.Format("2006-01-02"),
		UpdatedAt: param.UpdatedAt.Format("2006-01-02"),
		DeletedAt: func() string {
			if m.DeletedAt == nil {
				return ""
			}
			return m.DeletedAt.Format("2006-01-02")
		}(),
	}
}*/
