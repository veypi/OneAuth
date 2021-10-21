package base

import (
	"OneAuth/libs/oerr"
	"OneAuth/libs/token"
	"OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

type UserHandler struct {
	Payload      *token.PayLoad
	ignoreMethod map[rfc.Method]bool
}

func (a *UserHandler) Init(m OneBD.Meta) error {
	if a.ignoreMethod != nil && a.ignoreMethod[m.Method()] {
		return nil
	}
	return a.ParsePayload(m)
}

func (a *UserHandler) ParsePayload(m OneBD.Meta) error {
	a.Payload = new(token.PayLoad)
	tokenStr := m.GetHeader("auth_token")
	if tokenStr == "" {
		return oerr.NotLogin
	}
	ok, err := token.ParseToken(tokenStr, a.Payload)
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
