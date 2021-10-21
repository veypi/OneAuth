package key

import "OneAuth/cfg"

func App(id uint) string {
	if id == cfg.CFG.APPID {
		return cfg.CFG.APPKey
	}
	// TODO
	return ""
}
