package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/veypi/OneAuth/api/app"
	"github.com/veypi/OneAuth/api/resource"
	"github.com/veypi/OneAuth/api/role"
	"github.com/veypi/OneAuth/api/token"
	"github.com/veypi/OneAuth/api/user"
	"github.com/veypi/OneAuth/api/wx"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/libs/base"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/core"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils"
	"github.com/veypi/utils/log"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func Router(r OneBD.Router) {
	r.SetNotFoundFunc(func(m core.Meta) {
		log.Info().Msgf("%s", m.Request().RequestURI)
		base.JSONResponse(m, nil, nil)
	})
	r.SetInternalErrorFunc(func(m core.Meta) {
		base.JSONResponse(m, nil, nil)
	})
	user.Router(r.SubRouter("/user"))
	wx.Router(r.SubRouter("wx"))

	app.Router(r.SubRouter("app"))
	appRouter := r.SubRouter("/app/:uuid")
	role.Router(appRouter.SubRouter("role"))
	resource.Router(appRouter.SubRouter("resource"))
	token.Router(appRouter.SubRouter("token"))

	r.Set("upload", handleUpload, rfc.MethodPost)

	//message.Router(r.SubRouter("/message"))
}

func handleUpload(m OneBD.Meta) {
	r := m.Request()
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		base.JSONResponse(m, nil, err)
		return
	}
	bufLen := 32 << 10
	buf := make([]byte, bufLen)
	// 缓存上一级读取数据, 使用copy 复制数据 最多缓存2倍buf, 为倒数第1,2桢
	bufCache := make([]byte, 2*bufLen)
	bufCacheLen := 0
	headerReg := regexp.MustCompile(`------WebKitFormBoundary.*\r\n.*?filename="([^"]+)".*\r\n(.*)\r\n\r\n(?s:(.*))`)
	footReg := regexp.MustCompile(`(?s:(.*))\r\n------WebKitFormBoundary.*?--\r\n`)
	fileName := ""
	fileDir, err := filepath.Abs(cfg.CFG.MediaDir)
	if err != nil {
		base.JSONResponse(m, nil, err)
		return
	}
	hash := md5.New()
	var written int64 = 0
	firstRead := true
	nr := 0
	var er error
	var out *os.File
	var filePath string
	for {
		nr, er = r.Body.Read(buf)
		if er != nil && er != io.EOF {
			base.JSONResponse(m, nil, err)
			return
		}
		if firstRead {
			firstRead = false
			res := headerReg.FindAllSubmatch(buf[0:nr], -1)
			if len(res) == 0 || len(res[0]) == 0 {
				log.Warn().Msgf("reg failed for header form: %s: %d %d", buf, bufLen, length)
				base.JSONResponse(m, nil, oerr.ApiArgsError)
				return
			}
			fileName = string(res[0][1])
			copy(bufCache, res[0][3])
			bufCacheLen = len(res[0][3])
			nr = 0
			filePath = fmt.Sprintf("upload/%s_%s_%s", time.Now().Format("2006-01-02-15-04-05"), utils.RandSeq(5), fileName)
			out, err = os.OpenFile(path.Join(fileDir, filePath),
				os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				base.JSONResponse(m, nil, err)
				return
			}
			defer out.Close()
		}
		if er == io.EOF {
			res := footReg.FindAllSubmatch(append(bufCache[0:bufCacheLen], buf[0:nr]...), -1)
			if len(res) == 0 || len(res[0]) == 0 {
				log.Warn().Msgf("reg failed for foot form: %s", buf[0:nr])
				base.JSONResponse(m, nil, oerr.ApiArgsError)
				return
			}
			copy(bufCache, res[0][1])
			bufCacheLen = len(res[0][1])
		}
		if bufCacheLen > 0 {
			if nw, ew := out.Write(bufCache[0:bufCacheLen]); ew != nil {
				base.JSONResponse(m, nil, ew)
				return
			} else if nw > 0 {
				written += int64(nw)
			}
			if _, ew := hash.Write(bufCache[0:bufCacheLen]); ew != nil {
				base.JSONResponse(m, nil, ew)
				return
			}
		}
		copy(bufCache, buf[0:nr])
		bufCacheLen = nr
		if er == io.EOF {
			break
		}
	}
	_ = hex.EncodeToString(hash.Sum(nil))
	base.JSONResponse(m, "/media/"+filePath, nil)
	return
}
