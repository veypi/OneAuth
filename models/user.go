package models

import (
	"github.com/veypi/OneAuth/oalib"
	"github.com/veypi/utils"
	"gorm.io/gorm"
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
	Roles []*Role `json:"roles" gorm:"many2many:user_roles;"`
	Apps  []*App  `json:"apps" gorm:"many2many:app_users;"`
	Auths []*Auth `json:"auths" gorm:"foreignkey:UserID;references:ID"`
}

func (u *User) String() string {
	return u.Username + ":" + u.Nickname
}

func (u *User) LoadAuths(tx *gorm.DB) error {
	return tx.Where("id = ?", u.ID).Preload("Auths").Preload("Roles.Auths").First(u).Error
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
