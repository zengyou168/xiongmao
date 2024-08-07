package model

import (
	"gorm.io/gorm"
	"xiongmao/pkg/custom"
	"xiongmao/pkg/utils"
)

// Admin 管理员模型
type Admin struct {
	custom.Id
	Name string `gorm:"type:varchar(20);index;comment:姓名" json:"name,omitempty"`
	Pwd  string `gorm:"type:varchar(60);comment:密码" json:"pwd,omitempty"`
	custom.TenantId
	custom.At
	custom.DeletedAt
}

// AdminLoginParam 管理员登录请求参数
type AdminLoginParam struct {
	Name string
	Pwd  string
}

// AdminLoginVO 管理员登录返回数据
type AdminLoginVO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

// AdminAddParam 管理员添加请求参数
type AdminAddParam struct {
	Name string
	Pwd  string
}

// BeforeCreate GORM 钩子方法，生成 UUID
func (r *Admin) BeforeCreate(db *gorm.DB) (err error) {
	r.ID = utils.UUID()
	return
}

// TableName 表名
func (Admin) TableName() string {
	return "admin"
}

// TableComment 方法返回表的注释
func (Admin) TableComment() string {
	return "管理员"
}
