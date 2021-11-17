package role

import (
	"OneAuth/cfg"
	"OneAuth/libs/auth"
	"OneAuth/libs/base"
	"OneAuth/libs/oerr"
	"OneAuth/libs/token"
	"OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/core"
	"strconv"
)

var authP = OneBD.NewHandlerPool(func() core.Handler {
	return &authHandler{}
})

type authHandler struct {
	base.ApiHandler
}

func (h *authHandler) Get() (interface{}, error) {
	if !h.GetAuth(auth.Auth).CanRead() {
		return nil, oerr.NoAuth
	}
	aid := h.Meta().ParamsInt("id")
	query := &models.Auth{}
	var err error
	if aid > 0 {
		err = cfg.DB().Where("id = ?", aid).First(query).Error
		return query, err
	}
	id, _ := strconv.Atoi(h.Meta().Query("id"))
	uuid := h.Meta().Query("uuid")
	if id == 0 || uuid == "" {
		return nil, oerr.ApiArgsMissing
	}
	target := &models.App{}
	err = cfg.DB().Where("uuid = ?", uuid).First(target).Error
	if err != nil {
		return nil, err
	}
	u := &models.User{}
	err = cfg.DB().Preload("Roles.Auths").Preload("Auths").Where("id = ?", id).First(u).Error
	if err != nil {
		return nil, err
	}
	l := make([]*token.SimpleAuth, 0, 10)
	for _, as := range u.GetAuths() {
		if as.AppUUID == uuid {
			l = append(l, &token.SimpleAuth{
				RID:   as.RID,
				RUID:  as.RUID,
				Level: as.Level,
			})
		}
	}
	return l, nil
}
