package base

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/veypi/OneAuth/libs/oerr"
	"github.com/veypi/OneAuth/libs/tools"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils/log"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

var json = jsoniter.ConfigFastest

func JSONResponse(m OneBD.Meta, data interface{}, err error) {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = oerr.ResourceNotExist
		}
	}
	if m.Method() == rfc.MethodHead {
		if err != nil {
			m.SetHeader("status", "0")
			m.SetHeader("code", strconv.Itoa(int(oerr.OfType(err.Error()))))
			m.SetHeader("err", err.Error())
		} else {
			m.SetHeader("status", "1")
		}
		return
	}
	res := map[string]interface{}{
		"status": 1,
	}
	if err != nil {
		res["status"] = 0
		res["code"] = oerr.OfType(err.Error())
		res["err"] = err.Error()
	} else {
		res["status"] = 1
		res["content"] = data
	}
	p, err := json.Marshal(res)
	if err != nil {
		log.Warn().Err(err).Msg("encode json data error")
		return
	}
	_, _ = m.Write(p)
}

type ApiHandler struct {
	OneBD.BaseHandler
	UserHandler
}

func (h *ApiHandler) Init(m OneBD.Meta) error {
	return tools.MultiIniter(m, &h.BaseHandler, &h.UserHandler)
}

func (h *ApiHandler) OnResponse(data interface{}) {
	JSONResponse(h.Meta(), data, nil)
}

func (h *ApiHandler) OnError(err error) {
	log.WithNoCaller.Warn().Err(err).Msg(h.Meta().RequestPath())
	JSONResponse(h.Meta(), nil, err)
}

var ioNumLimit = make(map[string]time.Time)
var limitLocker = sync.RWMutex{}

func (h *ApiHandler) SetAccessDelta(d time.Duration) error {
	// 尽量对写操作加频率限制
	now := time.Now()
	limitLocker.Lock()
	label := h.Meta().RemoteAddr() + h.Meta().RequestPath()
	last, ok := ioNumLimit[label]
	defer func() {
		ioNumLimit[label] = now
		limitLocker.Unlock()
	}()
	if !ok {
		return nil
	} else if now.Sub(last) >= d {
		return nil
	}
	return oerr.AccessTooFast
}
