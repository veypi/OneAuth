package main

import (
	"github.com/urfave/cli/v2"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/sub"
	"github.com/veypi/utils/cmd"
	"github.com/veypi/utils/log"
	"os"
)

const Version = "v0.1.0"

func main() {
	cmd.LoadCfg(cfg.Path, cfg.CFG)
	app := cli.NewApp()
	app.Name = "oneauth"
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
		sub.File,
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
		return nil
	}
	_ = app.Run(os.Args)
	srv.Run()

}
