package user

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/app"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/libs/token"
	"github.com/veypi/OneAuth/models"
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
	userID := uint(h.Meta().ParamsInt("user_id"))
	if userID != h.Payload.ID && !h.Payload.GetAuth(auth.User, "").CanRead() {
		return nil, oerr.NoAuth.AttachStr("to read user data")
	}
	if userID != 0 {
		user := &models.User{}
		user.ID = userID
		return user, cfg.DB().Where(user).First(user).Error
	} else {
		username := h.Meta().Query("username")
		if username != "" {
			users := make([]*models.User, 0, 10)
			err := cfg.DB().Where("username LIKE ? OR nickname LIKE ?", "%"+username+"%", "%"+username+"%").Find(&users).Error
			if err != nil {
				return nil, err
			}
			return users, nil
		}
		users := make([]models.User, 10)
		skip, err := strconv.Atoi(h.Meta().Query("skip"))
		if err != nil || skip < 0 {
			skip = 0
		}
		if err := cfg.DB().Offset(skip).Find(&users).Error; err != nil {
			return nil, err
		}
		return users, nil
	}
}

// Post register user
func (h *handler) Post() (interface{}, error) {
	self := &models.App{}
	self.UUID = cfg.CFG.APPUUID
	err := cfg.DB().Where(self).First(self).Error
	if err != nil {
		return nil, oerr.DBErr.Attach(err)
	}
	log.Warn().Msgf("%v %v", self.EnableRegister, h.GetAuth(auth.User, "").CanCreate())
	if !self.EnableRegister && !h.Payload.GetAuth(auth.User, "").CanCreate() {
		return nil, errors.New("register disabled")
	}
	var userdata = struct {
		Username string
		Password string
		Nickname string
		Phone    string
		Email    string
		Domain   string
		Title    string
		Position string
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
		_, err := app.AddUser(tx, self.UUID, h.User.ID, self.InitRoleID, models.AUOK)
		if err != nil {
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
		Password string
		Icon     string
		Nickname string
		Phone    string `gorm:"type:varchar(20);unique;default:null" json:",omitempty"`
		Email    string `gorm:"type:varchar(50);unique;default:null" json:",omitempty"`
		Status   string
		Position string
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
	if err := cfg.DB().Where(&target).First(&target).Error; err != nil {
		return nil, err
	}
	if target.ID != h.Payload.ID && h.Payload.GetAuth(auth.User, strconv.Itoa(int(target.ID))).CanUpdate() {
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
	if opts.Icon != "" {
		target.Icon = opts.Icon
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
	if err := cfg.DB().Updates(&target).Error; err != nil {
		return nil, err
	}
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
	h.User = new(models.User)
	uidType := h.Meta().Query("UidType")
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
	target.UUID = cfg.CFG.APPUUID
	err = cfg.DB().Where(target).Find(target).Error
	if err != nil {
		return nil, oerr.DBErr.Attach(err)
	}
	if err := cfg.DB().Where(h.User).First(h.User).Error; err != nil {
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
	au.AppUUID = target.UUID
	err = cfg.DB().Where(au).First(au).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if !target.EnableRegister {
			return nil, errors.New("该应用不允许自主加入，请联系管理员")
		}
		_, err := app.AddUser(cfg.DB(), target.UUID, h.User.ID, target.InitRoleID, models.AUOK)
		if err != nil {
			return nil, err
		}
	} else if au.Status != models.AUOK {
		return nil, oerr.DisableLogin
	}
	err = h.User.LoadAuths(cfg.DB())
	if err != nil {
		return nil, err
	}
	tokenStr, err := token.GetToken(h.User, target.UUID, cfg.CFG.APPKey)
	if err != nil {
		log.HandlerErrs(err)
		return nil, oerr.Unknown.Attach(err)
	}
	h.Meta().SetHeader("auth_token", tokenStr)
	log.Info().Msg(h.User.Username + " login")
	return nil, nil
}
