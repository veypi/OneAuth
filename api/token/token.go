package token

import (
	"OneAuth/cfg"
	"OneAuth/libs/app"
	"OneAuth/libs/base"
	"OneAuth/libs/oerr"
	"OneAuth/libs/token"
	"OneAuth/models"
	"errors"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"gorm.io/gorm"
)

func Router(r OneBD.Router) {
	p := OneBD.NewHandlerPool(func() OneBD.Handler {
		return &tokenHandler{}
	})
	r.Set("/:uuid", p, rfc.MethodGet)
}

type tokenHandler struct {
	base.ApiHandler
}

func (h *tokenHandler) Get() (interface{}, error) {
	uuid := h.Meta().Params("uuid")
	if uuid == "" {
		return nil, oerr.ApiArgsMissing.AttachStr("uuid")
	}
	a := &models.App{}
	a.UUID = uuid
	err := cfg.DB().Where("uuid = ?", uuid).First(a).Error
	if err != nil {
		return nil, err
	}
	au := &models.AppUser{
		UserID: h.Payload.ID,
		AppID:  a.ID,
	}
	err = cfg.DB().Where(au).First(au).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if a.EnableRegister {
				err = cfg.DB().Transaction(func(tx *gorm.DB) error {
					return app.AddUser(cfg.DB(), au.AppID, au.UserID, a.InitRoleID, models.AUOK)
				})
				if err != nil {
					return nil, err
				}
				au.Status = models.AUOK
			} else {
				return nil, oerr.AppNotJoin.AttachStr(a.Name)
			}
		}
		return nil, oerr.DBErr.Attach(err)
	}
	if au.Status != models.AUOK {
		return nil, oerr.NoAuth.AttachStr(string(au.Status))
	}
	u := &models.User{}
	err = cfg.DB().Preload("Auths").Preload("Roles.Auths").Where("id = ?", h.Payload.ID).First(u).Error
	if err != nil {
		return nil, err
	}
	t, err := token.GetToken(u, a.ID, a.Key)
	return t, err
}
