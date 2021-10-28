package token

import (
	"OneAuth/libs/key"
	"OneAuth/models"
	"github.com/veypi/utils/jwt"
)

type simpleAuth struct {
	RID string `json:"rid"`
	// 具体某个资源的id
	RUID  string           `json:"ruid"`
	Level models.AuthLevel `json:"level"`
}

// TODO:: roles 是否会造成token过大 ?
type PayLoad struct {
	jwt.Payload
	ID    uint                 `json:"id"`
	AppID uint                 `json:"app_id"`
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
	payload := &PayLoad{
		ID:    u.ID,
		AppID: appID,
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
	return jwt.GetToken(payload, []byte(key.User(payload.ID, payload.AppID)))
}

func ParseToken(token string, payload *PayLoad) (bool, error) {
	return jwt.ParseToken(token, payload, []byte(key.User(payload.ID, payload.AppID)))
}
