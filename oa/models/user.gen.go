package models

import ()

type UserDelete struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_id"`
}

type UserGet struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_id"`
}

type UserList struct {
	Username *string `json:"username" gorm:"varchar(100);unique;default:not null"  parse:"json"`
	Nickname *string `json:"nickname"  parse:"json"`
	Email    *string `json:"email" gorm:"varchar(20);unique;default:null"  parse:"json"`
	Phone    *string `json:"phone" gorm:"varchar(50);unique;default:null"  parse:"json"`
	Status   *uint   `json:"status"  parse:"json"`
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

type UserPost struct {
	Username string  `json:"username" gorm:"varchar(100);unique;default:not null"  parse:"json"`
	Nickname *string `json:"nickname"  parse:"json"`
	Icon     *string `json:"icon"  parse:"json"`
	Email    *string `json:"email" gorm:"varchar(20);unique;default:null"  parse:"json"`
	Phone    *string `json:"phone" gorm:"varchar(50);unique;default:null"  parse:"json"`
}

type UserRoleDelete struct {
	ID     string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_role_id"`
	UserID string `json:"user_id"  parse:"path"`
	RoleID string `json:"role_id"  parse:"path"`
}

type UserRoleGet struct {
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_role_id"`
}

type UserRoleList struct {
	Status *string `json:"status"  parse:"json"`
}

type UserRolePatch struct {
	ID     string  `json:"id" gorm:"primaryKey;type:varchar(32)"  parse:"path@user_role_id"`
	Status *string `json:"status"  parse:"json"`
}

type UserRolePost struct {
	UserID string `json:"user_id"  parse:"path"`
	RoleID string `json:"role_id"  parse:"path"`
	Status string `json:"status"  parse:"json"`
}
