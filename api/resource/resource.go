package resource

import (
	"errors"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/libs/tools"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneBD"
)

/**
* @name: resource
* @author: veypi <i@veypi.com>
* @date: 2021-11-18 15:25
* @description：resource
**/

var resP = OneBD.NewHandlerPool(func() OneBD.Handler {
	return &resourceHandler{}
})

type resourceHandler struct {
	base.AppHandler
}

func (h *resourceHandler) Get() (interface{}, error) {
	if !h.GetAuth(auth.Res, h.UUID).CanRead() {
		return nil, oerr.NoAuth
	}
	list := make([]*models.Resource, 0, 10)
	err := cfg.DB().Where("AppUUID = ?", h.UUID).Find(&list).Error
	return list, err
}

func (h *resourceHandler) Post() (interface{}, error) {
	if !h.GetAuth(auth.Res, h.UUID).CanCreate() {
		return nil, oerr.NoAuth
	}
	props := &struct {
		Name string
		Des  string
	}{}
	err := h.Meta().ReadJson(props)
	if err != nil {
		return nil, err
	}
	res := &models.Resource{
		AppUUID: h.UUID,
		Name:    props.Name,
		Des:     props.Des,
	}
	err = cfg.DB().Create(res).Error
	return res, err
}

func (h *resourceHandler) Patch() (interface{}, error) {
	if !h.GetAuth(auth.Res, h.UUID).CanUpdate() {
		return nil, oerr.NoAuth
	}
	props := struct {
		Des *string
	}{}
	err := h.Meta().ReadJson(&props)
	if err != nil {
		return nil, err
	}

	query := tools.Struct2Map(props)
	id := h.Meta().ParamsInt("id")
	if len(query) == 0 || id <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	if err := cfg.DB().Table("Resources").Where("id = ?", id).Updates(query).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *resourceHandler) Delete() (interface{}, error) {
	if !h.GetAuth(auth.Res, h.UUID).CanDelete() {
		return nil, oerr.NoAuth
	}
	id := uint(h.Meta().ParamsInt("id"))
	if id <= 0 {
		return nil, oerr.ApiArgsError
	}
	list := make([]*models.Auth, 0, 10)
	err := cfg.DB().Where("ResourceID = ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return nil, errors.New("关联权限未删除")
	}
	err = cfg.DB().Delete(&models.Resource{}, id).Error
	return nil, err
}
