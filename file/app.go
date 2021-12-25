package file

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/oalib"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"net/http"
	"strconv"
)

/**
* @name: user
* @author: veypi <i@veypi.com>
* @date: 2021-12-04 11:49
* @description：user
**/
func appFileChecker(w http.ResponseWriter, r *http.Request) (prefix string, mountPoint string, ownerID string, actorID string, err error) {
	m := w.(OneBD.Meta)
	uuid := m.Params("uuid")
	p := &oalib.PayLoad{}
	h := r.Header.Get("auth_token")
	if h == "" {
		h = m.Query("auth_token")
	}
	var ok bool
	ok, err = p.ParseToken(h, cfg.CFG.APPKey)
	if !ok {
		err = oerr.NoAuth
		return
	}
	l := p.GetAuth(auth.APP, uuid)
	if !l.CanRead() {
		err = oerr.NoAuth
	}
	if !l.CanDelete() && r.Method == rfc.MethodDelete {
		err = oerr.NoAuth
	}
	if !l.CanUpdate() && (r.Method == "PUT" || r.Method == "MKCOL" || r.Method == "COPY" || r.Method == "MOVE") {
		err = oerr.NoAuth
	}
	if err != nil {
		return
	}
	actorID = strconv.Itoa(int(p.ID))
	ownerID = uuid
	mountPoint = uuid
	prefix = uuid
	return
}
