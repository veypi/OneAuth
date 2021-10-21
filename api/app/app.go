package app

import (
	"OneAuth/cfg"
	"OneAuth/libs/auth"
	"OneAuth/libs/base"
	"OneAuth/libs/oerr"
	"OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
)

func Router(r OneBD.Router) {
	r.Set("/:id", appHandlerP, rfc.MethodGet)
}

var appHandlerP = OneBD.NewHandlerPool(func() OneBD.Handler {
	h := &appHandler{}
	h.Ignore(rfc.MethodGet)
	return h
})

type appHandler struct {
	base.ApiHandler
	query *models.App
}

func (h *appHandler) Get() (interface{}, error) {
	id := h.Meta().Params("id")
	h.query = &models.App{}
	isSelf := h.Meta().Query("is_self")
	if isSelf != "" {
		// 无权限可以获取本系统基本信息
		h.query.ID = cfg.CFG.APPID
		err := cfg.DB().Where(h.query).First(h.query).Error
		return h.query, err
	}
	err := h.ParsePayload(h.Meta())
	if err != nil {
		return nil, err
	}
	if !h.GetAuth(auth.APP, id).CanRead() {
		return nil, oerr.NoAuth
	}
	if id != "" {
		h.query.UUID = id
		err := cfg.DB().Where(h.query).First(h.query).Error
		return h.query, err
	}
	// 注释代码为获取已经绑定的应用
	//user := &models.User{}
	//user.ID = h.Payload.ID
	//err := cfg.DB().Preload("Roles.Auths").Preload("Auths").Where(user).First(user).Error
	//if err != nil {
	//	return nil, oerr.DBErr.Attach(err)
	//}
	//ids := make([]string, 0, 10)
	//for _, a := range user.GetAuths() {
	//	if a.RID == auth.Login && a.Level.CanDo() {
	//		ids = append(ids, a.RUID)
	//	}
	//}
	list := make([]*models.App, 0, 10)
	err = cfg.DB().Find(&list).Error
	return list, err
}
