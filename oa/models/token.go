package models

import "time"

// refresh token，由oa 秘钥签发，有效期长, 存储在token表
// app token, 由app 秘钥签发，有效期短, 不存储
// OverPerm 非oa应用获取oa数据的权限，由用户设定
type Token struct {
	BaseModel
	UserID    string    `json:"user_id" gorm:"index;type:varchar(32)" methods:"post,list" parse:"json"`
	User      *User     `json:"-"`
	AppID     string    `json:"app_id" gorm:"index;type:varchar(32)" methods:"post,list" parse:"json"`
	App       *App      `json:"-"`
	ExpiredAt time.Time `json:"expired_at" methods:"*post,*patch" parse:"json"`
	OverPerm  string    `json:"over_perm" methods:"*post,*patch" parse:"json"`
	Device    string    `json:"device" methods:"*post" parse:"json"`
}
