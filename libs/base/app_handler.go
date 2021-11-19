package base

import "github.com/veypi/OneBD"

/**
* @name: app_handler
* @author: veypi <i@veypi.com>
* @date: 2021-11-18 15:27
* @description：app_handler
**/

type AppHandler struct {
	ApiHandler
	UUID string
}

func (h *AppHandler) Init(m OneBD.Meta) error {
	h.UUID = m.Params("uuid")
	return h.ApiHandler.Init(m)
}
