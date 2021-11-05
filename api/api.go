package api

import (
	"OneAuth/api/app"
	"OneAuth/api/role"
	"OneAuth/api/token"
	"OneAuth/api/user"
	"OneAuth/api/wx"
	"OneAuth/libs/base"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/core"
)

func Router(r OneBD.Router) {
	r.SetNotFoundFunc(func(m core.Meta) {
		base.JSONResponse(m, nil, nil)
	})
	r.SetInternalErrorFunc(func(m core.Meta) {
		base.JSONResponse(m, nil, nil)
	})
	user.Router(r.SubRouter("/user"))
	wx.Router(r.SubRouter("wx"))
	app.Router(r.SubRouter("app"))
	token.Router(r.SubRouter("token"))
	role.Router(r)

	//message.Router(r.SubRouter("/message"))
}
