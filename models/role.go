package models

import (
	"OneAuth/cfg"
	"github.com/veypi/utils/log"
)

var GlobalRoles = make(map[uint]*Role)

func SyncGlobalRoles() {
	roles := make([]*Role, 0, 10)
	err := cfg.DB().Preload("Auths").Find(&roles).Error
	if err != nil {
		log.Warn().Msgf("sync global roles error: %s", err.Error())
		return
	}
	for _, r := range roles {
		GlobalRoles[r.ID] = r
	}
}

type UserRole struct {
	BaseModel
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type RoleAuth struct {
	BaseModel
	RoleID uint `json:"role_id"`
	AuthID uint `json:"auth_id"`
}

type Role struct {
	BaseModel
	Name string `json:"name"`
	// 角色类型
	// 0: 系统角色   1: 用户角色
	Category uint `json:"category" gorm:"default:0"`
	// 角色标签
	Tag   string  `json:"tag" gorm:"default:''"`
	Users []*User `json:"users" gorm:"many2many:user_role;"`
	// 具体权限
	Auths    []*Auth `json:"auths" gorm:"many2many:role_auth;"`
	IsUnique bool    `json:"is_unique" gorm:"default:false"`
}

func (r Role) CheckAuth(name string, tags ...string) AuthLevel {
	res := AuthNone
	tag := ""
	if len(tags) > 0 {
		tag = tags[0]
	}
	for _, a := range r.Auths {
		if a.Name == "admin" && a.Tag == "" || (a.Name == "admin" && a.Tag == tag) || (a.Name == name && a.Tag == tag) {
			if a.Level > res {
				res = a.Level
			}
		}
	}
	return res
}

type AuthLevel uint

const (
	AuthNone   AuthLevel = 0
	AuthRead   AuthLevel = 1
	AuthCreate AuthLevel = 2
	AuthUpdate AuthLevel = 3
	AuthDelete AuthLevel = 4
)

func (a AuthLevel) CanRead() bool {
	return a >= AuthRead
}

func (a AuthLevel) CanCreate() bool {
	return a >= AuthCreate
}

func (a AuthLevel) CanUpdate() bool {
	return a >= AuthUpdate
}

func (a AuthLevel) CanDelete() bool {
	return a >= AuthDelete
}

func (a AuthLevel) CanDoAny() bool {
	return a >= AuthDelete
}

// 资源权限

type Auth struct {
	BaseModel
	Name  string `json:"name"`
	AppID uint   `json:"app_id"`
	// 权限标签
	Tag string `json:"tag"`
	// 权限等级 0 相当于没有 1 读权限  2 创建权限  3 修改权限  4 删除权限
	Level AuthLevel `json:"level"`
}
