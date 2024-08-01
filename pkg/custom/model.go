package custom

import (
	"time"
)

type Id struct {
	ID string `gorm:"type:char(32);primaryKey" json:"id"`
}

type At struct {
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt,omitempty"`
}
