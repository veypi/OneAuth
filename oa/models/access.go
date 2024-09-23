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
	App   *App   `json:"app" gorm:"foreignKey:ID;references:AppID"`

	UserID *string `json:"user_id" gorm:"index;type:varchar(32);default: null" methods:"post,list" parse:"json"`
	User   *User   `json:"user"`

	RoleID *string `json:"role_id" gorm:"index;type:varchar(32);default: null" methods:"post,list" parse:"json"`
	Role   *Role   `json:"role"`

	Name  string `json:"name" methods:"post,*list" parse:"json"`
	TID   string `json:"tid" methods:"post" parse:"json"`
	Level string `json:"level" methods:"post" parse:"json"`
}
