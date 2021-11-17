package token

import (
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/oalib"
	"github.com/veypi/utils/jwt"
)

func GetToken(u *models.User, uuid string, key string) (string, error) {
	payload := &oalib.PayLoad{
		ID:   u.ID,
		Auth: []*oalib.SimpleAuth{},
	}
	for _, a := range u.GetAuths() {
		if uuid == a.AppUUID {
			payload.Auth = append(payload.Auth, &oalib.SimpleAuth{
				RID:   a.RID,
				RUID:  a.RUID,
				Level: a.Level,
			})
		}
	}
	return jwt.GetToken(payload, []byte(key))
}

func ParseToken(token string, payload *oalib.PayLoad, key string) (bool, error) {
	return jwt.ParseToken(token, payload, []byte(key))
}
