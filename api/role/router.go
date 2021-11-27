package role

import (
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

func Router(r OneBD.Router) {
	r.Set("/", roleP, rfc.MethodGet, rfc.MethodPost)
	r.Set("/:id", roleP, rfc.MethodGet, rfc.MethodDelete, rfc.MethodPatch)
	r.Set("/:id/user/", ruP, rfc.MethodGet)
	r.Set("/:id/user/:uid", ruP, rfc.MethodPost, rfc.MethodDelete)
}
