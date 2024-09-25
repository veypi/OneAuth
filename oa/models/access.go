//
// access.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-09-23 16:21
// Distributed under terms of the MIT license.
//

package models

type Access struct {
	BaseDate
	AppID string `json:"app_id" gorm:"index;type:varchar(32)" methods:"post,list" parse:"json"`
	App   *App   `json:"-" gorm:"foreignKey:AppID;references:ID"`

	UserID *string `json:"user_id" gorm:"index;type:varchar(32);default: null" methods:"post,list" parse:"json"`
	User   *User   `json:"-" gorm:"foreignKey:UserID;references:ID"`

	RoleID *string `json:"role_id" gorm:"index;type:varchar(32);default: null" methods:"post,list" parse:"json"`
	Role   *Role   `json:"-" gorm:"foreignKey:RoleID;references:ID"`

	Name string `json:"name" methods:"post,*list" parse:"json"`

	TID   string `json:"tid" methods:"post" parse:"json"`
	Level uint   `json:"level" methods:"post" parse:"json"`
}
