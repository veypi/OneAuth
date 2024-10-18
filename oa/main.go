//
// main.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-09-20 16:10:16
// Distributed under terms of the MIT license.
//

package main

import (
	"net/http"
	"oa/api"
	"oa/builtin"
	"oa/cfg"
	"oa/errs"
	_ "oa/models"

	"github.com/veypi/OneBD/rest"
	"github.com/veypi/utils/logv"
)

func main() {
	cfg.CMD.Command = runWeb
	cfg.CMD.Parse()
	err := cfg.CMD.Run()
	if err != nil {
		logv.Warn().Msg(err.Error())
	}
}

func runWeb() error {
	go cfg.RunNats()
	app, err := rest.New(&cfg.Config.RestConf)
	if err != nil {
		return err
	}
	builtin.Enable(app)
	app.Router().Any("/media/*", http.StripPrefix("/media", http.FileServer(http.Dir("/home/v/cache/"))).ServeHTTP)
	apiRouter := app.Router().SubRouter("api")
	api.Use(apiRouter)
	apiRouter.Use(errs.JsonResponse)
	apiRouter.SetErrFunc(errs.JsonErrorResponse)
	app.Router().Print()
	return app.Run()
}
