package key

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/utils"
	"sync"
)

var keyCache = sync.Map{}

func User(uid uint, uuid string) string {
	if uuid == cfg.CFG.APPUUID {
		key, _ := keyCache.LoadOrStore(uid, utils.RandSeq(16))
		return cfg.CFG.APPKey + key.(string)
	}
	// TODO: 获取其他应用user_key
	return ""
}

func RefreshUser(uid uint, uuid string) string {
	if uuid == cfg.CFG.APPUUID {
		key := utils.RandSeq(16)
		keyCache.Store(uid, key)
		return key
	}
	return ""
}
