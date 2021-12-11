package file

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/oalib"
	"net/http"
	"strconv"
)

/**
* @name: user
* @author: veypi <i@veypi.com>
* @date: 2021-12-04 11:49
* @description：user
**/
func userFileChecker(w http.ResponseWriter, r *http.Request) (prefix string, mountPoint string, ownerID string, actorID string, err error) {
	user := &models.User{}
	u, p, ok := r.BasicAuth()
	if ok {
		user.Username = u
		err = cfg.DB().Where("username = ?", u).First(user).Error
		if err != nil {
			return
		}
		var isAuth bool
		isAuth, err = user.CheckLogin(p)
		if err != nil || !isAuth {
			return
		}
	} else {
		p := &oalib.PayLoad{}
		h := r.Header.Get("auth_token")
		ok, err = p.ParseToken(h, cfg.CFG.APPKey)
		if !ok {
			return
		}
		user.ID = p.ID
		err = cfg.DB().Where("ID = ?", p.ID).First(user).Error
		if err != nil {
			return
		}
	}
	if user.ID > 0 {
		actorID = strconv.Itoa(int(user.ID))
		ownerID = actorID
		mountPoint = actorID
		prefix = cfg.CFG.FileUrlPrefix + "/usr/"
		return
	}
	return
}
