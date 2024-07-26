package model

import "panda/pkg/custom"

// User 用户模型
type User struct {
	custom.Model
	Name string `gorm:"type:varchar(100);comment:姓名" json:"name,omitempty"`
	Pwd  string `gorm:"type:varchar(100);comment:密码" json:"pwd,omitempty"`
}

// UserLoginParam 用户登录请求参数
type UserLoginParam struct {
	Name string
	Pwd  string
}

// UserLoginVO 用户登录返回数据
type UserLoginVO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserCreate 用于接收创建用户请求的数据
type UserCreate struct {
	Username string `gorm:"type:varchar(100);comment:用户姓名"`
	Password string
}

// 定义请求结构体
type RequestUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Custom JSON representation of User
type UserDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt,omitempty"`
}
