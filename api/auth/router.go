package auth

import (
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

func Router(r OneBD.Router) {
	r.Set("/", authP, rfc.MethodGet, rfc.MethodPost)
	r.Set("/:id", authP, rfc.MethodPatch, rfc.MethodDelete)
}
