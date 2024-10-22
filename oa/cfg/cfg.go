//
// cfg.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-09-20 16:10:16
// Distributed under terms of the MIT license.
//

package cfg

import (
	"github.com/veypi/OneBD/rest"
	"github.com/veypi/utils"
	"github.com/veypi/utils/flags"
	"github.com/veypi/utils/logv"
)

type config struct {
	rest.RestConf
	DSN       string `json:"dsn"`
	ID        string `json:"id"`
	Key       string `json:"key"`
	FsPath    string `json:"fs"`
	AccessUrl string `json:"access_url"`
}

var Config = &config{}

var CMD = flags.New("oa", "the backend server of oa")
var CfgDump = CMD.SubCommand("cfg", "generate cfg file")

var configFile = CMD.String("f", "./dev.yaml", "the config file")

func init() {
	CMD.StringVar(&Config.Host, "h", "0.0.0.0", "host")
	CMD.IntVar(&Config.Port, "p", 4000, "port")
	CMD.StringVar(&Config.LoggerLevel, "l", "info", "log level")
	CMD.StringVar(&Config.DSN, "dsn", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local", "data source name")
	CMD.Before = func() error {
		flags.LoadCfg(*configFile, Config)
		CMD.Parse()
		if Config.Key == "" {
			Config.Key = utils.RandSeq(32)
		}
		if Config.ID == "" {
			Config.ID = utils.RandSeq(32)
		}
		logv.SetLevel(logv.AssertFuncErr(logv.ParseLevel(Config.LoggerLevel)))
		return nil
	}
	CfgDump.Command = func() error {
		flags.DumpCfg(*configFile, Config)
		return nil
	}
}
