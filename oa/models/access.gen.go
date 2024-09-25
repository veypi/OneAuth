package models

import "time"

type AccessList struct {
	CreatedAt *time.Time `json:"created_at"  parse:"query"`
	UpdatedAt *time.Time `json:"updated_at"  parse:"query"`
	AppID     string     `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
	UserID    *string    `json:"user_id" gorm:"index;type:varchar(32);default: null"  parse:"json"`
	RoleID    *string    `json:"role_id" gorm:"index;type:varchar(32);default: null"  parse:"json"`
	Name      *string    `json:"name"  parse:"json"`
}

type AccessPost struct {
	AppID  string  `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
	UserID *string `json:"user_id" gorm:"index;type:varchar(32);default: null"  parse:"json"`
	RoleID *string `json:"role_id" gorm:"index;type:varchar(32);default: null"  parse:"json"`
	Name   string  `json:"name"  parse:"json"`
	TID    string  `json:"tid"  parse:"json"`
	Level  uint    `json:"level"  parse:"json"`
}
