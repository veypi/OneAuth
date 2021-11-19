package app

import (
	"errors"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/models"
	"gorm.io/gorm"
)

func AddUser(tx *gorm.DB, uuid string, userID uint, roleID uint, status models.AUStatus) (*models.AppUser, error) {
	if uuid == "" || userID == 0 {
		return nil, oerr.FuncArgsError
	}
	au := &models.AppUser{
		AppUUID: uuid,
	}
	au.UserID = userID
	err := tx.Where(au).First(au).Error
	if err == nil {
		return nil, oerr.ResourceDuplicated
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		au.Status = status
		err = tx.Create(au).Error
		if err != nil {
			return nil, err
		}
		if roleID > 0 {
			err = auth.BindUserRole(tx, userID, roleID)
			if err != nil {
				return nil, err
			}
		}
		err = tx.Model(&models.App{}).Where("UUID = ?", uuid).Update("UserCount", gorm.Expr("UserCount + ?", 1)).Error
		return au, err
	}
	return nil, err
}
func EnableUser(tx *gorm.DB, uuid string, userID uint) error {
	if uuid == "" || userID == 0 {
		return oerr.FuncArgsError
	}
	au := &models.AppUser{
		AppUUID: uuid,
	}
	au.UserID = userID
	err := tx.Where(au).First(au).Error
	if err != nil {
		return err
	}
	if au.Status != models.AUOK {
		return tx.Where(au).Update("Status", models.AUOK).Error
	}
	return nil
}

func DisableUser(tx *gorm.DB, uuid string, userID uint) error {
	if uuid == "" || userID == 0 {
		return oerr.FuncArgsError
	}
	au := &models.AppUser{
		AppUUID: uuid,
	}
	au.UserID = userID
	return tx.Where(au).Update("Status", models.AUDisable).Error
}
