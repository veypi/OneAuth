package models

type App struct {
	BaseModel
	Name string  `json:"name"`
	UUID string  `json:"uuid"`
	Host string  `json:"host"`
	WxID string  `json:"wx_id" gorm:""`
	Wx   *Wechat `json:"wx" gorm:"association_foreignkey:ID"`
}

type Wechat struct {
	BaseModel
	// 网页授权登录用
	WxID    string `json:"wx_id"`
	AgentID string `json:"agent_id"`
	Url     string `json:"url"`

	// 获取access_token用
	CorpID     string `json:"corp_id"`
	CorpSecret string `json:"corp_secret"`
}
