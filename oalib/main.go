package oalib

import (
	"fmt"
	"github.com/veypi/utils/jwt"
)

/**
* @name: main
* @author: veypi <i@veypi.com>
* @date: 2021-11-17 16:28
* @description：main
**/

func New(c *Config) *OA {
	if !c.Valid() {
		panic("invalid oa config")
	}
	return &OA{cfg: c, Key: []byte(c.Key)}
}

type OA struct {
	cfg *Config
	Key []byte
}

func (oa *OA) Ping() {
}

func (oa *OA) LoginUrl() string {
	return fmt.Sprintf("%s/login?uuid=%s", oa.cfg.Host, oa.cfg.Key)
}

func (oa *OA) Parse(token string, payload jwt.PayloadInterface) (bool, error) {
	return jwt.ParseToken(token, payload, oa.Key)
}
