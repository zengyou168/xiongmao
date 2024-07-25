package custom

import (
	"time"
)

type Model struct {
	ID        string     `gorm:"type:char(32);primaryKey" json:"id"`
	CreatedAt time.Time  `gorm:"type:timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time  `gorm:"type:timestamp" json:"updatedAt,omitempty"`
	DeletedAt *time.Time `gorm:"type:timestamp" json:"deletedAt,omitempty"`
}
