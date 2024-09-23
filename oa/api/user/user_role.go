package user

import (
	"github.com/google/uuid"
	"github.com/veypi/OneBD/rest"
	"oa/cfg"
	M "oa/models"
	"strings"
)

func useUserRole(r rest.Router) {
	r.Delete("/:user_role_id", userRoleDelete)
	r.Get("/:user_role_id", userRoleGet)
	r.Get("/", userRoleList)
	r.Patch("/:user_role_id", userRolePatch)
	r.Post("/", userRolePost)
}
func userRoleDelete(x *rest.X) (any, error) {
	opts := &M.UserRoleDelete{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.UserRole{}

	err = cfg.DB().Where("id = ?", opts.ID).Delete(data).Error

	return data, err
}
func userRoleGet(x *rest.X) (any, error) {
	opts := &M.UserRoleGet{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.UserRole{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error

	return data, err
}
func userRoleList(x *rest.X) (any, error) {
	opts := &M.UserRoleList{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := make([]*M.UserRole, 0, 10)

	query := cfg.DB()
	if opts.Status != nil {
		query = query.Where("status LIKE ?", opts.Status)
	}
	err = query.Find(&data).Error

	return data, err
}
func userRolePatch(x *rest.X) (any, error) {
	opts := &M.UserRolePatch{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.UserRole{}

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
func userRolePost(x *rest.X) (any, error) {
	opts := &M.UserRolePost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.UserRole{}

	data.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	data.UserID = opts.UserID
	data.RoleID = opts.RoleID
	data.Status = opts.Status
	err = cfg.DB().Create(data).Error

	return data, err
}
