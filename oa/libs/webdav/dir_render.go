//
// dir_render.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-10-18 19:35
// Distributed under terms of the GPL license.
//

package webdav

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"net/url"
	"strings"

	"github.com/veypi/utils/logv"
)

type anyDirs interface {
	len() int
	name(i int) string
	isDir(i int) bool
}

type fileInfoDirs []fs.FileInfo

func (d fileInfoDirs) len() int          { return len(d) }
func (d fileInfoDirs) isDir(i int) bool  { return d[i].IsDir() }
func (d fileInfoDirs) name(i int) string { return d[i].Name() }

type dirEntryDirs []fs.DirEntry

func (d dirEntryDirs) len() int          { return len(d) }
func (d dirEntryDirs) isDir(i int) bool  { return d[i].IsDir() }
func (d dirEntryDirs) name(i int) string { return d[i].Name() }

//go:embed dir.html
var dirTemplate embed.FS // 嵌入文件系统

func size2Label(s int64) string {
	if s < 1024 {
		return fmt.Sprintf("%d B", s)
	} else if s < 1048576 {
		return fmt.Sprintf("%d KB", s/1024)
	} else if s < 1073741824 {
		return fmt.Sprintf("%d MB", s/1024/1024)
	} else {
		return fmt.Sprintf("%d MB", s/1024/1024/1024)
	}
}

func dirList(w http.ResponseWriter, r *http.Request, f File, rootPath string) {

	// Prefer to use ReadDir instead of Readdir,
	// because the former doesn't require calling
	// Stat on every entry of a directory on Unix.
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
		if err != nil {
			logv.Warn().Msgf("http: error reading directory: %v", err)
			http.Error(w, "Error reading directory", http.StatusInternalServerError)
		}
	}()
	dir_count := 0
	f_count := 0
	var file_bytes int64
	files := make([][4]any, 0, 5)
	dirs := make([][4]any, 0, 5)
	if d, ok := f.(fs.ReadDirFile); ok {
		for _, item := range logv.AssertFuncErr(d.ReadDir(-1)) {
			if item.IsDir() {
				name := item.Name() + "/"
				dir_count += 1
				fstat, _ := item.Info()
				furl := url.URL{Path: name}
				dirs = append(dirs, [4]any{name, furl.String(), "-----", fstat.ModTime().UTC()})
			} else {
				name := item.Name()
				f_count += 1
				fstat, _ := item.Info()
				file_bytes += fstat.Size()
				furl := url.URL{Path: name}
				files = append(files, [4]any{name, furl.String(), size2Label(fstat.Size()), fstat.ModTime().UTC()})
			}
		}
	} else {
		for _, item := range logv.AssertFuncErr(d.ReadDir(-1)) {
			name := item.Name()
			f_count += 1
			fstat, _ := item.Info()
			file_bytes += fstat.Size()
			furl := url.URL{Path: name}
			files = append(files, [4]any{name, furl.String(), size2Label(fstat.Size()), fstat.ModTime().UTC()})
		}
	}

	dirBody := logv.AssertFuncErr(dirTemplate.ReadFile("dir.html"))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl := logv.AssertFuncErr(template.New("").Parse(string(dirBody)))
	logv.AssertError(tpl.Execute(w, map[string]any{"files": files, "dirs": dirs, "path": strings.Split(rootPath, "/"), "cdir": dir_count, "cfile": f_count, "size": size2Label(file_bytes)}))
}
