package oalib

/**
* @name: payload
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 16:45
* @description：payload
**/

import (
	"github.com/veypi/utils/jwt"
)

type SimpleAuth struct {
	RID string
	// 具体某个资源的id
	RUID  string
	Level AuthLevel
}

// PayLoad TODO:: roles 是否会造成token过大 ?
type PayLoad struct {
	jwt.Payload
	ID   uint
	Auth []*SimpleAuth
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

func (p *PayLoad) ParseToken(token string, key string) (bool, error) {
	return jwt.ParseToken(token, p, []byte(key))
}
