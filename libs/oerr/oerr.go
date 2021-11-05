package oerr

import (
	"gorm.io/gorm"
	"strconv"
)

// 错误描述

type Code uint

/*

5位10进制码表示错误, 00000 etc.

0 代表未知，或不必定义的有通用意义的错误

## 第1位 错误类型
	- 1 : 系统级错误 比如 内存申请失败, 系统调用失败，文件打开失败等等
	- 2 : 数据库错误
	- 3 : 保留
	- 4 : 权限错误
	- 5 : 配置错误
	- 6 : 参数错误
	- 7 : 时序(控制)错误

## 第2位 2级错误类型

## 第3,4位 具体错误编号

## 第5位 错误严重程度
	- 0 : 无任何影响错误，简单重试可以解决
	- 1 : 无影响错误，重试不可解决
	- 2 : 有影响用户体验或系统性能错误, 重试可解决
    - 3 : 有影响用户体验或系统性能错误, 重试不可解决
	- 4 : 有影响组件功能的错误, 重试可解决
	- 5 : 有影响组件功能的错误， 重试不可解决
	- 6 : 有影响服务运行的错误， 重启可解决
	- 7 : 有影响服务运行的错误，重启不可解决
	- 8 : 有影响系统运行的错误
	- 9 : 本不可能发生的错误，例如被人攻击导致数据异常产生的逻辑错误

*/

// Unknown error
const (
	Unknown Code = 0
)
const (
	// DBErr 2 数据库错误
	//		-1 系统错误
	// 		-2 数据读写错误
	DBErr                 Code = 20001
	ResourceCreatedFailed Code = 22012
	ResourceDuplicated    Code = 22021
	ResourceNotExist      Code = 22031
)

const (
	// LogicErr 3 系统内逻辑错误
	LogicErr   Code = 30000
	AppNotJoin Code = 30001
)

const (
	// NotLogin
	// 4 权限类型错误
	// 1: 登录权限
	// 2: 资源操作权限
	NotLogin        Code = 41001
	LoginExpired    Code = 41011
	PassError       Code = 41021
	DisableLogin    Code = 41031
	AccountNotExist Code = 41041
	NoAuth          Code = 42011
)

// 6 参数类型错误
/*
	-1: 协议参数
	-2: 接口参数
	-3: 函数参数
	-4: 数据依赖错误
*/
const (
	MethodNotSupport Code = 61111
	MethodNotAllowed Code = 61121

	ApiArgsError     Code = 62001
	ApiArgsMissing   Code = 62011
	TableArgsMissing Code = 62021
	TableArgsErr     Code = 62031

	FuncArgsError        Code = 63001
	UrlPatternNotSupport Code = 63117
	UrlDefinedDuplicate  Code = 63127
	UrlParamDuplicate    Code = 63137

	DataError Code = 64009
)

// 7 : 时序(控制)错误
/*
	-1: 访问控制
*/

const (
	AccessErr     Code = 71001
	AccessTooFast Code = 71010
)

var codeMap = map[Code]string{
	Unknown:               "unknown error",
	DBErr:                 "db error",
	ResourceCreatedFailed: "resource created failed",
	ResourceDuplicated:    "resource duplicated",
	ResourceNotExist:      "Resource not exist",
	MethodNotSupport:      "this http method is not supported",
	MethodNotAllowed:      "this http method is not allowed",
	ApiArgsError:          "base args error",
	ApiArgsMissing:        "missing args",
	TableArgsMissing:      "missing data",
	TableArgsErr:          "invalid table data",
	FuncArgsError:         "func args error",
	UrlPatternNotSupport:  "this router's url pattern is not supported.",
	UrlDefinedDuplicate:   "this router's url has been defined",
	UrlParamDuplicate:     "this param defined in router's url duplicated",
	DataError:             "data error",
	NotLogin:              "not login",
	LoginExpired:          "login expired",
	DisableLogin:          "disabled to login",
	PassError:             "password/account error",
	AccountNotExist:       "account not exist",
	NoAuth:                "no auth to access",
	AccessErr:             "access error",
	AccessTooFast:         "access too fast",
	LogicErr:              "logic error",
	AppNotJoin:            "not join in app",
}

func (c Code) Error() string {
	return strconv.Itoa(int(c)) + ":" + c.String()
}

func (c Code) String() string {
	s, ok := codeMap[c]
	if ok && len(s) > 0 {
		return s
	}
	return codeMap[Unknown]
}

// Attach 附加错误详细原因
func (c Code) Attach(errs ...error) (e error) {
	e = c
	for _, err := range errs {
		if err != nil {
			e = &wrapErr{msg: e.Error() + "\n" + err.Error(), err: e}
		}
	}
	return e
}

func (c Code) AttachStr(errs ...string) (e error) {
	e = c
	for _, m := range errs {
		if m != "" {
			e = &wrapErr{
				msg: e.Error() + "\n" + m,
				err: e,
			}
		}
	}
	return e
}

func OfType(errMsg string) Code {
	s := ""
	if gorm.ErrRecordNotFound.Error() == errMsg {
		return ResourceNotExist
	}
	for _, v := range errMsg {
		if v == ':' {
			break
		}
		s += string(v)
	}
	c, _ := strconv.Atoi(s)
	return Code(c)
}

type wrapErr struct {
	msg string
	err error
}

func (w *wrapErr) Error() string {
	return w.msg
}

func (w *wrapErr) UnWrap() error {
	return w.err
}
