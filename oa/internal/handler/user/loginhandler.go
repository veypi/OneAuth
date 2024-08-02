package user

import (
	"net/http"

	"oa/errs"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oa/internal/logic/user"
	"oa/internal/svc"
	"oa/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			errs.Response(w, nil, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		_, err := l.Login(&req)
		errs.Response(w, nil, err)
	}
}
