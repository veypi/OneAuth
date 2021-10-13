package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/veypi/utils"
	"strings"
	"time"
)

// User db user model
type User struct {
	BaseModel
	Username  string `json:"username" gorm:"type:varchar(100);unique;not null"`
	Nickname  string `json:"nickname" gorm:"type:varchar(100)" json:",omitempty"`
	Phone     string `json:"phone" gorm:"type:varchar(20);unique;default:null" json:",omitempty"`
	Email     string `json:"email" gorm:"type:varchar(50);unique;default:null" json:",omitempty"`
	CheckCode string `gorm:"type:varchar(64);not null" json:"-"`
	RealCode  string `gorm:"type:varchar(32);not null" json:"-"`
	Position  string `json:"position"`
	// disabled 禁用
	Status string `json:"status"`

	Icon  string  `json:"icon"`
	Roles []*Role `json:"roles" gorm:"many2many:user_role;"`
	Auths []*Auth `json:"auths" gorm:"foreignkey:UserID;references:ID"`
}

type simpleAuth struct {
	RID string `json:"rid"`
	// 具体某个资源的id
	RUID  string    `json:"ruid"`
	Level AuthLevel `json:"level"`
}

// TODO:: roles 是否会造成token过大 ?
type PayLoad struct {
	ID   uint                 `json:"id"`
	Iat  int64                `json:"iat"` //token time
	Exp  int64                `json:"exp"`
	Auth map[uint]*simpleAuth `json:"auth"`
}

// GetAuth resource_uuid 缺省或仅第一个有效 权限会被更高权限覆盖
func (p *PayLoad) GetAuth(ResourceID string, ResourceUUID ...string) AuthLevel {
	res := AuthNone
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

func (u *User) String() string {
	return u.Username + ":" + u.Nickname
}

func (u *User) GetToken(key string, appID uint) (string, error) {
	header := map[string]string{
		"typ": "JWT",
		"alg": "HS256",
	}
	//header := "{\"typ\": \"JWT\", \"alg\": \"HS256\"}"
	now := time.Now().Unix()
	payload := PayLoad{
		ID:   u.ID,
		Iat:  now,
		Exp:  now + 60*60*24,
		Auth: map[uint]*simpleAuth{},
	}
	for _, r := range u.Roles {
		for _, a := range r.Auths {
			if appID == a.AppID {
				payload.Auth[a.ID] = &simpleAuth{
					RID:   a.RID,
					RUID:  a.RUID,
					Level: a.Level,
				}
			}
		}
	}
	for _, a := range u.Auths {
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
	hmacCipher := hmac.New(sha256.New, []byte(key))
	hmacCipher.Write([]byte(A + "." + B))
	C := hmacCipher.Sum(nil)
	return A + "." + B + "." + base64.StdEncoding.EncodeToString(C), nil
}

var (
	InvalidToken = errors.New("invalid token")
	ExpiredToken = errors.New("expired token")
)

func ParseToken(token string, key string, payload *PayLoad) (bool, error) {
	var A, B, C string
	if seqs := strings.Split(token, "."); len(seqs) == 3 {
		A, B, C = seqs[0], seqs[1], seqs[2]
	} else {
		return false, InvalidToken
	}
	hmacCipher := hmac.New(sha256.New, []byte(key))
	hmacCipher.Write([]byte(A + "." + B))
	tempC := hmacCipher.Sum(nil)
	if !hmac.Equal([]byte(C), []byte(base64.StdEncoding.EncodeToString(tempC))) {
		return false, nil
	}
	tempPayload, err := base64.StdEncoding.DecodeString(B)
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(tempPayload, payload); err != nil {
		return false, err
	}
	if time.Now().Unix() > payload.Exp {
		return false, ExpiredToken
	}
	return true, nil
}

func (u *User) UpdateAuth(ps string) (err error) {
	u.RealCode = utils.RandSeq(32)
	u.CheckCode, err = utils.AesEncrypt(u.RealCode, []byte(ps))
	return err
}

func (u *User) CheckLogin(ps string) (bool, error) {
	temp, err := utils.AesDecrypt(u.CheckCode, []byte(ps))
	return temp == u.RealCode, err
}
