package role

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/core"
)

/**
* @name: user
* @author: veypi <i@veypi.com>
* @date: 2021-11-23 10:17
* @description：user
**/

var ruP = OneBD.NewHandlerPool(func() core.Handler {
	return &roleUserHandler{}
})

type roleUserHandler struct {
	base.AppHandler
}

func (h *roleUserHandler) Get() (interface{}, error) {
	if !h.GetAuth(auth.Role, h.UUID).CanRead() {
		return nil, oerr.NoAuth
	}
	id := h.Meta().ParamsInt("id")
	if id <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	r := &models.Role{}
	err := cfg.DB().Preload("Users").Where("ID = ?", id).First(r).Error
	if err != nil {
		return nil, err
	}
	if r.AppUUID != h.UUID {
		return nil, oerr.NoAuth
	}
	return r.Users, nil
}

func (h *roleUserHandler) Post() (interface{}, error) {
	if !h.GetAuth(auth.Role, h.UUID).CanCreate() {
		return nil, oerr.NoAuth
	}
	id := h.Meta().ParamsInt("id")
	uid := h.Meta().ParamsInt("uid")
	if id <= 0 || uid <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	r := &models.Role{}
	err := cfg.DB().Where("ID = ?", id).First(r).Error
	if err != nil {
		return nil, err
	}
	if r.AppUUID != h.UUID {
		return nil, oerr.NoAuth
	}
	err = auth.BindUserRole(cfg.DB(), uint(uid), uint(id))
	return nil, err
}
func (h *roleUserHandler) Delete() (interface{}, error) {
	if !h.GetAuth(auth.Role, h.UUID).CanCreate() {
		return nil, oerr.NoAuth
	}
	id := h.Meta().ParamsInt("id")
	uid := h.Meta().ParamsInt("uid")
	if id <= 0 || uid <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	r := &models.Role{}
	err := cfg.DB().Where("ID = ?", id).First(r).Error
	if err != nil {
		return nil, err
	}
	if r.AppUUID != h.UUID {
		return nil, oerr.NoAuth
	}
	err = auth.UnBindUserRole(cfg.DB(), uint(uid), uint(id))
	return nil, err
}
