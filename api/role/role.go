package role

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneBD"
	"gorm.io/gorm"
)

var roleP = OneBD.NewHandlerPool(func() OneBD.Handler {
	return &roleHandler{}
})

type roleHandler struct {
	base.AppHandler
}

func (h *roleHandler) Get() (interface{}, error) {
	id := h.Meta().ParamsInt("id")
	if !h.GetAuth(auth.Role, h.UUID).CanRead() {
		return nil, oerr.NoAuth
	}
	if id > 0 {
		role := &models.Role{}
		role.AppUUID = h.UUID
		role.ID = uint(id)
		err := cfg.DB().Preload("Auths").Preload("Users").First(role).Error
		if err != nil {
			return nil, err
		}
		return role, nil
	}
	roles := make([]*models.Role, 0, 10)
	err := cfg.DB().Where("AppUUID = ?", h.UUID).Find(&roles).Error
	return roles, err
}

func (h *roleHandler) Post() (interface{}, error) {
	if !h.GetAuth(auth.Role).CanCreate() {
		return nil, oerr.NoAuth
	}
	role := &models.Role{
		AppUUID: h.UUID,
	}
	err := h.Meta().ReadJson(role)
	if err != nil {
		return nil, err
	}
	role.ID = 0
	if role.Name == "" {
		return nil, oerr.ApiArgsMissing
	}
	return role, cfg.DB().Where(role).FirstOrCreate(role).Error
}

func (h *roleHandler) Patch() (interface{}, error) {
	if !h.GetAuth(auth.Role).CanUpdate() {
		return nil, oerr.NoAuth
	}
	query := &struct {
		Name *string
		// 角色标签
		Tag *string `gorm:"default:''"`
	}{}
	err := h.Meta().ReadJson(query)
	if err != nil {
		return nil, err
	}
	rid := h.Meta().ParamsInt("id")
	if rid <= 0 {
		return nil, oerr.ApiArgsError
	}
	role := &models.Role{}
	role.ID = uint(rid)
	err = cfg.DB().Preload("Users").Where(role).First(role).Error
	if err != nil {
		return nil, err
	}
	return nil, cfg.DB().Transaction(func(tx *gorm.DB) error {
		var err error
		if query.Tag != nil && *query.Tag != role.Tag {
			err = tx.Model(role).Update("Tag", *query.Tag).Error
			if err != nil {
				return err
			}
		}
		if query.Name != nil && *query.Name != role.Name {
			err = tx.Model(role).Update("Name", *query.Name).Error
			if err != nil {
				return err
			}
		}
		return err
	})
}

func (h *roleHandler) Delete() (interface{}, error) {
	if !h.GetAuth(auth.Role).CanDelete() {
		return nil, oerr.NoAuth
	}
	rid := h.Meta().ParamsInt("id")
	if rid <= 2 {
		return nil, oerr.NoAuth
	}
	role := &models.Role{}
	role.ID = uint(rid)
	err := cfg.DB().Where(role).First(role).Error
	if err != nil {
		return nil, err
	}
	return nil, cfg.DB().Delete(role).Error
}
