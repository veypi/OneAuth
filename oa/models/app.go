package models

import ()

type App struct {
	BaseModel
	Name        string `json:"name" methods:"get,post,put,*patch,*list" parse:"json"`
	Icon        string `json:"icon" methods:"post,put,*patch" parse:"json"`
	Des         string `json:"des" methods:"post,put,*patch" parse:"json"`
	Participate string `json:"participate" gorm:"default:auto" methods:"post,put,*patch" parse:"json"`
	InitRoleID  string `json:"init_role_id" gorm:"index;type:varchar(32)" methods:"put,*patch" parse:"json"`
	InitRole    *Role  `json:"init_role" gorm:"foreignKey:ID;references:InitRoleID"`
	InitUrl     string `json:"init_url"`
	UserCount   uint   `json:"user_count"`
	Key         string `json:"-"`
}

type AppUser struct {
	BaseModel
	AppID  string `json:"app_id" methods:"get,*list,post,*patch" parse:"path"`
	App    *App   `json:"app"`
	UserID string `json:"user_id" methods:"get,*list,post,*patch" parse:"path"`
	User   *User  `json:"user"`
	Status string `json:"status" methods:"post,put,*patch,*list" parse:"json"`
}

type Resource struct {
	BaseDate
	AppID string `json:"app_id" gorm:"primaryKey;type:varchar(32)" methods:"post,list,delete" parse:"json"`
	App   *App   `json:"app" gorm:"foreignKey:ID;references:AppID"`
	Name  string `json:"name" gorm:"primaryKey" methods:"post,delete" parse:"json"`
	Des   string `json:"des" methods:"post" parse:"json"`
}
