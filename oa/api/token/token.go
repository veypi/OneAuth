package token

import (
	"encoding/hex"
	"oa/cfg"
	"oa/errs"
	"oa/libs/auth"
	M "oa/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/veypi/OneBD/rest"
	"github.com/veypi/utils"
	"github.com/veypi/utils/logv"
)

func useToken(r rest.Router) {
	r.Post("/salt", tokenSalt)
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
	query := "username = ?"
	if opts.Typ == nil {
	} else if *opts.Typ == "email" {
		query = "email = ?"
	} else if *opts.Typ == "phone" {
		query = "phone = ?"
	}

	err = cfg.DB().Where(query, opts.Username).First(data).Error
	return map[string]string{"salt": data.Salt, "id": data.ID}, err
}

// for user login app
func tokenPost(x *rest.X) (any, error) {
	opts := &M.TokenPost{}
	err := x.Parse(opts)
	if err != nil {
		return nil, err
	}
	aid := cfg.Config.ID
	if opts.AppID != nil {
		aid = *opts.AppID
	}
	data := &M.Token{}
	claim := &auth.Claims{}
	claim.IssuedAt = jwt.NewNumericDate(time.Now())
	claim.Issuer = "oa"
	if opts.Token != nil {
		// for other app redirect
		oldClaim, err := auth.ParseJwt(*opts.Token)
		if err != nil {
			return nil, err
		}
		err = cfg.DB().Where("id = ?", oldClaim.ID).First(data).Error
		if err != nil {
			return nil, err
		}
		if oldClaim.AID == aid {
			// refresh token
			claim.AID = oldClaim.AID
			claim.UID = oldClaim.UID
			claim.Name = oldClaim.Name
			claim.Icon = oldClaim.Icon
			claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 10))
			acList := make(auth.Access, 0, 10)
			logv.AssertError(cfg.DB().Table("accesses a").
				Select("a.name, a.t_id, a.level").
				Joins("INNER JOIN user_roles ur ON ur.role_id = a.role_id AND ur.user_id = ?", oldClaim.UID).
				Scan(&acList).Error)
			claim.Access = acList
		} else {
			// gen other app token
		}
	} else if opts.Code != nil && aid == cfg.Config.ID {
		// for oa login
		user := &M.User{}
		err = cfg.DB().Where("id = ?", opts.UserID).Find(user).Error
		if err != nil {
			return nil, err
		}
		logv.Info().Str("user", user.ID).Msg("login")
		code := *opts.Code
		salt := logv.AssertFuncErr(hex.DecodeString(*opts.Salt))
		key := logv.AssertFuncErr(hex.DecodeString(user.Code))
		de, err := utils.AesDecrypt([]byte(code), key, salt)
		if err != nil || de != user.ID {
			return nil, errs.AuthFailed
		}
		data.UserID = opts.UserID
		data.AppID = aid
		if opts.ExpiredAt != nil {
			data.ExpiredAt = *opts.ExpiredAt
		} else {
			data.ExpiredAt = time.Now().Add(time.Hour * 72)
		}
		if opts.OverPerm != nil {
			data.OverPerm = *opts.OverPerm
		}
		if opts.Device != nil {
			data.Device = *opts.Device
		}
		data.Ip = x.GetRemoteIp()
		data.ExpiredAt = time.Now().Add(time.Hour)
		logv.AssertError(cfg.DB().Create(data).Error)
		claim.ID = data.ID
		claim.AID = aid
		claim.UID = user.ID
		claim.Name = user.Username
		claim.Icon = user.Icon
		claim.ExpiresAt = jwt.NewNumericDate(data.ExpiredAt)
		if user.Nickname != "" {
			claim.Name = user.Nickname
		}
	} else {
		return nil, errs.ArgsInvalid
	}

	token := logv.AssertFuncErr(auth.GenJwt(claim))
	return token, err
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
