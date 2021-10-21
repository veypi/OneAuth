package key

import (
	"OneAuth/cfg"
	"github.com/veypi/utils"
	"sync"
)

var keyCache = sync.Map{}

func User(uid uint, appID uint) string {
	if appID == cfg.CFG.APPID {
		key, _ := keyCache.LoadOrStore(uid, utils.RandSeq(16))
		return cfg.CFG.APPKey + key.(string)
	}
	// TODO: 获取其他应用user_key
	return ""
}

func RefreshUser(uid uint, appID uint) string {
	if appID == cfg.CFG.APPID {
		key := utils.RandSeq(16)
		keyCache.Store(uid, key)
		return key
	}
	return ""
}
