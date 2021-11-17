package auth

import (
	"OneAuth/models"
	"github.com/veypi/utils"
	"gorm.io/gorm"
)

// 定义oa系统权限

type Resource = string

const (
	// ruid 皆为app uuid
	User Resource = "user"
	APP  Resource = "app"
	Res  Resource = "resource"
	Role Resource = "role"
	Auth Resource = "auth"
)

func BindUserRole(tx *gorm.DB, userID uint, roleID uint) error {
	r := &models.Role{}
	r.ID = roleID
	err := tx.Where(r).First(r).Error
	if err != nil {
		return err
	}
	ur := &models.UserRole{}
	ur.RoleID = roleID
	ur.UserID = userID
	err = utils.MultiErr(
		tx.Where(ur).FirstOrCreate(ur).Error,
		tx.Model(&models.Role{}).Where("id = ?", roleID).
			Update("user_count", gorm.Expr("user_count + ?", 1)).Error,
	)
	return err
}

func BindUserAuth(tx *gorm.DB, userID uint, resID uint, level models.AuthLevel, ruid string) error {
	return bind(tx, userID, resID, level, ruid, false)
}

func BindRoleAuth(tx *gorm.DB, roleID uint, resID uint, level models.AuthLevel, ruid string) error {
	return bind(tx, roleID, resID, level, ruid, true)
}

func bind(tx *gorm.DB, id uint, resID uint, level models.AuthLevel, ruid string, isRole bool) error {
	r := &models.Resource{}
	r.ID = resID
	err := tx.Where(r).First(r).Error
	if err != nil {
		return err
	}
	au := &models.Auth{
		AppUUID:    r.AppUUID,
		ResourceID: resID,
		RID:        r.Name,
		RUID:       ruid,
		Level:      level,
	}
	if isRole {
		au.RoleID = &id
	} else {
		au.UserID = &id
	}
	return tx.Where(au).FirstOrCreate(au).Error
}
