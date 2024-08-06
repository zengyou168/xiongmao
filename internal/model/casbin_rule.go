package model

// CasbinRule 角色模型
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Ptype string `gorm:"type:varchar(100)" json:"ptype,omitempty"`
	V0    string `gorm:"type:varchar(100)" json:"v0,omitempty"`
	V1    string `gorm:"type:varchar(100)" json:"v1,omitempty"`
	V2    string `gorm:"type:varchar(100)" json:"v2,omitempty"`
	V3    string `gorm:"type:varchar(100)" json:"v3,omitempty"`
	V4    string `gorm:"type:varchar(100)" json:"v4,omitempty"`
	V5    string `gorm:"type:varchar(100)" json:"v5,omitempty"`
}

// CasbinRuleAdd 角色添加
type CasbinRuleAdd struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Ptype string `gorm:"type:varchar(100)" json:"ptype,omitempty"`
	V0    string `gorm:"type:varchar(100)" json:"v0,omitempty"`
	V1    string `gorm:"type:varchar(100)" json:"v1,omitempty"`
	V2    string `gorm:"type:varchar(100)" json:"v2,omitempty"`
	V3    string `gorm:"type:varchar(100)" json:"v3,omitempty"`
	V4    string `gorm:"type:varchar(100)" json:"v4,omitempty"`
	V5    string `gorm:"type:varchar(100)" json:"v5,omitempty"`
}
