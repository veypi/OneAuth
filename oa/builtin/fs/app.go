//
// app.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-22 15:42
// Distributed under terms of the GPL license.
//

package fs

import (
	"net/http"
	"oa/cfg"
	"oa/libs/webdav"
	"os"

	"github.com/veypi/utils"
	"github.com/veypi/utils/logv"
)

func NewAppFs(prefix string) func(http.ResponseWriter, *http.Request) {
	apPath := utils.PathJoin(cfg.Config.FsPath, "app")
	if !utils.FileExists(apPath) {
		logv.AssertError(os.MkdirAll(apPath, 0744))
	}

	client := webdav.NewWebdav(apPath)
	client.Prefix = prefix
	client.GenSubPathFunc = func(r *http.Request) string {
		return ""
	}
	return client.ServeHTTP
}
