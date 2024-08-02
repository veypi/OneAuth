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
)

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
	AuthFailed     = New(401, "auth failed")
	AuthExpired    = New(401, "auth expired")
	AuthInvalid    = New(401, "auth invalid")
	ArgsInvalid    = New(http.StatusBadRequest, "args invalid")
	UserNotFound   = New(400, "user not found")
	UserPwdInvalid = New(400, "password invalid")
)
