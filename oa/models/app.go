package models

import (
	"github.com/veypi/utils/logv"
	"gorm.io/gorm"
)

type App struct {
	BaseModel
	Name        string  `json:"name" methods:"get,post,*patch,*list" parse:"json"`
	Icon        string  `json:"icon" methods:"post,*patch" parse:"json"`
	Des         string  `json:"des" methods:"post,*patch" parse:"json"`
	Participate string  `json:"participate" gorm:"default:auto" methods:"post,*patch" parse:"json"`
	InitRoleID  *string `json:"init_role_id" gorm:"index;type:varchar(32);default: null" methods:"*patch" parse:"json"`
	InitRole    *Role   `json:"init_role" gorm:"foreignKey:InitRoleID;references:ID"`
	InitUrl     string  `json:"init_url"`
	UserCount   uint    `json:"user_count"`
	Key         string  `json:"-"`
}

type AppUser struct {
	BaseModel
	AppID string `json:"app_id" methods:"get,list,post,patch,delete" parse:"path"`
	App   *App   `json:"-" gorm:"foreignKey:AppID;references:ID"`

	UserID string `json:"user_id" methods:"get,*list,post" parse:"json"`
	User   *User  `json:"-" gorm:"foreignKey:UserID;references:ID"`

	Status string `json:"status" methods:"post,*patch,*list" parse:"json"`
}

func (m *AppUser) onOk(tx *gorm.DB) (err error) {
	app := &App{}
	logv.AssertError(tx.Where("id = ?", m.AppID).First(app).Error)
	if app.InitRoleID != nil {
		urList := make([]*UserRole, 0, 2)
		logv.AssertError(tx.Where("app_id = ? && user_id = ?", m.AppID, m.UserID).Find(&urList).Error)
		if len(urList) == 0 {
			return tx.Create(&UserRole{
				AppID:  m.AppID,
				UserID: m.UserID,
				RoleID: *app.InitRoleID,
				Status: "ok",
			}).Error
		}
	}
	return nil
}

func (m *AppUser) AfterCreate(tx *gorm.DB) error {
	if m.Status == "ok" {
		logv.AssertError(m.onOk(tx))
	}
	return tx.Model(&App{}).Where("id = ?", m.AppID).Update("user_count", gorm.Expr("user_count + ?", 1)).Error
}

func (m *AppUser) AfterUpdate(tx *gorm.DB) error {
	if m.Status == "ok" {
		return m.onOk(tx)
	}
	return nil
}

type Resource struct {
	BaseDate
	AppID string `json:"app_id" methods:"get,list,post,patch,delete" parse:"path"`
	App   *App   `json:"-" gorm:"foreignKey:AppID;references:ID"`
	Name  string `json:"name" gorm:"primaryKey" methods:"post,delete" parse:"json"`
	Des   string `json:"des" methods:"post" parse:"json"`
}
