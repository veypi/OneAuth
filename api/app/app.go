package app

import (
	"OneAuth/cfg"
	"OneAuth/libs/base"
	"OneAuth/libs/oerr"
	"OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

func Router(r OneBD.Router) {
	r.Set("/:id", appHandlerP, rfc.MethodGet)
}

var appHandlerP = OneBD.NewHandlerPool(func() OneBD.Handler {
	h := &appHandler{}
	h.Ignore(rfc.MethodGet)
	return h
})

type appHandler struct {
	base.ApiHandler
	query *models.App
}

func (h *appHandler) Get() (interface{}, error) {
	id := h.Meta().Params("id")
	if id == "" {
		return nil, oerr.ApiArgsMissing
	}
	h.query = &models.App{}
	h.query.UUID = id
	err := cfg.DB().Where(h.query).Preload("Wx").First(h.query).Error
	if err != nil {
		return nil, err
	}
	return h.query, nil
}
