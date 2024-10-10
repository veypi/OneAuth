package app

import (
	"github.com/google/uuid"
	"github.com/veypi/OneBD/rest"
	"oa/cfg"
	M "oa/models"
	"strings"
)

func useAppUser(r rest.Router) {
	r.Post("/", appUserPost)
	r.Delete("/:app_user_id", appUserDelete)
	r.Get("/:app_user_id", appUserGet)
	r.Get("/", appUserList)
	r.Patch("/:app_user_id", appUserPatch)
}
func appUserPost(x *rest.X) (any, error) {
	opts := &M.AppUserPost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.AppUser{}

	data.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	data.AppID = opts.AppID
	data.UserID = opts.UserID
	data.Status = opts.Status
	err = cfg.DB().Create(data).Error

	return data, err
}
func appUserDelete(x *rest.X) (any, error) {
	opts := &M.AppUserDelete{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.AppUser{}

	err = cfg.DB().Where("id = ?", opts.ID).Delete(data).Error

	return data, err
}
func appUserGet(x *rest.X) (any, error) {
	opts := &M.AppUserGet{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.AppUser{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error

	return data, err
}
func appUserList(x *rest.X) (any, error) {
	opts := &M.AppUserList{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := make([]*M.AppUser, 0, 10)

	query := cfg.DB()
	if opts.AppID != nil {
		query = query.Where("app_id LIKE ?", opts.AppID)
	}
	if opts.UserID != nil {
		query = query.Where("user_id LIKE ?", opts.UserID)
	}
	if opts.Status != nil {
		query = query.Where("status LIKE ?", opts.Status)
	}
	err = query.Find(&data).Error

	return data, err
}
func appUserPatch(x *rest.X) (any, error) {
	opts := &M.AppUserPatch{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.AppUser{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error
	if err != nil {
		return nil, err
	}
	optsMap := make(map[string]interface{})
	if opts.Status != nil {
		optsMap["status"] = opts.Status
	}
	err = cfg.DB().Model(data).Updates(optsMap).Error

	return data, err
}
