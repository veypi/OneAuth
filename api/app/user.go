package app

import (
	"OneAuth/cfg"
	"OneAuth/libs/app"
	"OneAuth/libs/auth"
	"OneAuth/libs/base"
	"OneAuth/libs/oerr"
	"OneAuth/models"
	"github.com/veypi/OneBD"
)

var auHandlerP = OneBD.NewHandlerPool(func() OneBD.Handler {
	h := &appUserHandler{}
	return h
})

type appUserHandler struct {
	base.ApiHandler
	uuid string
}

func (h *appUserHandler) Init(m OneBD.Meta) error {
	h.uuid = m.Params("uuid")
	if h.uuid == "-" {
		h.uuid = ""
	}
	return h.ApiHandler.Init(m)
}

func (h *appUserHandler) Get() (interface{}, error) {
	id := h.Meta().ParamsInt("id")
	if h.uuid == "" && id == 0 {
		return nil, oerr.ApiArgsMissing
	}
	if uint(id) != h.Payload.ID && !h.Payload.GetAuth(auth.User, h.uuid).CanRead() {
		return nil, oerr.NoAuth
	}
	au := &models.AppUser{
		UserID:  uint(id),
		AppUUID: h.uuid,
	}
	list := make([]*models.AppUser, 0, 10)
	err := cfg.DB().Preload("User").Where(au).Find(&list).Error
	return list, err
}

func (h *appUserHandler) Post() (interface{}, error) {
	id := h.Meta().ParamsInt("id")
	if h.uuid == "" || id <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	status := models.AUOK
	target := &models.App{}
	err := cfg.DB().Where("uuid = ?", h.uuid).First(target).Error
	if err != nil {
		return nil, err
	}
	if target.EnableRegister {
		status = models.AUApply
	}
	au, err := app.AddUser(cfg.DB(), h.uuid, uint(id), target.InitRoleID, status)
	return au, err
}

func (h *appUserHandler) Update() (interface{}, error) {
	id := h.Meta().ParamsInt("id")
	if h.uuid == "" || id <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	props := struct {
		Status string `json:"status"`
	}{}
	err := h.Meta().ReadJson(&props)
	if err != nil {
		return nil, err
	}
	if uint(id) != h.Payload.ID && !h.Payload.GetAuth(auth.User, h.uuid).CanUpdate() {
		return nil, oerr.NoAuth
	}
	au := &models.AppUser{
		UserID:  uint(id),
		AppUUID: h.uuid,
	}
	err = cfg.DB().Where(au).Update("status", props.Status).Error
	return nil, err
}

func (h *appUserHandler) Delete() (interface{}, error) {
	id := h.Meta().ParamsInt("id")
	if h.uuid == "" && id <= 0 {
		return nil, oerr.ApiArgsMissing
	}
	if uint(id) != h.Payload.ID && h.Payload.GetAuth(auth.User, h.uuid).CanDelete() {
		return nil, oerr.NoAuth
	}
	au := &models.AppUser{
		AppUUID: h.uuid,
		UserID:  uint(id),
	}
	list := make([]*models.AppUser, 0, 10)
	err := cfg.DB().Where(au).Delete(&list).Error
	if err != nil {
		return nil, err
	}
	err = cfg.DB().Delete(&list).Error
	return nil, err
}
