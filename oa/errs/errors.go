//
// errors.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-08-01 15:40
// Distributed under terms of the MIT license.
//

package errs

import (
	"fmt"
	"net/http"

	"github.com/veypi/OneBD/rest"
)

func JsonResponse(x *rest.X, data any) error {
	x.WriteHeader(http.StatusOK)
	return x.JSON(map[string]any{"code": 0, "data": data})
}

func JsonErrorResponse(x *rest.X, err error) {
	code := 50000
	var msg string
	if e := err.(*CodeErr); e != nil {
		code = e.Code
		msg = e.Msg
	} else {
		msg = err.Error()
	}
	x.WriteHeader(code / 100)
	x.JSON(map[string]any{"code": code, "err": msg})
}

type CodeErr struct {
	Code int
	Msg  string
}

func (c *CodeErr) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", c.Code, c.Msg)
}

func (c *CodeErr) WithErr(e error) error {
	c.Msg = fmt.Sprintf("%s: %s", c.Msg, e.Error())
	return c
}

// New creates a new CodeMsg.
func New(code int, msg string) *CodeErr {
	return &CodeErr{Code: code, Msg: msg}
}

var (
	ArgsInvalid  = New(40001, "args invalid")
	AuthNotFound = New(40100, "auth not found")
	AuthFailed   = New(40101, "auth failed")
	AuthExpired  = New(40102, "auth expired")
	AuthInvalid  = New(40103, "auth invalid")
	AuthNoPerm   = New(40104, "no permission")
	NotFound     = New(40400, "not found")
)
