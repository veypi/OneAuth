//
// jwt.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-09-23 18:28
// Distributed under terms of the MIT license.
//

package auth

import (
	"context"
	"errors"
	"fmt"
	"oa/cfg"
	"oa/errs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/veypi/OneBD/rest"
)

func GenJwt(claim *Claims) (string, error) {
	if claim.ExpiresAt == nil {
		claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour))
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(cfg.Config.Key))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJwt(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Config.Key), nil
	})
	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errs.AuthExpired
	}

	if err != nil || !token.Valid {
		return nil, errs.AuthInvalid
	}
	return claims, nil
}

func CheckJWT(x *rest.X) (*Claims, error) {
	authHeader := x.Request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errs.AuthNotFound
	}
	// Token is typically in the format "Bearer <token>"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return nil, errs.AuthInvalid
	}

	// Parse the token
	claims, err := ParseJwt(tokenString)
	if err != nil {
		return nil, err
	}

	x.Request = x.Request.WithContext(context.WithValue(x.Request.Context(), "uid", claims.ID))
	return claims, nil
}

func Check(target string, pid string, l AuthLevel) func(x *rest.X) error {
	return func(x *rest.X) error {
		claims, err := CheckJWT(x)
		if err != nil {
			return err
		}
		tid := ""
		if pid != "" {
			tid = x.Params.GetStr(pid)
		}
		if !claims.Access.Check(target, tid, l) {
			return errs.AuthNoPerm
		}
		return nil
	}
}
