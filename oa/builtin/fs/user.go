//
// user.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-22 15:49
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

func NewUserFs(prefix string) func(http.ResponseWriter, *http.Request) {
	tmp := utils.PathJoin(cfg.Config.FsPath, "u")
	if !utils.FileExists(tmp) {
		logv.AssertError(os.MkdirAll(tmp, 0744))
	}

	client := webdav.NewWebdav(tmp)
	client.Prefix = prefix
	client.GenSubPathFunc = func(r *http.Request) string {
		return ""
	}
	return client.ServeHTTP
}
