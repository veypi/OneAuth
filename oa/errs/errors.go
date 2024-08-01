//
// errors.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-08-01 15:40
// Distributed under terms of the MIT license.
//

package errs

import (
	"fmt"
)

type CodeMsg struct {
	Code int
	Msg  string
}

func (c *CodeMsg) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", c.Code, c.Msg)
}

// New creates a new CodeMsg.
func New(code int, msg string) error {
	return &CodeMsg{Code: code, Msg: msg}
}

var (
	AuthFailed  = New(401, "auth failed")
	AuthExpired = New(401, "auth expired")
	AuthInvalid = New(401, "auth invalid")
)
