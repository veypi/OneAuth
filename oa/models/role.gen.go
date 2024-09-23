package models

import ()

type RoleDelete struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@role_id"`
}

type RoleGet struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@role_id"`
}

type RoleList struct {
	Name *string `json:"name"  parse:"json"`
}

type RolePatch struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@role_id"`

	Name *string `json:"name"  parse:"json"`

	Des *string `json:"des"  parse:"json"`

	AppID *string `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
}

type RolePost struct {
	Name string `json:"name"  parse:"json"`

	Des string `json:"des"  parse:"json"`

	AppID string `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
}

type RolePut struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@role_id"`

	Name string `json:"name"  parse:"json"`

	Des string `json:"des"  parse:"json"`

	AppID string `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
}
