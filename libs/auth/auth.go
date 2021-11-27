package auth

import (
	"errors"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/oalib"
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
	ur := &models.UserRole{}
	ur.RoleID = roleID
	ur.UserID = userID
	err := tx.Where(ur).First(ur).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = tx.Create(ur).Error
		if err == nil {
			tx.Model(&models.Role{}).Where("ID = ?", roleID).
				Update("UserCount", gorm.Expr("UserCount + ?", 1))
		}
	}
	return err
}
func UnBindUserRole(tx *gorm.DB, userID uint, roleID uint) error {
	ur := &models.UserRole{}
	ur.RoleID = roleID
	ur.UserID = userID
	err := tx.Unscoped().Where(ur).Delete(ur).Error
	if err == nil {
		tx.Model(&models.Role{}).Where("ID = ?", roleID).
			Update("UserCount", gorm.Expr("UserCount - ?", 1))
	}
	return err
}

func BindUserAuth(tx *gorm.DB, userID uint, resID uint, level oalib.AuthLevel, ruid string) error {
	return bind(tx, userID, resID, level, ruid, false)
}

func BindRoleAuth(tx *gorm.DB, roleID uint, resID uint, level oalib.AuthLevel, ruid string) error {
	return bind(tx, roleID, resID, level, ruid, true)
}

func bind(tx *gorm.DB, id uint, resID uint, level oalib.AuthLevel, ruid string, isRole bool) error {
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
