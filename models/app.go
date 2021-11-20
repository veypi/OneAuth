package models

import "gorm.io/gorm"

var AppKeys = map[string]string{}

type App struct {
	UUID      string `gorm:"primaryKey;size:32"`
	CreatedAt JSONTime
	UpdatedAt JSONTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Icon      string
	Des       string
	Creator   uint
	UserCount uint
	Users     []*User `gorm:"many2many:AppUsers;"`
	// 初始用户角色
	InitRoleID uint
	InitRole   *Role
	// 是否在首页隐藏
	Hide bool
	// PubKey     string
	// PrivateKey string
	// 认证成功跳转链接
	Host string
	// 加解密用户token (key+key2)
	// 两个key都是请求获取时刷新
	// key oa发放给app 双方保存 针对app生成 每个应用有一个
	// key2 app发放给oa app保存 oa使用一次销毁 针对当个用户生成 每个用户有一个
	// 获取app用户加密秘钥key2
	// TODO
	UserRefreshUrl string
	// app 校验用户token时使用
	Key string `json:"-"`
	// 是否允许用户自动加入应用
	EnableRegister bool
	//
	EnableUserKey bool
	UserKeyUrl    string
	// 允许登录方式
	EnableUser  bool
	EnableWx    bool
	EnablePhone bool
	EnableEmail bool

	Wx *Wechat `gorm:"-"`
}

type AUStatus string

const (
	AUOK      AUStatus = "ok"
	AUDisable AUStatus = "disabled"
	AUApply   AUStatus = "apply"
	AUDeny    AUStatus = "deny"
)

type AppUser struct {
	BaseModel
	AppUUID string `gorm:"size:32"`
	App     *App   `gorm:"association_foreignkey:UUID"`
	UserID  uint
	User    *User
	Status  AUStatus
}

type Wechat struct {
	BaseModel
	AppUUID string `gorm:"size:32"`
	App     *App   `gorm:"association_foreignkey:UUID"`
	// 网页授权登录用
	WxID    string
	AgentID string
	Url     string

	// 获取access_token用
	CorpID     string
	CorpSecret string
}
