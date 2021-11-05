package main

import (
	"OneAuth/cfg"
	"OneAuth/sub"
	"github.com/urfave/cli/v2"
	"github.com/veypi/utils/cmd"
	"github.com/veypi/utils/log"
	"os"
)

const Version = "v0.1.0"

func main() {
	cmd.LoadCfg(cfg.Path, cfg.CFG)
	app := cli.NewApp()
	app.Name = "OneAuth"
	app.Usage = "one auth"
	app.Version = Version
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			Value:       cfg.CFG.Debug,
			Destination: &cfg.CFG.Debug,
		},
		&cli.StringFlag{
			Name:        "log_level,log",
			Value:       cfg.CFG.LoggerLevel,
			Destination: &cfg.CFG.LoggerLevel,
		},
		&cli.StringFlag{
			Name:        "log_path",
			Value:       cfg.CFG.LoggerPath,
			Destination: &cfg.CFG.LoggerPath,
		},
		&cli.StringFlag{
			Name:        "host",
			Value:       cfg.CFG.Host,
			Destination: &cfg.CFG.Host,
		},
	}
	app.Commands = []*cli.Command{
		sub.Web,
		sub.App,
		sub.Role,
		sub.Resource,
		sub.Init,
	}
	srv, err := cmd.NewSrv(app, sub.RunWeb, cfg.CFG, cfg.Path)
	if err != nil {
		log.Warn().Msg(err.Error())
		return
	}
	srv.SetExecMax(1)
	srv.SetStopFunc(func() {
	})
	app.Before = func(c *cli.Context) error {
		if cfg.CFG.Debug {
			cfg.CFG.LoggerLevel = "debug"
		}
		cfg.ConnectDB()
		return nil
	}
	_ = app.Run(os.Args)

}
