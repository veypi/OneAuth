package models

import "github.com/veypi/OneAuth/oalib"

type UserRole struct {
	BaseModel
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type Role struct {
	BaseModel
	AppUUID string `json:"app_uuid" gorm:"size:32"`
	App     *App   `json:"app" gorm:"association_foreignkey:UUID"`
	Name    string `json:"name"`
	// 角色标签
	Tag   string  `json:"tag" gorm:"default:''"`
	Users []*User `json:"users" gorm:"many2many:user_roles;"`
	// 具体权限
	Auths     []*Auth `json:"auths" gorm:"foreignkey:RoleID;references:ID"`
	UserCount uint    `json:"user_count"`
}

// Auth 资源权限
type Auth struct {
	BaseModel
	// 该权限作用的应用
	AppUUID string `json:"app_uuid" gorm:"size:32"`
	App     *App   `json:"app" gorm:"association_foreignkey:UUID"`
	// 权限绑定只能绑定一个
	RoleID *uint `json:"role_id" gorm:""`
	Role   *Role `json:"role"`
	UserID *uint `json:"user_id"`
	User   *User `json:"user"`
	// 资源id
	ResourceID uint      `json:"resource_id" gorm:"not null"`
	Resource   *Resource `json:"resource"`
	// resource_name 用于其他系统方便区分权限的名字
	RID string `json:"rid" gorm:""`
	// 具体某个资源的id
	RUID  string          `json:"ruid"`
	Level oalib.AuthLevel `json:"level"`
}

type Resource struct {
	BaseModel
	AppUUID string `json:"app_uuid"  gorm:"size:32"`
	App     *App   `json:"app" gorm:"association_foreignkey:UUID"`
	Name    string `json:"name"`
	// 权限标签
	Tag string `json:"tag"`
	Des string `json:"des"`
}
