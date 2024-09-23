package app

import (
	"github.com/veypi/OneBD/rest"
	M "oa/models"
	"oa/cfg"
	"strings"
	"github.com/google/uuid"
)

func useApp(r rest.Router) {
	r.Delete("/:app_id", appDelete)
	r.Get("/:app_id", appGet)
	r.Get("/", appList)
	r.Patch("/:app_id", appPatch)
	r.Post("/", appPost)
	r.Put("/:app_id", appPut)
}
func appDelete(x *rest.X) (any, error) {
	opts := &M.AppDelete{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.App{}

	err = cfg.DB().Where("id = ?", opts.ID).Delete(data).Error

	return data, err
}
func appGet(x *rest.X) (any, error) {
	opts := &M.AppGet{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.App{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error

	return data, err
}
func appList(x *rest.X) (any, error) {
	opts := &M.AppList{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := make([]*M.App, 0, 10)

	query := cfg.DB()
	if opts.Name != nil {
		query = query.Where("name LIKE ?", opts.Name)
	}
	err = query.Find(&data).Error

	return data, err
}
func appPatch(x *rest.X) (any, error) {
	opts := &M.AppPatch{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.App{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error
	if err != nil {
		return nil, err
	}
	optsMap := make(map[string]interface{})
	if opts.Name != nil {
		optsMap["name"] = opts.Name
	}
	if opts.Icon != nil {
		optsMap["icon"] = opts.Icon
	}
	if opts.Des != nil {
		optsMap["des"] = opts.Des
	}
	if opts.Participate != nil {
		optsMap["participate"] = opts.Participate
	}
	if opts.InitRoleID != nil {
		optsMap["init_role_id"] = opts.InitRoleID
	}
	err = cfg.DB().Model(data).Updates(optsMap).Error

	return data, err
}
func appPost(x *rest.X) (any, error) {
	opts := &M.AppPost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.App{}

	data.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	data.Name = opts.Name
	data.Icon = opts.Icon
	data.Des = opts.Des
	data.Participate = opts.Participate
	err = cfg.DB().Create(data).Error

	return data, err
}
func appPut(x *rest.X) (any, error) {
	opts := &M.AppPut{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.App{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error
	if err != nil {
		return nil, err
	}
	optsMap := map[string]interface{}{
		"id":		opts.ID,
		"name":		opts.Name,
		"icon":		opts.Icon,
		"des":		opts.Des,
		"participate":	opts.Participate,
		"init_role_id":	opts.InitRoleID,
	}
	err = cfg.DB().Model(data).Updates(optsMap).Error

	return data, err
}
