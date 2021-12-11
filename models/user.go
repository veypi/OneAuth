package models

import (
	"github.com/veypi/OneAuth/oalib"
	"github.com/veypi/utils"
	"github.com/veypi/utils/jwt"
	"gorm.io/gorm"
)

// User db user model
type User struct {
	BaseModel
	Username  string `gorm:"type:varchar(100);unique;not null"`
	Nickname  string `gorm:"type:varchar(100)"`
	Phone     string `gorm:"type:varchar(20);unique;default:null"`
	Email     string `gorm:"type:varchar(50);unique;default:null"`
	CheckCode string `gorm:"type:varchar(64);not null" json:"-"`
	RealCode  string `gorm:"type:varchar(32);not null" json:"-"`
	Position  string
	// disabled 禁用
	Status string

	Icon  string
	Roles []*Role `gorm:"many2many:UserRoles;"`
	Apps  []*App  `gorm:"many2many:AppUsers;"`
	Auths []*Auth `gorm:"foreignkey:UserID;references:ID"`
	Used  uint    `gorm:"default:0"`
	Space uint    `gorm:"default:300"`
}

func (u *User) String() string {
	return u.Username + ":" + u.Nickname
}

func (u *User) LoadAuths(tx *gorm.DB) error {
	return tx.Where("ID = ?", u.ID).Preload("Auths").Preload("Roles.Auths").First(u).Error
}

func (u *User) GetAuths() []*Auth {
	list := make([]*Auth, 0, 10)
	for _, r := range u.Roles {
		for _, a := range r.Auths {
			list = append(list, a)
		}
	}
	for _, a := range u.Auths {
		list = append(list, a)
	}
	return list
}

func (u *User) GetAuth(uuid, ResourceID string, ResourceUUID ...string) oalib.AuthLevel {
	var res = oalib.AuthNone
	ruid := ""
	if len(ResourceUUID) > 0 {
		ruid = ResourceUUID[0]
	}
	for _, a := range u.GetAuths() {
		if a.RID == ResourceID && a.AppUUID == uuid {
			if a.RUID != "" {
				if a.RUID == ruid {
					if a.Level.Upper(res) {
						res = a.Level
					}
				} else {
					continue
				}
			} else if a.Level.Upper(res) {
				res = a.Level
			}
		}
	}
	return res
}

func (u *User) UpdatePass(ps string) (err error) {
	u.RealCode = utils.RandSeq(32)
	u.CheckCode, err = utils.AesEncrypt(u.RealCode, []byte(ps))
	return err
}

func (u *User) CheckLogin(ps string) (bool, error) {
	temp, err := utils.AesDecrypt(u.CheckCode, []byte(ps))
	return temp == u.RealCode, err
}

func (u *User) GetToken(uuid string, key string) (string, error) {
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
