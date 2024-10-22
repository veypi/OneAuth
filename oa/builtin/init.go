//
// init.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-18 17:07
// Distributed under terms of the GPL license.
//

package builtin

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"oa/builtin/fs"
	"oa/cfg"

	"github.com/veypi/OneBD/rest"
	"github.com/veypi/utils/logv"
)

func Enable(app *rest.Application) {
	if cfg.Config.FsPath != "" {
		r := app.Router().SubRouter("fs")
		r.Any("/app/*", fs.NewAppFs("/fs/app"))
		r.Any("/u/*", fs.NewUserFs("/fs/u"))
	}
	tsPorxy := httputil.NewSingleHostReverseProxy(logv.AssertFuncErr(url.Parse("http://v.v:8428")))
	fsProxy := fs.NewFs("/home/v/cache/", "")

	app.SetMux(func(w http.ResponseWriter, r *http.Request) func(http.ResponseWriter, *http.Request) {
		if r.Host == "ts.oa.v" || r.Header.Get("mux") == "ts" {
			return tsPorxy.ServeHTTP
		} else if r.Host == "fs.oa.v" {
			return fsProxy.ServeHTTP
		}
		return nil
	})
}
