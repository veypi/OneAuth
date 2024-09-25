//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-09-20 16:10:16
// Distributed under terms of the MIT license.
//

package models

import (
	"oa/cfg"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/veypi/utils/logv"
	"gorm.io/gorm"
)

type BaseModel struct {
	// ID        uint           `json:"id" gorm:"primaryKey" methods:"get,patch,delete" parse:"path"`
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)" methods:"get,patch,delete" parse:"path"`
	BaseDate
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return nil
}

type BaseDate struct {
	CreatedAt time.Time      `json:"created_at" methods:"*list" parse:"query"`
	UpdatedAt time.Time      `json:"updated_at" methods:"*list" parse:"query"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func init() {
	cfg.CmdDB.SubCommand("init", "init db data").Command = InitDBData
	cfg.ObjList = append(cfg.ObjList, &AppUser{})
	cfg.ObjList = append(cfg.ObjList, &Resource{})
	cfg.ObjList = append(cfg.ObjList, &Access{})
	cfg.ObjList = append(cfg.ObjList, &Role{})
	cfg.ObjList = append(cfg.ObjList, &User{})
	cfg.ObjList = append(cfg.ObjList, &UserRole{})
	cfg.ObjList = append(cfg.ObjList, &Token{})
	cfg.ObjList = append(cfg.ObjList, &App{})
}

func InitDBData() error {
	app := &App{}
	app.ID = cfg.Config.ID
	logv.AssertError(cfg.DB().Where("id = ?", app.ID).Attrs(app).FirstOrCreate(app).Error)
	initRole := map[string]map[string]uint{
		"user": {"admin": 5, "normal": 1},
		"app":  {"admin": 5, "normal": 2},
	}
	adminID := ""
	for r, roles := range initRole {
		logv.AssertError(cfg.DB().Where("app_id = ? AND name = ?", app.ID, r).FirstOrCreate(&Resource{
			AppID: app.ID,
			Name:  r,
		}).Error)
		for rName, l := range roles {
			role := &Role{}
			logv.AssertError(cfg.DB().Where("app_id = ? AND name = ?", app.ID, rName).Attrs(&Role{
				BaseModel: BaseModel{
					ID: strings.ReplaceAll(uuid.New().String(), "-", ""),
				},
				AppID: app.ID,
				Name:  rName,
			}).FirstOrCreate(role).Error)
			logv.AssertError(cfg.DB().Where("app_id = ? AND role_id = ? AND name = ?", app.ID, role.ID, r).FirstOrCreate(&Access{
				AppID:  app.ID,
				RoleID: &role.ID,
				Name:   r,
				Level:  l,
			}).Error)
			if rName == "admin" {
				adminID = role.ID
			}
		}
	}
	if app.InitRoleID == nil {
		logv.AssertError(cfg.DB().Model(app).Update("init_role_id", adminID).Error)
	}
	return nil
}
