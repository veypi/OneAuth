package role

import (
	"github.com/google/uuid"
	"github.com/veypi/OneBD/rest"
	"oa/cfg"
	M "oa/models"
	"strings"
)

func useRole(r rest.Router) {
	r.Get("/:role_id", roleGet)
	r.Get("/", roleList)
	r.Patch("/:role_id", rolePatch)
	r.Post("/", rolePost)
	r.Delete("/:role_id", roleDelete)
}
func roleGet(x *rest.X) (any, error) {
	opts := &M.RoleGet{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Role{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error

	return data, err
}
func roleList(x *rest.X) (any, error) {
	opts := &M.RoleList{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := make([]*M.Role, 0, 10)

	query := cfg.DB()
	if opts.Name != nil {
		query = query.Where("name LIKE ?", opts.Name)
	}
	err = query.Find(&data).Error

	return data, err
}
func rolePatch(x *rest.X) (any, error) {
	opts := &M.RolePatch{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Role{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error
	if err != nil {
		return nil, err
	}
	optsMap := make(map[string]interface{})
	if opts.Name != nil {
		optsMap["name"] = opts.Name
	}
	if opts.Des != nil {
		optsMap["des"] = opts.Des
	}
	if opts.AppID != nil {
		optsMap["app_id"] = opts.AppID
	}
	err = cfg.DB().Model(data).Updates(optsMap).Error

	return data, err
}
func rolePost(x *rest.X) (any, error) {
	opts := &M.RolePost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Role{}

	data.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	data.Name = opts.Name
	data.Des = opts.Des
	data.AppID = opts.AppID
	err = cfg.DB().Create(data).Error

	return data, err
}
func roleDelete(x *rest.X) (any, error) {
	opts := &M.RoleDelete{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.Role{}

	err = cfg.DB().Where("id = ?", opts.ID).Delete(data).Error

	return data, err
}
