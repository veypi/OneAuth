//
// access.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-09-23 19:37
// Distributed under terms of the MIT license.
//

package auth

import "github.com/golang-jwt/jwt/v5"

type AuthLevel uint

const (
	DoNone   = 0
	Do       = 1
	DoRead   = 1
	DoCreate = 2
	DoUpdate = 3
	DoDelete = 4
	DoAll    = 5
)

type Access []*struct {
	Name  string    `json:"name"`
	TID   string    `json:"tid"`
	Level AuthLevel `json:"level"`
}

func (a *Access) Check(target string, tid string, l AuthLevel) bool {
	if l == DoNone {
		return true
	}
	for _, line := range *a {
		if line.Name == target && line.Level > l {
			if line.TID == "" || line.TID == tid {
				return true
			}
		}
	}
	return false
}

type Claims struct {
	UID    string `json:"uid"`
	AID    string `json:"aid"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Access Access `json:"access"`
	jwt.RegisteredClaims
}
