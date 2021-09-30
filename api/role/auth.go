package role

import (
	"OneAuth/cfg"
	"OneAuth/libs/base"
	"OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/core"
)

var authP = OneBD.NewHandlerPool(func() core.Handler {
	return &authHandler{}
})

type authHandler struct {
	base.ApiHandler
}

func (h *authHandler) Get() (interface{}, error) {
	l := make([]*models.Auth, 0, 10)
	return &l, cfg.DB().Find(&l).Error
}
