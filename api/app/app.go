package app

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/auth"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils"
	"github.com/veypi/utils/log"
	"gorm.io/gorm"
	"reflect"
)

var appHandlerP = OneBD.NewHandlerPool(func() OneBD.Handler {
	h := &appHandler{}
	h.Ignore(rfc.MethodGet)
	return h
})

type appHandler struct {
	base.ApiHandler
	query *models.App
}

func (h *appHandler) Get() (interface{}, error) {
	uuid := h.Meta().Params("uuid")
	h.query = &models.App{}
	option := h.Meta().Query("option")
	if option == "oa" {
		// 无权限可以获取本系统基本信息
		h.query.UUID = cfg.CFG.APPUUID
		err := cfg.DB().Where(h.query).First(h.query).Error
		return h.query, err
	}
	err := h.ParsePayload(h.Meta())
	if err != nil {
		return nil, err
	}
	if option == "key" {
		if uuid == "" {
			return nil, oerr.ApiArgsError
		}
		if !h.GetAuth(auth.APP, uuid).CanDoAny() {
			return nil, oerr.NoAuth
		}
		key := utils.RandSeq(32)
		err = cfg.DB().Model(&models.App{}).Where("UUID = ?", uuid).Update("Key", key).Error
		if err != nil {
			return nil, err
		}
		return key, nil
	}
	if !h.GetAuth(auth.APP, uuid).CanRead() {
		return nil, oerr.NoAuth
	}
	if uuid != "" {
		h.query.UUID = uuid
		err := cfg.DB().Where(h.query).First(h.query).Error
		return h.query, err
	}
	// 注释代码为获取已经绑定的应用
	//user := &models.User{}
	//user.ID = h.Payload.ID
	//err := cfg.DB().Preload("Roles.Auths").Preload("Auths").Where(user).First(user).Error
	//if err != nil {
	//	return nil, oerr.DBErr.Attach(err)
	//}
	//ids := make([]string, 0, 10)
	//for _, a := range user.GetAuths() {
	//	if a.RID == auth.Login && a.Level.CanDo() {
	//		ids = append(ids, a.RUID)
	//	}
	//}
	list := make([]*models.App, 0, 10)
	err = cfg.DB().Find(&list).Error
	return list, err
}

func (h *appHandler) Post() (interface{}, error) {
	if !h.Payload.GetAuth(auth.APP, "").CanCreate() {
		return nil, oerr.NoAuth
	}
	data := &struct {
		Name string
		Icon string
	}{}
	err := h.Meta().ReadJson(data)
	if err != nil {
		return nil, err
	}
	if data.Name == "" {
		return nil, oerr.ApiArgsMissing.AttachStr("name")
	}
	a := &models.App{
		UUID:           utils.RandSeq(16),
		Name:           data.Name,
		Icon:           data.Icon,
		Creator:        h.Payload.ID,
		EnableRegister: false,
	}
	a.UUID = utils.RandSeq(16)
	err = cfg.DB().Transaction(func(tx *gorm.DB) error {
		e := tx.Create(a).Error
		if e != nil {
			return e
		}
		au := &models.AppUser{
			AppUUID: a.UUID,
			UserID:  h.Payload.ID,
			Status:  models.AUOK,
		}
		return tx.Create(au).Error
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}

func Struct2Map(obj interface{}) (data map[string]interface{}) {
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	var item reflect.Value
	var k reflect.StructField
	for i := 0; i < objT.NumField(); i++ {
		k = objT.Field(i)
		item = objV.Field(i)
		if !item.IsNil() {
			data[k.Name] = item.Interface()
		}
	}
	return
}

func (h *appHandler) Patch() (interface{}, error) {
	uid := h.Meta().Params("uuid")
	if uid == "" || !h.Payload.GetAuth(auth.APP, uid).CanUpdate() {
		return nil, oerr.NoAuth
	}
	opts := struct {
		Icon           *string
		Name           *string
		EnableRegister *bool
		Des            *string
		Host           *string
		UserRefreshUrl *string
	}{}
	if err := h.Meta().ReadJson(&opts); err != nil {
		return nil, err
	}
	query := Struct2Map(opts)
	log.Warn().Msgf("%#v", query)
	if len(query) == 0 {
		return nil, nil
	}
	if err := cfg.DB().Table("Apps").Where("UUID = ?", uid).Updates(query).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
