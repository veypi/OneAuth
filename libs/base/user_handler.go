package base

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/oalib"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

type UserHandler struct {
	Payload      *oalib.PayLoad
	ignoreMethod map[rfc.Method]bool
}

func (h *UserHandler) Init(m OneBD.Meta) error {
	if h.ignoreMethod != nil && h.ignoreMethod[m.Method()] {
		return nil
	}
	return h.ParsePayload(m)
}

func (h *UserHandler) ParsePayload(m OneBD.Meta) error {
	h.Payload = new(oalib.PayLoad)
	tokenStr := m.GetHeader("auth_token")
	if tokenStr == "" {
		return oerr.NotLogin
	}
	uuid := m.GetHeader("uuid")
	var ok bool
	var err error
	if uuid != "" {
		a := &models.App{}
		err = cfg.DB().Where("UUID = ?", uuid).First(a).Error
		if err != nil {
			return err
		}
		ok, err = h.Payload.ParseToken(tokenStr, []byte(a.Key))
		h.Payload.Auth = nil
	} else {
		ok, err = h.Payload.ParseToken(tokenStr, cfg.CFG.APPKey)
	}
	if ok {
		return nil
	}
	return oerr.NotLogin.Attach(err)
}

func (h *UserHandler) Ignore(methods ...rfc.Method) {
	if h.ignoreMethod == nil {
		h.ignoreMethod = make(map[rfc.Method]bool)
	}
	for _, m := range methods {
		h.ignoreMethod[m] = true
	}
}

func (h *UserHandler) GetAuth(ResourceID string, ResourceUUID ...string) oalib.AuthLevel {
	return h.Payload.GetAuth(ResourceID, ResourceUUID...)
}
