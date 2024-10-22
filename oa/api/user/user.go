package user

import (
	"fmt"
	"math/rand"
	"oa/cfg"
	"oa/errs"
	"oa/libs/auth"
	M "oa/models"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/veypi/OneBD/rest"
	"gorm.io/gorm"
)

func useUser(r rest.Router) {
	r.Delete("/:user_id", auth.Check("user", "user_id", auth.DoDelete), userDelete)
	r.Get("/:user_id", auth.Check("user", "user_id", auth.DoRead), userGet)
	r.Get("/", auth.Check("user", "", auth.DoRead), userList)
	r.Patch("/:user_id", auth.Check("user", "user_id", auth.DoUpdate), userPatch)
	r.Post("/", userPost)
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
	data.Salt = opts.Salt
	data.Code = opts.Code
	if data.Username == "" || len(data.Salt) != 32 || len(data.Code) != 64 {
		return nil, errs.ArgsInvalid.WithStr("username/salt/code length")
	}
	if opts.Nickname != nil {
		data.Nickname = *opts.Nickname
	}
	if opts.Icon != nil {
		data.Icon = *opts.Icon
	} else {
		data.Icon = fmt.Sprintf("https://public.veypi.com/img/avatar/%04d.jpg", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(220))
	}
	if opts.Email != nil {
		data.Email = *opts.Email
	}
	if opts.Phone != nil {
		data.Phone = *opts.Phone
	}
	data.Status = 1
	err = cfg.DB().Transaction(func(tx *gorm.DB) error {
		err := tx.Create(data).Error
		if err != nil {
			return err
		}
		app := &M.App{}
		err = tx.Where("id = ?", cfg.Config.ID).First(app).Error
		if err != nil {
			return err
		}
		status := "ok"
		if app.Participate != "auto" {
			status = "applying"
		}

		return tx.Create(&M.AppUser{
			UserID: data.ID,
			AppID:  cfg.Config.ID,
			Status: status,
		}).Error
	})
	return data, nil
}
