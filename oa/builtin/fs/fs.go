//
// fs.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-22 15:51
// Distributed under terms of the GPL license.
//

package fs

import (
	"net/http"
	"oa/libs/webdav"
)

func NewFs(dir_path, prefix string) *webdav.Handler {
	client := webdav.NewWebdav(dir_path)
	client.Prefix = prefix
	client.GenSubPathFunc = func(r *http.Request) string {
		return ""
	}
	return client
}
