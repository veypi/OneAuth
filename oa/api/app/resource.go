package app

import (
	"github.com/veypi/OneBD/rest"
	"oa/cfg"
	M "oa/models"
)

func useResource(r rest.Router) {
	r.Post("/", resourcePost)
	r.Delete("/", resourceDelete)
	r.Get("/", resourceList)
}
func resourcePost(x *rest.X) (any, error) {
	opts := &M.ResourcePost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Resource{}

	data.AppID = opts.AppID
	data.Name = opts.Name
	data.Des = opts.Des
	err = cfg.DB().Create(data).Error

	return data, err
}
func resourceDelete(x *rest.X) (any, error) {
	opts := &M.ResourceDelete{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Resource{
		AppID: opts.AppID,
		Name:  opts.Name,
	}

	err = cfg.DB().Delete(data).Error

	return data, err
}
func resourceList(x *rest.X) (any, error) {
	opts := &M.ResourceList{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := make([]*M.Resource, 0, 10)

	query := cfg.DB()
	if opts.CreatedAt != nil {
		query = query.Where("created_at > ?", opts.CreatedAt)
	}
	if opts.UpdatedAt != nil {
		query = query.Where("updated_at > ?", opts.UpdatedAt)
	}
	query = query.Where("app_id = ?", opts.AppID)
	err = query.Find(&data).Error

	return data, err
}
