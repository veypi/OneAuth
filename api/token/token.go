package token

import (
	"errors"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/app"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/libs/token"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"gorm.io/gorm"
)

func Router(r OneBD.Router) {
	p := OneBD.NewHandlerPool(func() OneBD.Handler {
		return &tokenHandler{}
	})
	r.Set("/", p, rfc.MethodGet)
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
	err := cfg.DB().Where("UUID = ?", uuid).First(a).Error
	if err != nil {
		return nil, err
	}
	au := &models.AppUser{
		UserID:  h.Payload.ID,
		AppUUID: a.UUID,
	}
	err = cfg.DB().Where(au).First(au).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if a.EnableRegister {
				err = cfg.DB().Transaction(func(tx *gorm.DB) error {
					_, err := app.AddUser(cfg.DB(), au.AppUUID, au.UserID, a.InitRoleID, models.AUOK)
					return err
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
	err = cfg.DB().Preload("Auths").Preload("Roles.Auths").Where("ID = ?", h.Payload.ID).First(u).Error
	if err != nil {
		return nil, err
	}
	t, err := token.GetToken(u, a.UUID, a.Key)
	return t, err
}
