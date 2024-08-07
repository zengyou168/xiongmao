package model

import (
	"gorm.io/gorm"
	"xiongmao/pkg/custom"
	"xiongmao/pkg/utils"
)

// Role 角色模型
type Role struct {
	custom.Id
	Name     string `gorm:"type:varchar(20);comment:名称" json:"name,omitempty"`
	Code     string `gorm:"type:varchar(20);comment:标识符" json:"code,omitempty"`
	Describe string `gorm:"type:varchar(100);comment:描述" json:"describe,omitempty"`
	custom.TenantId
	custom.At
	custom.DeletedAt
}

// RoleAddParam 角色添加请求参数
type RoleAddParam struct {
	Name     string `json:"name,omitempty"`
	Code     string `json:"code,omitempty"`
	Describe string `json:"describe,omitempty"`
}

// BeforeCreate GORM 钩子方法，生成 UUID
func (r *Role) BeforeCreate(db *gorm.DB) (err error) {
	r.ID = utils.UUID()
	return
}

// TableName 表名
func (Role) TableName() string {
	return "role"
}

// TableComment 方法返回表的注释
func (Role) TableComment() string {
	return "角色"
}
