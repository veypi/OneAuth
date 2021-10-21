package user

import (
	"OneAuth/cfg"
	"OneAuth/libs/app"
	"OneAuth/libs/base"
	"OneAuth/libs/oerr"
	"OneAuth/libs/token"
	"OneAuth/models"
	"errors"

	//"OneAuth/ws"
	"encoding/base64"
	"fmt"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils/log"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

func Router(r OneBD.Router) {
	pool := OneBD.NewHandlerPool(func() OneBD.Handler {
		h := &handler{}
		h.Ignore(rfc.MethodHead, rfc.MethodPost)
		return h
	})
	r.Set("/", pool, rfc.MethodGet, rfc.MethodPost) // list
	r.Set("/:user_id", pool, rfc.MethodGet, rfc.MethodPatch, rfc.MethodHead, rfc.MethodDelete)
	r.Set("/:user_id/role/", userRoleP, rfc.MethodPost)
	r.Set("/:user_id/role/:role_id", userRoleP, rfc.MethodDelete)
	//r.WS("/ws", func(m OneBD.Meta) (conn OneBD.WebsocketConn, err error) {
	//return ws.User.Upgrade(m.ResponseWriter(), m.Request())
	//})
}

type handler struct {
	base.ApiHandler
	User *models.User
}

// Get get user data
func (h *handler) Get() (interface{}, error) {
	username := h.Meta().Query("username")
	if username != "" {
		users := make([]*models.User, 0, 10)
		err := cfg.DB().Preload("Scores").Preload("Roles.Auths").Where("username LIKE ? OR nickname LIKE ?", "%"+username+"%", "%"+username+"%").Find(&users).Error
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	userID := h.Meta().ParamsInt("user_id")
	if userID != 0 {
		user := &models.User{}
		user.ID = uint(userID)
		return user, cfg.DB().Where(user).Preload("Scores").Preload("Roles.Auths").Preload("Favorites").First(user).Error
	} else {
		users := make([]models.User, 10)
		skip, err := strconv.Atoi(h.Meta().Query("skip"))
		if err != nil || skip < 0 {
			skip = 0
		}
		if err := cfg.DB().Preload("Scores").Preload("Roles.Auths").Offset(skip).Find(&users).Error; err != nil {
			return nil, err
		}
		return users, nil
	}
}

// Post register user
func (h *handler) Post() (interface{}, error) {
	self := &models.App{}
	self.ID = cfg.CFG.APPID
	err := cfg.DB().Where(self).First(self).Error
	if err != nil {
		return nil, oerr.DBErr.Attach(err)
	}
	if !self.EnableRegister {
		return nil, oerr.NoAuth.AttachStr("register disabled")
	}
	var userdata = struct {
		UUID     string `json:"uuid"`
		Username string `json:"username"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Domain   string `json:"domain"`
		Title    string `json:"title"`
		Position string `json:"position"`
	}{}
	if err := h.Meta().ReadJson(&userdata); err != nil {
		return nil, err
	}
	pass, err := base64.StdEncoding.DecodeString(userdata.Password)
	if err != nil {
		return nil, err
	}

	if len(pass) > 32 || len(pass) < 6 {
		return nil, oerr.PassError
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	h.User = new(models.User)
	h.User.Icon = fmt.Sprintf("/media/icon/default/%04d.jpg", r.Intn(230))
	h.User.Nickname = userdata.Nickname
	h.User.Phone = userdata.Phone
	h.User.Username = userdata.Username
	h.User.Email = userdata.Email
	h.User.Position = userdata.Position
	if err := h.User.UpdatePass(string(pass)); err != nil {
		log.HandlerErrs(err)
		return nil, oerr.ResourceCreatedFailed
	}
	err = cfg.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&h.User).Error; err != nil {
			return oerr.ResourceDuplicated
		}
		err := app.AddUser(tx, self.ID, h.User.ID, self.InitRoleID)
		if err != nil {
			return err
		}
		if userdata.UUID != self.UUID && userdata.UUID != "" {
			target := &models.App{}
			target.UUID = userdata.UUID
			err = tx.Where(target).First(target).Error
			if err != nil {
				return err
			}
			if target.EnableRegister {
				err := app.AddUser(tx, target.ID, h.User.ID, target.InitRoleID)
				if err != nil {
					return err
				}
			} else {
				// TODO
			}
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.User, nil
}

// Patch update user data
func (h *handler) Patch() (interface{}, error) {
	uid := h.Meta().Params("user_id")
	opts := struct {
		Password string `json:"password"`
		Nickname string `json:"nickname"`
		Phone    string `json:"phone" gorm:"type:varchar(20);unique;default:null" json:",omitempty"`
		Email    string `json:"email" gorm:"type:varchar(50);unique;default:null" json:",omitempty"`
		Status   string `json:"status"`
		Position string `json:"position"`
	}{}
	if err := h.Meta().ReadJson(&opts); err != nil {
		return nil, err
	}
	target := models.User{}
	if tempID, err := strconv.Atoi(uid); err != nil || tempID <= 0 {
		return nil, oerr.ApiArgsError.Attach(err)
	} else {
		target.ID = uint(tempID)
	}
	tx := cfg.DB().Begin()
	if err := cfg.DB().Where(&target).First(&target).Error; err != nil {
		return nil, err
	}
	if target.ID != h.Payload.ID {
		return nil, oerr.NoAuth
	}
	if len(opts.Password) >= 6 {
		if err := target.UpdatePass(opts.Password); err != nil {
			log.HandlerErrs(err)
			return nil, oerr.ApiArgsError.AttachStr(err.Error())
		}
	}
	if opts.Nickname != "" {
		target.Nickname = opts.Nickname
	}
	if opts.Position != "" {
		target.Position = opts.Position
	}
	if opts.Phone != "" {
		target.Phone = opts.Phone
	}
	if opts.Email != "" {
		target.Email = opts.Email
	}
	if opts.Status != "" {
		target.Status = opts.Status
	}
	if err := tx.Updates(&target).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return nil, nil
}

// Delete delete user
func (h *handler) Delete() (interface{}, error) {
	// TODO::
	return nil, nil
}

// Head user login
func (h *handler) Head() (interface{}, error) {
	uid := h.Meta().Params("user_id")
	pass, err := base64.StdEncoding.DecodeString(h.Meta().Query("password"))
	if err != nil {
		return nil, oerr.ApiArgsError.Attach(err)
	}
	password := string(pass)
	if len(uid) == 0 || len(password) == 0 {
		return nil, oerr.ApiArgsError
	}
	appUUID := h.Meta().Query("uuid")
	if appUUID == "" {
		return nil, oerr.ApiArgsMissing.AttachStr("uuid")
	}
	h.User = new(models.User)
	uidType := h.Meta().Query("uid_type")
	switch uidType {
	case "username":
		h.User.Username = uid
	case "phone":
		h.User.Phone = uid
	case "email":
		h.User.Email = uid
	default:
		h.User.Username = uid
	}
	target := &models.App{}
	target.UUID = appUUID
	err = cfg.DB().Where(target).Find(target).Error
	if err != nil {
		return nil, oerr.DBErr.Attach(err)
	}
	if err := cfg.DB().Preload("Roles.Auths").Preload("Auths").Where(h.User).First(h.User).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, oerr.AccountNotExist
		} else {
			log.HandlerErrs(err)
			return nil, oerr.DBErr.Attach(err)
		}
	}
	isAuth, err := h.User.CheckLogin(password)
	if err != nil || !isAuth {
		return nil, oerr.PassError.Attach(err)
	}
	au := &models.AppUser{}
	au.UserID = h.User.ID
	au.AppID = target.ID
	err = cfg.DB().Where(au).First(au).Error
	appID := target.ID
	h.Meta().SetHeader("content", target.UserRefreshUrl)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		appID = cfg.CFG.APPID
		h.Meta().SetHeader("content", "/app/"+target.UUID)
	} else if au.Disabled {
		return nil, oerr.DisableLogin
	}
	tokenStr, err := token.GetToken(h.User, appID)
	if err != nil {
		log.HandlerErrs(err)
		return nil, oerr.Unknown.Attach(err)
	}
	h.Meta().SetHeader("auth_token", tokenStr)
	log.Info().Msg(h.User.Username + " login")
	return nil, nil
}
