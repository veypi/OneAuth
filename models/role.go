package models

import (
	"github.com/veypi/OneAuth/oalib"
)

type UserRole struct {
	BaseModel
	UserID uint
	RoleID uint
}

type Role struct {
	BaseModel
	AppUUID string `gorm:"size:32"`
	App     *App   `gorm:"association_foreignkey:UUID"`
	Name    string
	// 角色标签
	Tag   string  `gorm:"default:''"`
	Users []*User `gorm:"many2many:UserRoles;"`
	// 具体权限
	Auths     []*Auth `gorm:"foreignkey:RoleID;references:ID"`
	UserCount uint
}

// Auth 资源权限
type Auth struct {
	BaseModel
	// 该权限作用的应用
	AppUUID string `gorm:"size:32"`
	App     *App   `gorm:"association_foreignkey:UUID"`
	// 权限绑定只能绑定一个
	RoleID *uint
	Role   *Role
	UserID *uint
	User   *User
	// 资源id
	ResourceID uint `gorm:"not null"`
	Resource   *Resource
	// resource_name 用于其他系统方便区分权限的名字
	RID string `gorm:""`
	// 具体某个资源的id
	RUID  string
	Level oalib.AuthLevel
}

type Resource struct {
	BaseModel
	AppUUID string ` gorm:"size:32"`
	App     *App   `gorm:"association_foreignkey:UUID"`
	Name    string
	Des     string
}
