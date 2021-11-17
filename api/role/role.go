package role

import (
	"OneAuth/cfg"
	"OneAuth/libs/auth"
	"OneAuth/libs/base"
	"OneAuth/libs/oerr"
	"OneAuth/models"
	"github.com/veypi/OneBD"
	"gorm.io/gorm"
)

var roleP = OneBD.NewHandlerPool(func() OneBD.Handler {
	return &roleHandler{}
})

type baseAppHandler struct {
	base.ApiHandler
	uuid string
}

func (h *baseAppHandler) Init(m OneBD.Meta) error {
	h.uuid = m.Params("uuid")
	return h.ApiHandler.Init(m)
}

type roleHandler struct {
	baseAppHandler
}

func (h *roleHandler) Get() (interface{}, error) {
	id := h.Meta().ParamsInt("id")
	if !h.GetAuth(auth.Role, h.uuid).CanRead() {
		return nil, oerr.NoAuth
	}
	if id > 0 {
		role := &models.Role{}
		role.AppUUID = h.uuid
		role.ID = uint(id)
		err := cfg.DB().Preload("Auths").Preload("Users").First(role).Error
		if err != nil {
			return nil, err
		}
		return role, nil
	}
	roles := make([]*models.Role, 0, 10)
	err := cfg.DB().Where("app_uuid = ?", h.uuid).Find(&roles).Error
	return roles, err
}

func (h *roleHandler) Post() (interface{}, error) {
	if !h.GetAuth(auth.Role).CanCreate() {
		return nil, oerr.NoAuth
	}
	role := &models.Role{}
	err := h.Meta().ReadJson(role)
	if err != nil {
		return nil, err
	}
	role.ID = 0
	if role.Name == "" {
		return nil, oerr.ApiArgsMissing
	}
	return role, cfg.DB().Where(role).FirstOrCreate(role).Error
}

func (h *roleHandler) Patch() (interface{}, error) {
	if !h.GetAuth(auth.Role).CanUpdate() {
		return nil, oerr.NoAuth
	}
	query := &struct {
		Name *string `json:"name"`
		// 角色标签
		Tag *string `json:"tag" gorm:"default:''"`
	}{}
	err := h.Meta().ReadJson(query)
	if err != nil {
		return nil, err
	}
	rid := h.Meta().ParamsInt("id")
	if rid <= 0 {
		return nil, oerr.ApiArgsError
	}
	role := &models.Role{}
	role.ID = uint(rid)
	err = cfg.DB().Preload("Users").Where(role).First(role).Error
	if err != nil {
		return nil, err
	}
	return nil, cfg.DB().Transaction(func(tx *gorm.DB) error {
		var err error
		if query.Tag != nil && *query.Tag != role.Tag {
			err = tx.Model(role).Update("tag", *query.Tag).Error
			if err != nil {
				return err
			}
		}
		if query.Name != nil && *query.Name != role.Name {
			err = tx.Model(role).Update("name", *query.Name).Error
			if err != nil {
				return err
			}
		}
		return err
	})
}

func (h *roleHandler) Delete() (interface{}, error) {
	if !h.GetAuth(auth.Role).CanDelete() {
		return nil, oerr.NoAuth
	}
	rid := h.Meta().ParamsInt("id")
	if rid <= 2 {
		return nil, oerr.NoAuth
	}
	role := &models.Role{}
	role.ID = uint(rid)
	err := cfg.DB().Where(role).First(role).Error
	if err != nil {
		return nil, err
	}
	return nil, cfg.DB().Delete(role).Error
}
