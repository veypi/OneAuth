package models

var AppKeys = map[string]string{}

type App struct {
	BaseModel
	Name      string  `json:"name"`
	Icon      string  `json:"icon"`
	UUID      string  `json:"uuid" gorm:"unique"`
	Des       string  `json:"des"`
	Creator   uint    `json:"creator"`
	UserCount uint    `json:"user_count"`
	Users     []*User `json:"users" gorm:"many2many:app_users;"`
	// 初始用户角色
	InitRoleID uint  `json:"init_role_id"`
	InitRole   *Role `json:"init_role"`
	// 是否在首页隐藏
	Hide bool `json:"hide"`
	// PubKey     string `json:"pub_key"`
	// PrivateKey string `json:"private_key"`
	// 认证成功跳转链接
	Host string `json:"host"`
	// 加解密用户token (key+key2)
	// 两个key都是请求获取时刷新
	// key oa发放给app 双方保存 针对app生成 每个应用有一个
	// key2 app发放给oa app保存 oa使用一次销毁 针对当个用户生成 每个用户有一个
	// 获取app用户加密秘钥key2
	// TODO
	UserRefreshUrl string `json:"user_refresh_url"`
	// app 校验用户token时使用
	Key string `json:"-"`
	// 是否允许用户自动加入应用
	EnableRegister bool `json:"enable_register"`
	//
	EnableUserKey bool   `json:"enable_user_key"`
	UserKeyUrl    string `json:"user_key_url"`
	// 允许登录方式
	EnableUser  bool    `json:"enable_user"`
	EnableWx    bool    `json:"enable_wx"`
	EnablePhone bool    `json:"enable_phone"`
	EnableEmail bool    `json:"enable_email"`
	Wx          *Wechat `json:"wx" gorm:"foreignkey:AppID;references:ID"`
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
	AppID  uint     `json:"app_id"`
	APP    *App     `json:"app"`
	UserID uint     `json:"user_id"`
	User   *User    `json:"user"`
	Status AUStatus `json:"status"`
}

type Wechat struct {
	BaseModel
	AppID uint `json:"app_id"`
	// 网页授权登录用
	WxID    string `json:"wx_id"`
	AgentID string `json:"agent_id"`
	Url     string `json:"url"`

	// 获取access_token用
	CorpID     string `json:"corp_id"`
	CorpSecret string `json:"corp_secret"`
}
