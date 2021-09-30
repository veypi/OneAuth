package sub

import (
	"OneAuth/api"
	"OneAuth/cfg"
	"OneAuth/models"
	"embed"
	"github.com/urfave/cli/v2"
	"github.com/veypi/OneBD"
	"github.com/veypi/OneBD/core"
	"github.com/veypi/OneBD/rfc"
	"github.com/veypi/utils/log"
	"net/http"
	"os"
)

//go:embed static
var staticFiles embed.FS

//go:embed static/index.html
var indexFile []byte

var Web = cli.Command{
	Name:        "web",
	Usage:       "",
	Description: "oa 核心http服务",
	Action:      RunWeb,
	Flags:       []cli.Flag{},
}

func RunWeb(c *cli.Context) error {
	_ = runSyncDB(c)
	ll := log.InfoLevel
	if l, err := log.ParseLevel(cfg.CFG.LoggerLevel); err == nil {
		ll = l
	}
	app := OneBD.New(&OneBD.Config{
		Host:        cfg.CFG.Host,
		LoggerPath:  cfg.CFG.LoggerPath,
		LoggerLevel: ll,
	})
	app.Router().EmbedFile("/", indexFile)
	app.Router().EmbedDir("/", staticFiles, "static/")

	// TODO media 文件需要检验权限
	//app.Router().SubRouter("/media/").Static("/", cfg.CFG.EXEDir+"/media")

	app.Router().SetNotFoundFunc(func(m core.Meta) {
		f, err := os.Open(cfg.CFG.EXEDir + "/static/index.html")
		if err != nil {
			m.WriteHeader(rfc.StatusNotFound)
			return
		}
		defer f.Close()
		info, err := f.Stat()
		if err != nil {
			m.WriteHeader(rfc.StatusNotFound)
			return
		}
		if info.IsDir() {
			// TODO:: dir list
			m.WriteHeader(rfc.StatusNotFound)
			return
		}
		http.ServeContent(m, m.Request(), info.Name(), info.ModTime(), f)
	})

	api.Router(app.Router().SubRouter("api"))

	log.Info().Msg("\nRouting Table\n" + app.Router().String())
	return app.Run()
}

func runSyncDB(*cli.Context) error {
	db := cfg.DB()
	log.HandlerErrs(
		db.SetupJoinTable(&models.User{}, "Roles", &models.UserRole{}),
		db.SetupJoinTable(&models.Role{}, "Users", &models.UserRole{}),
		db.SetupJoinTable(&models.Role{}, "Auths", &models.RoleAuth{}),
		db.AutoMigrate(&models.User{}, &models.Role{}, &models.Auth{}),
	)
	log.HandlerErrs(
		db.AutoMigrate(&models.App{}, &models.Wechat{}),
	)
	return nil
}
