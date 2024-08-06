package service

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
	"xiongmao/config"
	"xiongmao/internal/model"
	"xiongmao/pkg/db"
	"xiongmao/pkg/respond"
	"xiongmao/pkg/utils"
)

type admin struct {
	model.Admin
}

// AdminLogin 管理员登录
func AdminLogin(req model.AdminLoginParam) model.AdminLoginVO {

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

// AdminAdd 管理员添加
func AdminAdd(req model.AdminAddParam) {

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

	tx.Rollback()

	// 添加角色权限策略
	/*_, err = casbin.EnforcerVar.AddPolicy("admin", "/data1", "GET")
	  if err != nil {
	  	log.SugarVar.Error("Add policy error: %v", err)
	  }*/

	// 添加角色继承关系示例
	/*_, err = casbin.EnforcerVar.AddGroupingPolicy("alice", "admin")
	  if err != nil {
	  	log.SugarVar.Error("添加分组策略错误", err)
	  }

	  // 保存策略
	  err = casbin.EnforcerVar.SavePolicy()
	  if err != nil {
	  	log.SugarVar.Error("保存策略错误", err)
	  }*/
}

type ModelVO struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

// ToDTO converts a Model to a ModelDTO with formatted dates
func (param admin) ModelVO11() ModelVO {
	return ModelVO{
		ID:        param.ID,
		Name:      param.Name,
		CreatedAt: param.CreatedAt.Format("2006-01-02"),
		UpdatedAt: param.UpdatedAt.Format("2006-01-02"),
		/*DeletedAt: func() string {
			if m.DeletedAt == nil {
				return ""
			}
			return m.DeletedAt.Format("2006-01-02")
		}(),*/
	}
}

// BeforeCreate 请不到删除 db *gorm.DB，要不然重写不了id值，用于添加逻辑
func (param *admin) BeforeCreate(db *gorm.DB) (err error) {

	param.ID = utils.UUID()

	return
}
