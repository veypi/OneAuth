package access

import (
	"github.com/veypi/OneBD/rest"
	"oa/cfg"
	M "oa/models"
)

func useAccess(r rest.Router) {
	r.Get("/", accessList)
	r.Post("/", accessPost)
}
func accessList(x *rest.X) (any, error) {
	opts := &M.AccessList{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := make([]*M.Access, 0, 10)

	query := cfg.DB()
	if opts.CreatedAt != nil {
		query = query.Where("created_at > ?", opts.CreatedAt)
	}
	if opts.UpdatedAt != nil {
		query = query.Where("updated_at > ?", opts.UpdatedAt)
	}
	query = query.Where("app_id LIKE ?", opts.AppID)
	if opts.UserID != nil {
		query = query.Where("user_id LIKE ?", opts.UserID)
	}
	if opts.RoleID != nil {
		query = query.Where("role_id LIKE ?", opts.RoleID)
	}
	if opts.Name != nil {
		query = query.Where("name LIKE ?", opts.Name)
	}
	err = query.Find(&data).Error

	return data, err
}
func accessPost(x *rest.X) (any, error) {
	opts := &M.AccessPost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Access{}

	data.AppID = opts.AppID
	data.UserID = opts.UserID
	data.RoleID = opts.RoleID
	data.Name = opts.Name
	data.TID = opts.TID
	data.Level = opts.Level
	err = cfg.DB().Create(data).Error

	return data, err
}
