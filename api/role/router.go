package role

import (
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

func Router(r OneBD.Router) {
	r.Set("/", roleP, rfc.MethodGet, rfc.MethodPost)
	r.Set("/:id", roleP, rfc.MethodGet, rfc.MethodDelete, rfc.MethodPatch)
	r.Set("/:id/auth/:aid", roleP, rfc.MethodGet)
	r.Set("/auth/", authP, rfc.MethodGet)
	r.Set("/auth/:id", authP, rfc.MethodGet)
}
