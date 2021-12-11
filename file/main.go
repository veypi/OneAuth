package file

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/fs"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils/log"
)

/**
* @name: main
* @author: veypi <i@veypi.com>
* @date: 2021-11-27 16:27
* @description：main
**/

func Router(r OneBD.Router) {
	usrF := fs.FS{
		RootDir:   cfg.CFG.FireDir + "/usr/",
		MountFunc: userFileChecker,
	}
	log.Info().Msgf("start file server on %s", cfg.CFG.Host)
	r.Set("/usr/", usrF.ServeHTTP, rfc.MethodAll)
	r.Set("/usr/*", usrF.ServeHTTP, rfc.MethodAll)
	appF := fs.FS{
		RootDir:       cfg.CFG.FireDir + "/app/",
		MountFunc:     appFileChecker,
		DisabledAlert: true,
	}
	r.Set("/app/:uuid/", appF.ServeHTTP, rfc.MethodAll)
	r.Set("/app/:uuid/*preifx", appF.ServeHTTP, rfc.MethodAll)
	r.Set("/ursapp/:id/:uuid/*prefix", nil)
}
