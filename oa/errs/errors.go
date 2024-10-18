//
// errors.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-08-01 15:40
// Distributed under terms of the MIT license.
//

package errs

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/veypi/OneBD/rest"
	"github.com/veypi/utils/logv"
	"gorm.io/gorm"
)

func JsonResponse(x *rest.X, data any) error {
	x.WriteHeader(http.StatusOK)
	return x.JSON(map[string]any{"code": 0, "data": data})
}

func JsonErrorResponse(x *rest.X, err error) {
	code, msg := errIter(err)
	x.WriteHeader(code / 100)
	x.JSON(map[string]any{"code": code, "err": msg})
}

func errIter(err error) (code int, msg string) {
	code = 50000
	msg = err.Error()
	switch e := err.(type) {
	case *CodeErr:
		code = e.Code
		msg = e.Msg
	case *mysql.MySQLError:
		if e.Number == 1062 {
			code = DuplicateKey.Code
			msg = DuplicateKey.Msg
		} else {
			logv.Warn().Msgf("unhandled db error %d: %s", e.Number, err)
			msg = "db error"
		}
	case interface{ Unwrap() error }:
		code, _ = errIter(e.Unwrap())
	default:
		if errors.Is(e, gorm.ErrRecordNotFound) {
			code = ResourceNotFound.Code
			msg = ResourceNotFound.Msg
		} else if errors.Is(e, rest.ErrParse) {
			code = ArgsInvalid.Code
			msg = e.Error()
		} else {
			logv.Warn().Msgf("unhandled error type: %T\n%s", err, err)
			msg = e.Error()
		}
	}
	return
}

type CodeErr struct {
	Code int
	Msg  string
}

func (c *CodeErr) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", c.Code, c.Msg)
}

func (c *CodeErr) WithErr(e error) error {
	nerr := &CodeErr{
		Code: c.Code,
		Msg:  fmt.Errorf("%s: %w", c.Msg, e).Error(),
	}
	return nerr
}

func (c *CodeErr) WithStr(m string) error {
	nerr := &CodeErr{
		Code: c.Code,
		Msg:  fmt.Errorf("%s: %s", c.Msg, m).Error(),
	}
	return nerr
}

// New creates a new CodeMsg.
func New(code int, msg string) *CodeErr {
	return &CodeErr{Code: code, Msg: msg}
}

var (
	ArgsInvalid      = New(40001, "args invalid")
	DuplicateKey     = New(40002, "duplicate key")
	AuthNotFound     = New(40100, "auth not found")
	AuthFailed       = New(40101, "auth failed")
	AuthExpired      = New(40102, "auth expired")
	AuthInvalid      = New(40103, "auth invalid")
	AuthNoPerm       = New(40104, "no permission")
	NotFound         = New(40400, "not found")
	ResourceNotFound = New(40401, "resource not found")
	DBError          = New(50010, "db error")
)
