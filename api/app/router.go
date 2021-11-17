package app

import (
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

func Router(r OneBD.Router) {
	r.Set("/", appHandlerP, rfc.MethodPost, rfc.MethodGet)
	r.Set("/:uuid", appHandlerP, rfc.MethodGet, rfc.MethodPatch)
	r.Set("/:uuid/user/:id", auHandlerP, rfc.MethodAll)
}
