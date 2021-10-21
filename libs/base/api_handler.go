package base

import (
	"OneAuth/libs/oerr"
	"OneAuth/libs/tools"
	"github.com/json-iterator/go"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils/log"
	"strconv"
	"sync"
	"time"
)

var json = jsoniter.ConfigFastest

type ApiHandler struct {
	OneBD.BaseHandler
	UserHandler
}

func (h *ApiHandler) Init(m OneBD.Meta) error {
	return tools.MultiIniter(m, &h.BaseHandler, &h.UserHandler)
}

func (h *ApiHandler) OnResponse(data interface{}) {
	if h.Meta().Method() == rfc.MethodHead {
		h.Meta().SetHeader("status", "1")
		return
	}
	p, err := json.Marshal(map[string]interface{}{"status": 1, "content": data})
	if err != nil {
		log.Warn().Err(err).Msg("encode json data error")
		return
	}
	h.Meta().Write(p)
}

func (h *ApiHandler) OnError(err error) {
	log.WithNoCaller.Warn().Err(err).Msg(h.Meta().RequestPath())
	msg := err.Error()
	if h.Meta().Method() == rfc.MethodHead {
		h.Meta().SetHeader("status", "0")
		h.Meta().SetHeader("code", strconv.Itoa(int(oerr.OfType(msg))))
		h.Meta().SetHeader("err", msg)
	} else {
		p, _ := json.Marshal(map[string]interface{}{"status": 0, "code": oerr.OfType(msg), "err": msg})
		h.Meta().Write(p)
	}
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
