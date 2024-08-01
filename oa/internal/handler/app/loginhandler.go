package app

import (
	"net/http"

	"oa/errs"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oa/internal/logic/app"
	"oa/internal/svc"
	"oa/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppReq
		if err := httpx.Parse(r, &req); err != nil {
			errs.Response(w, nil, err)
			return
		}

		l := app.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		errs.Response(w, resp, err)
	}
}
