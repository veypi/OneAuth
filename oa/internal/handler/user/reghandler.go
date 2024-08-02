package user

import (
	"net/http"

	"oa/errs"

	"oa/internal/logic/user"
	"oa/internal/svc"
	"oa/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegReq
		if err := httpx.Parse(r, &req); err != nil {
			errs.Response(w, nil, err)
			return
		}
		l := user.NewRegLogic(r.Context(), svcCtx)
		err := l.Reg(&req)
		errs.Response(w, nil, err)
	}
}
