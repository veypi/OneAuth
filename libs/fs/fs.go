package fs

import (
	_ "embed"
	"fmt"
	"github.com/veypi/OneAuth/libs/webdav"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/utils"
	"github.com/veypi/utils/log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type MountFunc = func(w http.ResponseWriter, r *http.Request) (urlPrefix string, mountPoint string, ownerID string, actorID string, err error)
type FS struct {
	handlerCache  map[string]*webdav.Handler
	RootDir       string
	MountFunc     MountFunc
	DisabledAlert bool
}

//go:embed netErr.png
var errPng []byte

func (f *FS) mount(prefix, dir string) *webdav.Handler {
	p := f.RootDir
	if s, err := filepath.Abs(p); err == nil {
		p = s
	}
	p = filepath.Join(p, dir)
	ok, err := utils.PathExists(p)
	if err != nil || !ok {
		err = os.MkdirAll(p, 0777)
		if err != nil {
			log.Warn().Msgf("create user dir failed: %s", err)
			return nil
		}
	}
	log.Debug().Msgf("mount %s", p)
	fs := webdav.Dir(p)

	h := &webdav.Handler{
		Prefix:     prefix,
		FileSystem: fs,
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				msg := err.Error()
				if !strings.HasSuffix(msg, "no such file or directory") && !strings.HasSuffix(msg, "is a directory") {
					log.Debug().Msgf("\n%s %s %s %s", utils.CallPath(1), r.Method, r.RequestURI, msg)
				}
			}
		},
	}
	return h
}

func (f *FS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prefix, dir, ownerID, actorID, err := f.MountFunc(w, r)
	id := fmt.Sprintf("%s_%s", ownerID, actorID)
	if err != nil {
		if !f.DisabledAlert {
			w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
		}
		FileError(w, http.StatusUnauthorized)
		return
	}
	if f.handlerCache == nil {
		f.handlerCache = map[string]*webdav.Handler{}
	}
	h := f.handlerCache[id]
	if h == nil {
		h = f.mount(prefix, dir)
		h.OwnerID = ownerID
		h.ActorID = actorID
		f.handlerCache[id] = h
		if h == nil {
			if !f.DisabledAlert {
				w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
			}
			FileError(w, http.StatusBadRequest)
			return
		}
		f.setEvent(h)
	}
	h.ServeHTTP(w, r)
}

func (f *FS) setEvent(h *webdav.Handler) {
	// 记录下载历史
	h.On(webdav.EventAfterRead, webdav.ReadFunc(func(r *http.Request, path string) (int, error) {
		defer handlePanic()
		tf, err := getFile(r.Context(), path, h)
		if err != nil {
			return 0, err
		}
		return 0, addHistory(r, h, models.ActGet, tf.ID(), path)
	}))
	h.On(webdav.EventAfterUpdate, webdav.AfterUpdateFunc(func(r *http.Request, path string, size int64, md5Str string) (int, error) {
		defer handlePanic()
		tf, err := getFile(r.Context(), path, h)
		if err != nil {
			return 0, err
		}
		var delta = size - int64(tf.Size)
		err = updateFile(r.Context(), tf.ID(), map[string]interface{}{"Size": size, "MD5": md5Str})
		if err != nil {
			return 0, err
		}
		uid, _ := strconv.Atoi(h.OwnerID)
		err = updateUserSize(uint(uid), delta)
		if err != nil {
			return 0, err
		}
		return 0, addHistory(r, h, models.ActPut, tf.ID(), path)
	}))

	h.On(webdav.EventAfterDelete, webdav.DeleteFunc(func(r *http.Request, path string) (int, error) {
		defer handlePanic()
		return 0, removeFile(r, path, h)
	}))
}

func FileError(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(code)
	w.Write(errPng)
}

func handlePanic() {
	if e := recover(); e != nil {
		log.Warn().Msgf("%v", e)
	}
}
