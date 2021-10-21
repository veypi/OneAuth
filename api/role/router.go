package role

import (
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

func Router(r OneBD.Router) {
	r.Set("/role/", roleP, rfc.MethodGet, rfc.MethodPost)
	r.Set("/role/:id", roleP, rfc.MethodGet, rfc.MethodDelete, rfc.MethodPatch)
	r.Set("/role/:id/:action/:rid", roleP, rfc.MethodGet)
	r.Set("/auth/", authP, rfc.MethodGet)
}
