package auth

import (
	"OneAuth/cfg"
	"OneAuth/libs/oerr"
	"OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

type UserHandler struct {
	Payload      *models.PayLoad
	ignoreMethod map[rfc.Method]bool
}

func (a *UserHandler) Init(m OneBD.Meta) error {
	if a.ignoreMethod != nil && a.ignoreMethod[m.Method()] {
		return nil
	}
	a.Payload = new(models.PayLoad)
	token := m.GetHeader("auth_token")
	if token == "" {
		return oerr.NotLogin
	}
	ok, err := models.ParseToken(token, cfg.CFG.Key, a.Payload)
	if ok {
		return nil
	}
	return oerr.NotLogin.Attach(err)
}

func (a *UserHandler) Ignore(methods ...rfc.Method) {
	if a.ignoreMethod == nil {
		a.ignoreMethod = make(map[rfc.Method]bool)
	}
	for _, m := range methods {
		a.ignoreMethod[m] = true
	}
}

func (a *UserHandler) GetAuth(ResourceID string, ResourceUUID ...string) models.AuthLevel {
	return a.Payload.GetAuth(ResourceID, ResourceUUID...)
}
