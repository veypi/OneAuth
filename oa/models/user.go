package models

import (
	"gorm.io/gorm"
)

// salt for user user password gen aes code
// salt 32 hex / 16 byte / 128 bit
// code 64 hex / 32 byte / 256 bit
type User struct {
	BaseModel
	Username string `json:"username" gorm:"type:varchar(100);unique;default:not null" methods:"post,*patch,*list" parse:"json"`
	Nickname string `json:"nickname" gorm:"type:varchar(100)" methods:"*post,*patch,*list" parse:"json"`
	Icon     string `json:"icon" methods:"*post,*patch" parse:"json"`

	Email string `json:"email" gorm:"unique;type:varchar(50);default:null" methods:"*post,*patch,*list" parse:"json"`
	Phone string `json:"phone" gorm:"type:varchar(30);unique;default:null" methods:"*post,*patch,*list" parse:"json"`

	Status uint `json:"status" methods:"*patch,*list" parse:"json"`

	Salt string `json:"-" gorm:"type:varchar(32)"`
	Code string `json:"-" gorm:"type:varchar(64)" methods:"post" parse:"json"`
}

type UserRole struct {
	BaseModel
	UserID string `json:"user_id" methods:"get,list,post,patch,delete" parse:"path"`
	User   *User  `json:"-" gorm:"foreignKey:UserID;references:ID"`

	RoleID string `json:"role_id" methods:"post,delete" parse:"json"`
	Role   *Role  `json:"-" gorm:"foreignKey:RoleID;references:ID"`

	AppID string `json:"app_id" methods:"post,delete" parse:"json"`
	App   *App   `json:"-" gorm:"foreignKey:AppID;references:ID"`

	Status string `json:"status" methods:"post,*patch,*list" parse:"json"`
}

func (m *UserRole) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&Role{}).Where("id = ?", m.RoleID).Update("user_count", gorm.Expr("user_count + ?", 1)).Error
}
