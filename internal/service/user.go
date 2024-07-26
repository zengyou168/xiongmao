package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"panda/internal/model"
	"panda/pkg/db"
	"panda/pkg/respond"
	"panda/pkg/utils"
)

type user struct {
	model.User
}

// Login 用户登录
func Login(req model.UserLoginParam) model.UserLoginVO {

	var user user

	result := db.Gorm.Where("name = ?", req.Name).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(respond.Error("用户或密码错误"))
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(req.Pwd))

	if err != nil {
		panic(respond.Error("用户或密码错误..."))
	}

	userLoginVO := model.UserLoginVO{
		Name: user.Name,
		Pwd:  user.Pwd,
	}

	return userLoginVO
}

// CreateUser 创建新用户
func CreateUser(user1 model.User) (ModelVO, error) {

	localUser := user{User: user1}

	// 创建新用户

	if err := db.Gorm.Create(&localUser).Error; err != nil {
		return ModelVO{}, respond.Error("保存失败")
	}

	modelJSON := localUser.ModelVO11()

	return modelJSON, nil
}

type ModelVO struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

// ToDTO converts a Model to a ModelDTO with formatted dates
func (m user) ModelVO11() ModelVO {
	return ModelVO{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt.Format("2006-01-02"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02"),
		DeletedAt: func() string {
			if m.DeletedAt == nil {
				return ""
			}
			return m.DeletedAt.Format("2006-01-02")
		}(),
	}
}

func (user *user) BeforeCreate(tx *gorm.DB) (err error) {

	user.ID = utils.UUID()

	return
}
