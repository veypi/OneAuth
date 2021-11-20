package auth

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/oalib"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/core"
	"strconv"
)

var authP = OneBD.NewHandlerPool(func() core.Handler {
	return &authHandler{}
})

type authHandler struct {
	base.AppHandler
}

func (h *authHandler) Get() (interface{}, error) {
	if !h.GetAuth(auth.Auth, h.UUID).CanRead() {
		return nil, oerr.NoAuth
	}
	var err error
	id, _ := strconv.Atoi(h.Meta().Query("rid"))
	rid := uint(id)
	id, _ = strconv.Atoi(h.Meta().Query("uid"))
	uid := uint(id)
	if rid == 0 && uid == 0 {
		return nil, err
	}
	query := &models.Auth{
		AppUUID: h.UUID,
	}
	if rid != 0 {
		query.RoleID = &rid
	} else if uid != 0 {
		query.UserID = &uid
	}
	l := make([]*models.Auth, 0, 10)
	err = cfg.DB().Where(query).Find(&l).Error
	return l, err
}

func (h *authHandler) Post() (interface{}, error) {
	if !h.GetAuth(auth.Auth, h.UUID).CanCreate() {
		return nil, oerr.NoAuth
	}
	query := &models.Auth{}
	err := h.Meta().ReadJson(query)
	if err != nil {
		return nil, err
	}
	if query.ResourceID == 0 {
		return nil, oerr.ApiArgsError
	}
	query.AppUUID = h.UUID
	res := &models.Resource{}
	res.ID = query.ResourceID
	err = cfg.DB().First(res).Error
	if err != nil {
		return nil, err
	}
	query.RID = res.Name
	err = cfg.DB().Create(query).Error
	return query, err
}

func (h *authHandler) Patch() (interface{}, error) {
	if !h.GetAuth(auth.Auth, h.UUID).CanUpdate() {
		return nil, oerr.NoAuth
	}
	id := h.Meta().ParamsInt("id")
	if id <= 0 {
		return nil, oerr.ApiArgsError
	}
	a := &models.Auth{}
	a.ID = uint(id)
	err := cfg.DB().First(a).Error
	if err != nil {
		return nil, err
	}
	opts := struct {
		ResourceID *uint `gorm:"not null"`
		RUID       *string
		Level      *oalib.AuthLevel
	}{}
	err = h.Meta().ReadJson(&opts)
	if err != nil {
		return nil, err
	}
	if a.AppUUID != h.UUID {
		return nil, oerr.ApiArgsError
	}
	query := map[string]interface{}{}
	if opts.ResourceID != nil && a.ResourceID != *opts.ResourceID {
		query["ResourceID"] = *opts.ResourceID
		res := &models.Resource{}
		res.ID = *opts.ResourceID
		err = cfg.DB().First(res).Error
		if err != nil {
			return nil, err
		}
		query["RID"] = res.Name
		a.ResourceID = *opts.ResourceID
	}
	if opts.RUID != nil {
		query["RUID"] = *opts.RUID
		a.RUID = *opts.RUID
	}
	if opts.Level != nil {
		query["Level"] = *opts.Level
		a.Level = *opts.Level
	}
	err = cfg.DB().Model(a).Where("id = ?", id).Updates(query).Error
	return a, err
}

func (h *authHandler) Delete() (interface{}, error) {
	if !h.GetAuth(auth.Auth, h.UUID).CanDelete() {
		return nil, oerr.NoAuth
	}
	id := h.Meta().ParamsInt("id")
	if id <= 0 {
		return nil, oerr.ApiArgsError
	}
	a := &models.Auth{}
	a.ID = uint(id)
	err := cfg.DB().First(a).Error
	if err != nil {
		return nil, err
	}
	if a.AppUUID != h.UUID {
		return nil, oerr.ApiArgsError
	}
	return nil, cfg.DB().Delete(a).Error
}
