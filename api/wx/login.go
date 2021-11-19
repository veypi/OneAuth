package wx

import (
	"errors"
	"fmt"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/tools"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils"
	"github.com/veypi/utils/log"
	"net/url"
	"strings"
	"time"
)

var tokens = map[uint]string{
	1: "",
}

func login(m OneBD.Meta) {
	var loc = ""
	defer func() {
		m.Header().Set("Location", loc)
		log.Warn().Msg(loc)
		m.WriteHeader(rfc.StatusPermanentRedirect)
	}()
	app := &models.App{
		UUID: m.Params("id"),
	}
	err := cfg.DB().Preload("Wx").Where(app).First(app).Error
	loc = fmt.Sprintf("/#/wx?uuid=%s&msg=", app.UUID)
	if err != nil {
		loc += err.Error()
		return
	}
	if app.Wx == nil {
		loc += "微信登录未绑定"
		return
	}
	if tokens[app.Wx.ID] == "" {
		tokens[app.Wx.ID], err = requestCorpToken(app.Wx.CorpID, app.Wx.CorpSecret)
		if err != nil {
			log.Warn().Msg("get corp token failed: " + err.Error())
			loc += err.Error()
			return
		}
	}
	user, err := getUserID(tokens[app.Wx.ID], m.Query("code"))
	if err != nil {
		if strings.Contains(err.Error(), "access_token expired") {
			tokens[app.Wx.ID], err = requestCorpToken(app.Wx.CorpID, app.Wx.CorpSecret)
			if err != nil {
				log.Warn().Msg("refresh corp token failed: " + err.Error())
				loc += err.Error()
				return
			}
			user, err = getUserID(tokens[app.Wx.ID], m.Query("code"))
			if err != nil {
				log.Warn().Msg("get user token failed: " + err.Error())
				loc += err.Error()
				return
			}
		} else {
			log.Warn().Msg("get user token failed: " + err.Error())
			loc += err.Error()
			return
		}
	}
	info, err := getUserInfo(tokens[app.Wx.ID], user)
	if err != nil {
		log.Warn().Msg("get user info failed: " + err.Error())
		loc += err.Error()
		return
	}
	log.Warn().Msgf("\ncode= %s\nstate= %s\nu = %s\n%v",
		m.Query("code"), m.Query("state"), user, info)
	pass, err := utils.AesEncrypt(fmt.Sprintf("%s.%d", user, time.Now().Unix()), []byte(app.UUID))
	if err != nil {
		loc += err.Error()
		return
	}
	log.Warn().Msgf("pass: %s", pass)
	v := url.Values{}
	v.Add("wid", pass)
	u, err := url.Parse(app.Host)
	u.RawQuery = v.Encode()
	if err != nil {
		loc += err.Error()
		return
	}
	loc = u.String()
}

func requestCorpToken(corpid, corpsecret string) (string, error) {
	addr := "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	query := map[string]string{
		"corpid":     corpid,
		"corpsecret": corpsecret,
	}
	res := &struct {
		Errmsg      string
		Errcode     *uint
		AccessToken string
	}{}
	err := tools.Query(addr, query, res)
	if err != nil {
		return "", errors.New("request token response json parse err :" + err.Error())
	}
	if res.Errcode != nil && *res.Errcode == 0 {
		return res.AccessToken, nil
	} else {
		//返回错误信息
		err = errors.New(fmt.Sprintf("%d:%s", res.Errcode, res.Errmsg))
		return "", err
	}
}

func getUserID(token, code string) (string, error) {
	addr := "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
	res := &struct {
		Errmsg   string
		Errcode  *uint
		UserId   string
		DeviceId string
	}{}
	query := map[string]string{
		"access_token": token,
		"code":         code,
	}
	err := tools.Query(addr, query, res)
	if err != nil {
		return "", err
	}
	if res.Errcode != nil && *res.Errcode == 0 {
		return res.UserId, nil
	}
	return "", errors.New(fmt.Sprintf("%d:%s", res.Errcode, res.Errmsg))
}

func getUserInfo(token, id string) (interface{}, error) {
	addr := "https://qyapi.weixin.qq.com/cgi-bin/user/get"
	res := map[string]interface{}{}
	query := map[string]string{
		"access_token": token,
		"userid":       id,
	}
	err := tools.Query(addr, query, &res)
	if err != nil {
		return "", err
	}
	errcode := int(res["errcode"].(float64))
	errmsg := res["errmsg"].(string)
	if errcode == 0 {
		log.Warn().Msgf("%v", res)
		return res, nil
	}
	return "", errors.New(fmt.Sprintf("%d:%s", errcode, errmsg))
}
