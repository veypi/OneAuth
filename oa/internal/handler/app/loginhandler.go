package app

import (
	"fmt"
	"net/http"

	"oa/errs"

	"oa/internal/logic/app"
	"oa/internal/svc"
	"oa/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppReq
		if err := httpx.Parse(r, &req); err != nil {
			errs.Response(w, nil, err)
			return
		}
		fmt.Printf("\n|%v|\n", req)

		l := app.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		errs.Response(w, resp, err)
	}
}
