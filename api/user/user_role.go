package user

import (
	"errors"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneBD"
	"gorm.io/gorm"
)

var userRoleP = OneBD.NewHandlerPool(func() OneBD.Handler {
	return &userRoleHandler{}
})

type userRoleHandler struct {
	base.ApiHandler
}

func (h *userRoleHandler) Post() (interface{}, error) {
	if !h.GetAuth("role").CanCreate() {
		return nil, oerr.NoAuth
	}
	uid := h.Meta().ParamsInt("user_id")
	if uid <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	query := &models.Role{}
	err := h.Meta().ReadJson(query)
	if err != nil {
		return nil, err
	}
	if query.ID != 0 {
		err = cfg.DB().First(query, query.ID).Error
	} else if query.Name != "" {
		err = cfg.DB().Where(map[string]interface{}{
			"name": query.Name,
			"tag":  query.Tag,
		}).First(query).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = cfg.DB().Create(query).Error
		}
	} else {
		return nil, oerr.ApiArgsMissing
	}
	if err != nil {
		return nil, err
	}
	link := &models.UserRole{}
	link.UserID = uint(uid)
	link.RoleID = query.ID
	err = cfg.DB().Where(link).FirstOrCreate(link).Error
	return link, err
}

func (h *userRoleHandler) Delete() (interface{}, error) {
	if !h.GetAuth("role").CanDelete() {
		return nil, oerr.NoAuth
	}
	uid := h.Meta().ParamsInt("user_id")
	id := h.Meta().ParamsInt("role_id")
	if uid <= 0 || id <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	link := &models.UserRole{}
	link.UserID = uint(uid)
	link.RoleID = uint(id)
	err := cfg.DB().Where(link).First(link).Error
	if err != nil {
		return nil, err
	}
	return nil, cfg.DB().Delete(link).Error
}
