//
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-09-20 16:10:16
// Distributed under terms of the MIT license.
//

package models

import (
	"gorm.io/gorm"
	"oa/cfg"
	"time"
)

type BaseModel struct {
	// ID        uint           `json:"id" gorm:"primaryKey" methods:"get,patch,delete" parse:"path"`
	ID string `json:"id" gorm:"primaryKey;type:varchar(32)" methods:"get,put,patch,delete" parse:"path"`
	BaseDate
}

type BaseDate struct {
	CreatedAt time.Time      `json:"created_at" methods:"*list" parse:"query"`
	UpdatedAt time.Time      `json:"updated_at" methods:"*list" parse:"query"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func init() {
	cfg.ObjList = append(cfg.ObjList, &AppUser{})
	cfg.ObjList = append(cfg.ObjList, &Resource{})
	cfg.ObjList = append(cfg.ObjList, &Access{})
	cfg.ObjList = append(cfg.ObjList, &Role{})
	cfg.ObjList = append(cfg.ObjList, &User{})
	cfg.ObjList = append(cfg.ObjList, &UserRole{})
	cfg.ObjList = append(cfg.ObjList, &App{})
}
