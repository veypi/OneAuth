package key

import "github.com/veypi/OneAuth/cfg"

func App(uuid string) string {
	if uuid == cfg.CFG.APPUUID {
		return cfg.CFG.APPKey
	}
	// TODO
	return ""
}
