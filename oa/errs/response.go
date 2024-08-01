//
// response.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-08-01 16:42
// Distributed under terms of the MIT license.
//

package errs

import (
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Response(w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		httpx.Error(w, err)
	} else if resp != nil {
		httpx.OkJson(w, resp)
	} else {
		httpx.Ok(w)
	}
}

func ErrorHandler(err error) (int, any) {
	switch e := err.(type) {
	case *CodeMsg:
		return e.Code, e.Msg
	case *mysql.MySQLError:
		fmt.Printf("\nerror: %v| %v\n", e.SQLState, e.Number)
		return http.StatusUnprocessableEntity, e.Message
	default:
		fmt.Printf("\nerror: %T| %v\n", err, err)
		return http.StatusInternalServerError, err.Error()
	}
}
