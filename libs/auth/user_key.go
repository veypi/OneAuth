package auth

import (
	"OneAuth/models"
	"github.com/veypi/utils"
	"sync"
)

var keyCache = sync.Map{}

func GetUserKey(uid uint, app *models.App) string {
	if app.ID == 1 {
		key, _ := keyCache.LoadOrStore(uid, utils.RandSeq(16))
		return key.(string)
	}
	// TODO: 获取其他应用user_key
	return ""
}

func RefreshUserKey(uid uint, app *models.App) string {
	if app.ID == 1 {
		key := utils.RandSeq(16)
		keyCache.Store(uid, key)
		return key
	}
	return ""
}
