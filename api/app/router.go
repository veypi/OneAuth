package app

import (
	"errors"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/oalib"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils/jwt"
)

func Router(r OneBD.Router) {
	r.Set("/", appHandlerP, rfc.MethodPost, rfc.MethodGet)
	r.Set("/:uuid", appHandlerP, rfc.MethodGet, rfc.MethodPatch)
	r.Set("/:uuid/user/:id", auHandlerP, rfc.MethodAll)
	r.Set("/:uuid/ping", ping, rfc.MethodGet)
}

func ping(m OneBD.Meta) {
	var err error
	defer func() {
		if err != nil {
			m.WriteHeader(rfc.StatusBadRequest)
			m.Write([]byte(err.Error()))
		} else {
			m.WriteHeader(rfc.StatusOK)
			m.Write([]byte("ok"))
		}
	}()
	t := m.GetHeader("auth_token")
	uuid := m.Params("uuid")
	a := &models.App{}
	err = cfg.DB().Where("UUID = ?", uuid).First(a).Error
	if err != nil {
		return
	}
	p := &oalib.PayLoad{}
	ok, err := jwt.ParseToken(t, p, []byte(a.Key))
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("invalid key")
	}
}
