package models

var AppKeys = map[string]string{}

type App struct {
	BaseModel
	Name string `json:"name"`
	Icon string `json:"icon"`
	UUID string `json:"uuid"`
	// 认证成功跳转链接
	Host string `json:"host"`
	// 加解密用户token (key+key2)
	// 两个key都是请求获取时刷新
	// key oa发放给app 双方保存 针对app生成 每个应用有一个
	// key2 app发放给oa app保存 oa使用一次销毁 针对当个用户生成 每个用户有一个
	// 获取app用户加密秘钥key2
	UserRefreshUrl string `json:"user_refresh_url"`
	// app 校验用户token时使用
	Key string `json:"key"`
	// 是否允许用户注册
	EnableRegister string  `json:"enable_register"`
	EnableUser     bool    `json:"enable_user"`
	EnableWx       bool    `json:"enable_wx"`
	EnablePhone    bool    `json:"enable_phone"`
	EnableEmail    bool    `json:"enable_email"`
	Wx             *Wechat `json:"wx" gorm:"foreignkey:AppID;references:ID"`
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
