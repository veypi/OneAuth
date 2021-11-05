package sub

import (
	"OneAuth/api"
	"OneAuth/cfg"
	"embed"
	"github.com/urfave/cli/v2"
	"github.com/veypi/OneBD"
	"github.com/veypi/utils/log"
)

//go:embed static/static
var staticFiles embed.FS

//go:embed static/favicon.ico
var icon []byte

//go:embed static/index.html
var indexFile []byte

var Web = &cli.Command{
	Name:        "web",
	Usage:       "",
	Description: "oa 核心http服务",
	Action:      RunWeb,
	Flags:       []cli.Flag{},
}

func RunWeb(c *cli.Context) error {
	ll := log.InfoLevel
	if l, err := log.ParseLevel(cfg.CFG.LoggerLevel); err == nil {
		ll = l
	}
	app := OneBD.New(&OneBD.Config{
		Host:        cfg.CFG.Host,
		LoggerPath:  cfg.CFG.LoggerPath,
		LoggerLevel: ll,
	})

	api.Router(app.Router().SubRouter("api"))

	// TODO media 文件需要检验权限
	app.Router().SubRouter("/media/").Static("/", cfg.CFG.MediaDir)
	app.Router().EmbedDir("/static", staticFiles, "static/static/")
	app.Router().EmbedFile("/favicon.ico", icon)
	app.Router().EmbedFile("/*", indexFile)

	log.Info().Msg("\nRouting Table\n" + app.Router().String())
	return app.Run()
}
