package app

import (
	"OneAuth/libs/auth"
	"OneAuth/libs/oerr"
	"OneAuth/models"
	"errors"
	"gorm.io/gorm"
)

func AddUser(tx *gorm.DB, appID uint, userID uint, roleID uint, status models.AUStatus) error {
	if appID == 0 || userID == 0 {
		return oerr.FuncArgsError
	}
	au := &models.AppUser{}
	au.AppID = appID
	au.UserID = userID
	err := tx.Where(au).First(au).Error
	if err == nil {
		return oerr.ResourceDuplicated
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		au.Status = status
		err = tx.Create(au).Error
		if err != nil {
			return err
		}
		if roleID > 0 {
			err = auth.BindUserRole(tx, userID, roleID)
			if err != nil {
				return err
			}
		}
		return tx.Model(&models.App{}).Where("id = ?", appID).Update("user_count", gorm.Expr("user_count + ?", 1)).Error
	}
	return err
}
func EnableUser(tx *gorm.DB, appID uint, userID uint) error {
	if appID == 0 || userID == 0 {
		return oerr.FuncArgsError
	}
	au := &models.AppUser{}
	au.AppID = appID
	au.UserID = userID
	err := tx.Where(au).First(au).Error
	if err != nil {
		return err
	}
	if au.Status != models.AUOK {
		return tx.Where(au).Update("status", models.AUOK).Error
	}
	return nil
}

func DisableUser(tx *gorm.DB, appID uint, userID uint) error {
	if appID == 0 || userID == 0 {
		return oerr.FuncArgsError
	}
	au := &models.AppUser{}
	au.AppID = appID
	au.UserID = userID
	return tx.Where(au).Update("status", models.AUDisable).Error
}
