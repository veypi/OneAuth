package models

type UserRole struct {
	BaseModel
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type Role struct {
	BaseModel
	Name string `json:"name"`
	// 角色类型
	// 1: 系统定义角色   2: 用户自定义角色
	Category uint `json:"category" gorm:"default:1"`
	// 角色标签
	Tag   string  `json:"tag" gorm:"default:''"`
	Users []*User `json:"users" gorm:"many2many:user_role;"`
	// 具体权限
	Auths    []*Auth `json:"auths" gorm:"foreignkey:RoleID;references:ID"`
	IsUnique bool    `json:"is_unique" gorm:"default:false"`
}

// AuthLevel 权限等级
// 0 相当于没有
// 1 有限读权限
// 2 读权限
// 3 创建权限
// 4 修改权限
// 5 删除权限
// 6 赋予其余人权限
type AuthLevel uint

const (
	AuthNone AuthLevel = 0
	// AuthPart TODO: 临时权限
	AuthPart   AuthLevel = 1
	AuthRead   AuthLevel = 2
	AuthCreate AuthLevel = 3
	AuthUpdate AuthLevel = 4
	AuthDelete AuthLevel = 5
	AuthAll    AuthLevel = 6
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
	return a >= AuthAll
}

// 资源权限

type Auth struct {
	BaseModel
	Name string `json:"name"`
	// 该权限作用的应用
	AppID uint `json:"app_id"`
	// 权限绑定只能绑定一个
	RoleID uint `json:"role_id"`
	UserID uint `json:"user_id"`
	// 资源id
	RID string `json:"rid" gorm:""`
	// 具体某个资源的id
	RUID string `json:"ruid"`
	// 权限标签
	Tag   string    `json:"tag"`
	Level AuthLevel `json:"level"`
	Des   string    `json:"des"`
}
