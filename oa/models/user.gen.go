package models

import ()

type UserLogin struct {
	ID  string `json:"id" parse:"path@user_id"`
	Pwd string `json:"pwd" parse:"json"`
	Typ string `json:"typ" parse:"json"`
}

type UserGet struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_id"`
}

type UserPatch struct {
	ID       string  `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_id"`
	Username *string `json:"username" gorm:"varchar(100);unique;default:not null"  parse:"json"`
	Nickname *string `json:"nickname"  parse:"json"`
	Icon     *string `json:"icon"  parse:"json"`
	Email    *string `json:"email" gorm:"varchar(20);unique;default:null"  parse:"json"`
	Phone    *string `json:"phone" gorm:"varchar(50);unique;default:null"  parse:"json"`
	Status   *uint   `json:"status"  parse:"json"`
}

type UserDelete struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_id"`
}

type UserPost struct {
	Username string  `json:"username" gorm:"varchar(100);unique;default:not null"  parse:"json"`
	Nickname *string `json:"nickname"  parse:"json"`
	Icon     *string `json:"icon"  parse:"json"`
	Email    *string `json:"email" gorm:"varchar(20);unique;default:null"  parse:"json"`
	Phone    *string `json:"phone" gorm:"varchar(50);unique;default:null"  parse:"json"`
	Salt     string  `json:"salt" gorm:"varchar(32)"  parse:"json"`
	Code     string  `json:"code" gorm:"varchar(128)"  parse:"json"`
}

type UserList struct {
	Username *string `json:"username" gorm:"varchar(100);unique;default:not null"  parse:"json"`
	Nickname *string `json:"nickname"  parse:"json"`
	Email    *string `json:"email" gorm:"varchar(20);unique;default:null"  parse:"json"`
	Phone    *string `json:"phone" gorm:"varchar(50);unique;default:null"  parse:"json"`
	Status   *uint   `json:"status"  parse:"json"`
}

type UserRoleGet struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_role_id"`
}

type UserRolePatch struct {
	ID     string  `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_role_id"`
	Status *string `json:"status"  parse:"json"`
}

type UserRoleDelete struct {
	ID     string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_role_id"`
	UserID string `json:"user_id"  parse:"path"`
	RoleID string `json:"role_id"  parse:"path"`
}

type UserRolePost struct {
	UserID string `json:"user_id"  parse:"path"`
	RoleID string `json:"role_id"  parse:"path"`
	Status string `json:"status"  parse:"json"`
}

type UserRoleList struct {
	Status *string `json:"status"  parse:"json"`
}
