package file

import (
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/fs"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils"
	"github.com/veypi/utils/log"
)

/**
* @name: main
* @author: veypi <i@veypi.com>
* @date: 2021-11-27 16:27
* @description：main
**/

func Router(r OneBD.Router) {
	// 用户私有文件
	usrF := fs.FS{
		RootDir:   utils.PathJoin(cfg.CFG.FireDir, "usr"),
		UrlRoot:   utils.PathJoin(cfg.CFG.FileUrlPrefix, "usr"),
		MountFunc: userFileChecker,
	}
	log.Info().Msgf("start file server on %s", cfg.CFG.Host)
	r.Set("/usr", usrF.ServeHTTP, rfc.MethodAll)
	r.Set("/usr/*", usrF.ServeHTTP, rfc.MethodAll)
	// 应用存储文件
	appF := fs.FS{
		RootDir:       utils.PathJoin(cfg.CFG.FireDir, "app"),
		UrlRoot:       utils.PathJoin(cfg.CFG.FileUrlPrefix, "app"),
		MountFunc:     appFileChecker,
		DisabledAlert: true,
	}
	r.Set("/app/:uuid/", appF.ServeHTTP, rfc.MethodAll)
	r.Set("/app/:uuid/*prefix", appF.ServeHTTP, rfc.MethodAll)
	r.Set("/ursapp/:id/:uuid/*prefix", nil)
	// 公共文件 读取无需权限
	pubF := fs.FS{
		RootDir:       utils.PathJoin(cfg.CFG.FireDir, "public"),
		UrlRoot:       utils.PathJoin(cfg.CFG.FileUrlPrefix, "public"),
		MountFunc:     pubFileChecker,
		DisabledAlert: true,
	}
	r.Set("/public/app/:uuid/*prefix", pubF.ServeHTTP, rfc.MethodAll)
	r.Set("/public/*prefix", pubF.ServeHTTP, rfc.MethodAll)
}
