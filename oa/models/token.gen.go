package models

import "time"

type TokenSalt struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path"`
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

type TokenPost struct {
	UserID    string     `json:"user_id" gorm:"index;type:varchar(32)"  parse:"json"`
	AppID     string     `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
	ExpiredAt *time.Time `json:"expired_at"  parse:"json"`
	OverPerm  *string    `json:"over_perm"  parse:"json"`
}

type TokenList struct {
	UserID string `json:"user_id" gorm:"index;type:varchar(32)"  parse:"json"`
	AppID  string `json:"app_id" gorm:"index;type:varchar(32)"  parse:"json"`
}
