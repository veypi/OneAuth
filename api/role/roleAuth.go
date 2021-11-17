package role

import (
	"OneAuth/libs/auth"
	"OneAuth/libs/base"
	"OneAuth/libs/oerr"
	"github.com/veypi/OneBD"
)

/**
* @name: roleAuth
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 15:20
* @description：roleAuth
* @update: 2021-11-17 15:20
 */

var rap = OneBD.NewHandlerPool(func() OneBD.Handler {
	return &roleAuthHandler{}
})

type roleAuthHandler struct {
	base.ApiHandler
	id   uint
	aid  uint
	uuid string
}

func (h *roleAuthHandler) Init(m OneBD.Meta) error {
	id := uint(m.ParamsInt("id"))
	aid := uint(m.ParamsInt("aid"))
	if id == 0 || aid == 0 {
		return oerr.ApiArgsError
	}
	return h.ApiHandler.Init(m)
}

func (h *roleAuthHandler) Post() (interface{}, error) {
	if !h.Payload.GetAuth(auth.Auth, h.uuid).CanCreate() {
		return nil, oerr.NoAuth
	}
	return nil, nil
}
