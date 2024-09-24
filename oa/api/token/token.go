package token

import (
	"oa/cfg"
	M "oa/models"
	"strings"

	"github.com/google/uuid"
	"github.com/veypi/OneBD/rest"
)

func useToken(r rest.Router) {
	r.Get("/salt/:id", tokenSalt)
	r.Post("/", tokenPost)
	r.Get("/:token_id", tokenGet)
	r.Patch("/:token_id", tokenPatch)
	r.Delete("/:token_id", tokenDelete)
	r.Get("/", tokenList)
}
func tokenSalt(x *rest.X) (any, error) {
	opts := &M.TokenSalt{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.User{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error
	return data.Salt, err
}
func tokenGet(x *rest.X) (any, error) {
	opts := &M.TokenGet{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Token{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error

	return data, err
}
func tokenPatch(x *rest.X) (any, error) {
	opts := &M.TokenPatch{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Token{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error
	if err != nil {
		return nil, err
	}
	optsMap := make(map[string]interface{})
	if opts.ExpiredAt != nil {
		optsMap["expired_at"] = opts.ExpiredAt
	}
	if opts.OverPerm != nil {
		optsMap["over_perm"] = opts.OverPerm
	}
	err = cfg.DB().Model(data).Updates(optsMap).Error

	return data, err
}
func tokenDelete(x *rest.X) (any, error) {
	opts := &M.TokenDelete{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Token{}

	err = cfg.DB().Where("id = ?", opts.ID).Delete(data).Error

	return data, err
}
func tokenPost(x *rest.X) (any, error) {
	opts := &M.TokenPost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Token{}

	data.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	data.UserID = opts.UserID
	data.AppID = opts.AppID
	if opts.ExpiredAt != nil {
		data.ExpiredAt = *opts.ExpiredAt
	}
	if opts.OverPerm != nil {
		data.OverPerm = *opts.OverPerm
	}
	err = cfg.DB().Create(data).Error

	return data, err
}
func tokenList(x *rest.X) (any, error) {
	opts := &M.TokenList{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := make([]*M.Token, 0, 10)

	query := cfg.DB()
	query = query.Where("user_id = ?", opts.UserID)
	query = query.Where("app_id = ?", opts.AppID)
	err = query.Find(&data).Error

	return data, err
}
