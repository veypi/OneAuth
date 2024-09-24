package models

import "time"

type AppGet struct {
	ID   string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@app_id"`
	Name string `json:"name"  parse:"json"`
}

type AppPatch struct {
	ID          string  `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@app_id"`
	Name        *string `json:"name"  parse:"json"`
	Icon        *string `json:"icon"  parse:"json"`
	Des         *string `json:"des"  parse:"json"`
	Participate *string `json:"participate" gorm:"default:auto"  parse:"json"`
	InitRoleID  *string `json:"init_role_id" gorm:"index;type:varchar(32)"  parse:"json"`
}

type AppDelete struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@app_id"`
}

type AppPost struct {
	Name        string `json:"name"  parse:"json"`
	Icon        string `json:"icon"  parse:"json"`
	Des         string `json:"des"  parse:"json"`
	Participate string `json:"participate" gorm:"default:auto"  parse:"json"`
}

type AppList struct {
	Name *string `json:"name"  parse:"json"`
}

type AppUserGet struct {
	ID     string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@app_user_id"`
	AppID  string `json:"app_id"  parse:"path"`
	UserID string `json:"user_id"  parse:"path"`
}

type AppUserPatch struct {
	ID     string  `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@app_user_id"`
	AppID  *string `json:"app_id"  parse:"path"`
	UserID *string `json:"user_id"  parse:"path"`
	Status *string `json:"status"  parse:"json"`
}

type AppUserDelete struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@app_user_id"`
}

type AppUserList struct {
	AppID  *string `json:"app_id"  parse:"path"`
	UserID *string `json:"user_id"  parse:"path"`
	Status *string `json:"status"  parse:"json"`
}

type AppUserPost struct {
	AppID  string `json:"app_id"  parse:"path"`
	UserID string `json:"user_id"  parse:"path"`
	Status string `json:"status"  parse:"json"`
}

type ResourceList struct {
	CreatedAt *time.Time `json:"created_at"  parse:"query"`
	UpdatedAt *time.Time `json:"updated_at"  parse:"query"`
	AppID     string     `json:"app_id" gorm:"primaryKey;type:varchar(32)"  parse:"json"`
}

type ResourcePost struct {
	AppID string `json:"app_id" gorm:"primaryKey;type:varchar(32)"  parse:"json"`
	Name  string `json:"name" gorm:"primaryKey"  parse:"json"`
	Des   string `json:"des"  parse:"json"`
}

type ResourceDelete struct {
	AppID string `json:"app_id" gorm:"primaryKey;type:varchar(32)"  parse:"json"`
	Name  string `json:"name" gorm:"primaryKey"  parse:"json"`
}
