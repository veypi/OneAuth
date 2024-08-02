//
// response.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-08-01 16:42
// Distributed under terms of the MIT license.
//

package errs

import (
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Response(w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		code := http.StatusInternalServerError
		msg := err.Error()
		switch e := err.(type) {
		case *CodeErr:
			code = e.Code
			msg = e.Msg
		case *mysql.MySQLError:
			logx.Info(e.Error())
			code = http.StatusBadRequest
			msg = e.Message
		}
		w.WriteHeader(code)
		w.Write([]byte(msg))
	} else if resp != nil {
		httpx.OkJson(w, resp)
	} else {
		httpx.Ok(w)
	}
}
