package api

import (
	"OneAuth/api/app"
	"OneAuth/api/role"
	"OneAuth/api/user"
	"OneAuth/api/wx"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/core"
)

func Router(r OneBD.Router) {
	r.SetNotFoundFunc(func(m core.Meta) {
		m.Write([]byte("{\"status\": 0}"))
	})
	r.SetInternalErrorFunc(func(m core.Meta) {
		m.Write([]byte("{\"status\": 0}"))
	})
	user.Router(r.SubRouter("/user"))
	wx.Router(r.SubRouter("wx"))
	app.Router(r.SubRouter("app"))
	role.Router(r)

	//message.Router(r.SubRouter("/message"))
}
