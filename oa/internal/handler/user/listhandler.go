package user

import (
	"net/http"

	"oa/errs"

	"github.com/zeromicro/go-zero/rest/httpx"
	"oa/internal/logic/user"
	"oa/internal/svc"
	"oa/internal/types"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			errs.Response(w, nil, err)
			return
		}

		l := user.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		errs.Response(w, resp, err)
	}
}
