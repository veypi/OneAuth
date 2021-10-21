package token

import (
	"OneAuth/libs/key"
	"OneAuth/models"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var (
	InvalidToken = errors.New("invalid token")
	ExpiredToken = errors.New("expired token")
)

type simpleAuth struct {
	RID string `json:"rid"`
	// 具体某个资源的id
	RUID  string           `json:"ruid"`
	Level models.AuthLevel `json:"level"`
}

// TODO:: roles 是否会造成token过大 ?
type PayLoad struct {
	ID    uint                 `json:"id"`
	AppID uint                 `json:"app_id"`
	Iat   int64                `json:"iat"` //token time
	Exp   int64                `json:"exp"`
	Auth  map[uint]*simpleAuth `json:"auth"`
}

// GetAuth resource_uuid 缺省或仅第一个有效 权限会被更高权限覆盖
func (p *PayLoad) GetAuth(ResourceID string, ResourceUUID ...string) models.AuthLevel {
	res := models.AuthNone
	if p == nil || p.Auth == nil {
		return res
	}
	ruid := ""
	if len(ResourceUUID) > 0 {
		ruid = ResourceUUID[0]
	}
	for _, a := range p.Auth {
		if a.RID == ResourceID {
			if a.RUID != "" {
				if a.RUID == ruid {
					if a.Level > res {
						res = a.Level
					}
				} else {
					continue
				}
			} else if a.Level > res {
				res = a.Level
			}
		}
	}
	return res
}

func GetToken(u *models.User, appID uint) (string, error) {
	header := map[string]string{
		"typ": "JWT",
		"alg": "HS256",
	}
	//header := "{\"typ\": \"JWT\", \"alg\": \"HS256\"}"
	now := time.Now().Unix()
	payload := PayLoad{
		ID:    u.ID,
		AppID: appID,
		Iat:   now,
		Exp:   now + 60*60*24,
		Auth:  map[uint]*simpleAuth{},
	}
	for _, a := range u.GetAuths() {
		if appID == a.AppID {
			payload.Auth[a.ID] = &simpleAuth{
				RID:   a.RID,
				RUID:  a.RUID,
				Level: a.Level,
			}
		}
	}
	a, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	A := base64.StdEncoding.EncodeToString(a)
	B := base64.StdEncoding.EncodeToString(b)
	hmacCipher := hmac.New(sha256.New, []byte(key.User(payload.ID, payload.AppID)))
	hmacCipher.Write([]byte(A + "." + B))
	C := hmacCipher.Sum(nil)
	return A + "." + B + "." + base64.StdEncoding.EncodeToString(C), nil
}

func ParseToken(token string, payload *PayLoad) (bool, error) {
	var A, B, C string
	if seqs := strings.Split(token, "."); len(seqs) == 3 {
		A, B, C = seqs[0], seqs[1], seqs[2]
	} else {
		return false, InvalidToken
	}
	tempPayload, err := base64.StdEncoding.DecodeString(B)
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(tempPayload, payload); err != nil {
		return false, err
	}
	hmacCipher := hmac.New(sha256.New, []byte(key.User(payload.ID, payload.AppID)))
	hmacCipher.Write([]byte(A + "." + B))
	tempC := hmacCipher.Sum(nil)
	if !hmac.Equal([]byte(C), []byte(base64.StdEncoding.EncodeToString(tempC))) {
		return false, nil
	}
	if time.Now().Unix() > payload.Exp {
		return false, ExpiredToken
	}
	return true, nil
}
