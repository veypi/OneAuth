package resource

/**
* @name: router
* @author: veypi <i@veypi.com>
* @date: 2021-11-18 15:24
* @description：router
**/
import (
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

func Router(r OneBD.Router) {
	r.Set("/", resP, rfc.MethodGet, rfc.MethodPost)
	r.Set("/:id", resP, rfc.MethodDelete)
}
