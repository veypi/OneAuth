package sub

import (
	"github.com/urfave/cli/v2"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/utils"
	"github.com/veypi/utils/log"
)

var App = &cli.Command{
	Name: "app",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Action: runAppList,
		},
		{
			Name:   "create",
			Action: runAppCreate,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Required: true,
				},
			},
		},
	},
}

func runAppList(c *cli.Context) error {
	list := make([]*models.App, 0, 10)
	err := cfg.DB().Find(&list).Error
	if err != nil {
		return err
	}
	for _, a := range list {
		log.Info().Msgf("%d:  %s", a.UUID, a.Name)
	}
	return nil
}

func runAppCreate(c *cli.Context) error {
	app := &models.App{}
	app.Name = c.String("name")
	app.Key = utils.RandSeq(16)
	app.UUID = utils.RandSeq(8)
	err := cfg.DB().Create(app).Error
	if err != nil {
		return err
	}
	log.Info().Msgf("app: %s\nuuid: %s\nkey: %s", app.Name, app.UUID, app.Key)
	return nil
}
