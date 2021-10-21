package models

type UserRole struct {
	BaseModel
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type Role struct {
	BaseModel
	AppID uint   `json:"app_id"`
	App   *App   `json:"app"`
	Name  string `json:"name"`
	// 角色标签
	Tag   string  `json:"tag" gorm:"default:''"`
	Users []*User `json:"users" gorm:"many2many:user_roles;"`
	// 具体权限
	Auths    []*Auth `json:"auths" gorm:"foreignkey:RoleID;references:ID"`
	IsUnique bool    `json:"is_unique" gorm:"default:false"`
}

// AuthLevel 权限等级
// 对于操作类权限
// 0 禁止执行
// 1 允许执行
// 对于资源类权限
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
	AuthDo   AuthLevel = 1
	// AuthPart TODO: 临时权限
	AuthPart   AuthLevel = 1
	AuthRead   AuthLevel = 2
	AuthCreate AuthLevel = 3
	AuthUpdate AuthLevel = 4
	AuthDelete AuthLevel = 5
	AuthAll    AuthLevel = 6
)

func (a AuthLevel) Upper(b AuthLevel) bool {
	return a > b
}

func (a AuthLevel) CanDo() bool {
	return a > AuthNone
}

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

// Auth 资源权限
type Auth struct {
	BaseModel
	// 该权限作用的应用
	AppID uint `json:"app_id"`
	App   *App `json:"app"`
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
	RUID  string    `json:"ruid"`
	Level AuthLevel `json:"level"`
}

type Resource struct {
	BaseModel
	AppID uint   `json:"app_id"`
	App   *App   `json:"app"`
	Name  string `json:"name"`
	// 权限标签
	Tag string `json:"tag"`
	Des string `json:"des"`
}
