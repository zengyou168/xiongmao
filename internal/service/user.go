package service

import (
	"gorm.io/gorm"
	"panda/internal/model"
	"panda/pkg/db"
	"panda/pkg/respond"
	"panda/pkg/utils"
)

type User struct {
	model.User
}

// CreateUser 创建新用户
func CreateUser(user model.User) (ModelVO, error) {

	localUser := User{User: user}

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
func (m User) ModelVO11() ModelVO {
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

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

	user.ID = utils.UUID()

	return
}
