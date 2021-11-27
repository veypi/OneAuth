package oalib

import (
	"errors"
	"fmt"
	"github.com/veypi/utils/jwt"
	"io/ioutil"
	"net/http"
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
	return &OA{cfg: c}
}

type OA struct {
	cfg *Config
}

func (oa *OA) Ping() error {
	url := fmt.Sprintf("%s/api/app/%s/ping", oa.cfg.Host, oa.cfg.UUID)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	p := &PayLoad{}
	t, err := jwt.GetToken(p, oa.cfg.Key)
	if err != nil {
		return err
	}
	req.Header.Set("auth_token", t)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if string(body) != "ok" {
		return errors.New(string(body))
	}
	return nil
}

func (oa *OA) LoginUrl() string {
	return fmt.Sprintf("%s/login?uuid=%s", oa.cfg.Host, oa.cfg.UUID)
}

func (oa *OA) Parse(token string, payload jwt.PayloadInterface) (bool, error) {
	return jwt.ParseToken(token, payload, oa.cfg.Key)
}
