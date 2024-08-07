// Package custom 数据公共字段
package custom

import (
	"gorm.io/gorm"
	"time"
)

type Id struct {
	ID string `gorm:"type:char(32);primaryKey" json:"id,omitempty"`
}

// TenantId 租户ID
type TenantId struct {
	TenantId string `gorm:"type:char(32);index;comment:租户id" json:"tenantId,omitempty"`
}

type At struct {
	CreatedAt time.Time `gorm:"type:timestamp;comment:创建时间" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp;comment:更新时间" json:"updatedAt,omitempty"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;index;comment:删除时间" json:"updatedAt"`
}
