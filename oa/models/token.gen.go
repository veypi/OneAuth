package models

import "time"

type TokenSalt struct {
	Username string  `json:"username"  parse:"json"`
	Typ      *string `json:"typ" parse:"json"`
}

type TokenPost struct {
	UserID string `json:"user_id" gorm:"index;type:varchar(32)"  parse:"json"`

	// 两种获取token方式，一种用token换取(应用登录)，一种用密码加密code换(oa登录)
	Token *string `json:"token"  parse:"json"`
	// 登录方随机生成的salt，非用户salt
	Salt *string `json:"salt"  parse:"json"`
	Code *string `json:"code"  parse:"json"`

	AppID     *string    `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
	ExpiredAt *time.Time `json:"expired_at"  parse:"json"`
	OverPerm  *string    `json:"over_perm"  parse:"json"`
	Device    *string    `json:"device"  parse:"json"`
}

type TokenGet struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@token_id"`
}

type TokenPatch struct {
	ID        string     `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@token_id"`
	ExpiredAt *time.Time `json:"expired_at"  parse:"json"`
	OverPerm  *string    `json:"over_perm"  parse:"json"`
}

type TokenDelete struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@token_id"`
}

type TokenList struct {
	UserID string `json:"user_id" gorm:"index;type:varchar(32)"  parse:"json"`
	AppID  string `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
}
