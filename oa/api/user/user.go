package user

import (
	"github.com/veypi/OneBD/rest"
	M "oa/models"
	"oa/cfg"
	"strings"
	"github.com/google/uuid"
)

func useUser(r rest.Router) {
	r.Delete("/:user_id", userDelete)
	r.Get("/:user_id", userGet)
	r.Get("/", userList)
	r.Patch("/:user_id", userPatch)
	r.Post("/", userPost)
	r.Put("/:user_id", userPut)
}
func userDelete(x *rest.X) (any, error) {
	opts := &M.UserDelete{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.User{}

	err = cfg.DB().Where("id = ?", opts.ID).Delete(data).Error

	return data, err
}
func userGet(x *rest.X) (any, error) {
	opts := &M.UserGet{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.User{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error

	return data, err
}
func userList(x *rest.X) (any, error) {
	opts := &M.UserList{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := make([]*M.User, 0, 10)

	query := cfg.DB()
	if opts.Username != nil {
		query = query.Where("username LIKE ?", opts.Username)
	}
	if opts.Nickname != nil {
		query = query.Where("nickname LIKE ?", opts.Nickname)
	}
	if opts.Email != nil {
		query = query.Where("email LIKE ?", opts.Email)
	}
	if opts.Phone != nil {
		query = query.Where("phone LIKE ?", opts.Phone)
	}
	if opts.Status != nil {
		query = query.Where("status = ?", opts.Status)
	}
	err = query.Find(&data).Error

	return data, err
}
func userPatch(x *rest.X) (any, error) {
	opts := &M.UserPatch{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.User{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error
	if err != nil {
		return nil, err
	}
	optsMap := make(map[string]interface{})
	if opts.Username != nil {
		optsMap["username"] = opts.Username
	}
	if opts.Nickname != nil {
		optsMap["nickname"] = opts.Nickname
	}
	if opts.Icon != nil {
		optsMap["icon"] = opts.Icon
	}
	if opts.Email != nil {
		optsMap["email"] = opts.Email
	}
	if opts.Phone != nil {
		optsMap["phone"] = opts.Phone
	}
	if opts.Status != nil {
		optsMap["status"] = opts.Status
	}
	err = cfg.DB().Model(data).Updates(optsMap).Error

	return data, err
}
func userPost(x *rest.X) (any, error) {
	opts := &M.UserPost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.User{}

	data.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	data.Username = opts.Username
	data.Nickname = opts.Nickname
	data.Icon = opts.Icon
	data.Email = opts.Email
	data.Phone = opts.Phone
	data.Status = opts.Status
	err = cfg.DB().Create(data).Error

	return data, err
}
func userPut(x *rest.X) (any, error) {
	opts := &M.UserPut{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	data := &M.User{}

	err = cfg.DB().Where("id = ?", opts.ID).First(data).Error
	if err != nil {
		return nil, err
	}
	optsMap := map[string]interface{}{
		"id":		opts.ID,
		"username":	opts.Username,
		"nickname":	opts.Nickname,
		"icon":		opts.Icon,
		"email":	opts.Email,
		"phone":	opts.Phone,
		"status":	opts.Status,
	}
	err = cfg.DB().Model(data).Updates(optsMap).Error

	return data, err
}
