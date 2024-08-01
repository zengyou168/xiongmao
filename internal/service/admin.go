package service

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"panda/config"
	"panda/internal/model"
	"panda/pkg/db"
	"panda/pkg/respond"
	"panda/pkg/utils"
	"time"
)

type admin struct {
	model.Admin
}

// Login 管理员登录
func Login(req model.AdminLoginParam) model.AdminLoginVO {

	var admin admin

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"name": name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenStr, err := token.SignedString(config.CasbinSecretKey)

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

// Add 管理员添加
func Add(req model.AdminAddParam) {

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

	r = db.Gorm.Create(&admin)

	if r.Error != nil {
		panic(respond.Error("保存失败"))
	}
}

type ModelVO struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

// ToDTO converts a Model to a ModelDTO with formatted dates
func (m admin) ModelVO11() ModelVO {
	return ModelVO{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt.Format("2006-01-02"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02"),
		/*DeletedAt: func() string {
			if m.DeletedAt == nil {
				return ""
			}
			return m.DeletedAt.Format("2006-01-02")
		}(),*/
	}
}

func (admin *admin) BeforeCreate(tx *gorm.DB) (err error) {

	admin.ID = utils.UUID()

	return
}
