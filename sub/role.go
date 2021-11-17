package sub

import (
	"OneAuth/cfg"
	"OneAuth/models"
	"github.com/urfave/cli/v2"
	"github.com/veypi/utils/log"
)

var Role = &cli.Command{
	Name:        "role",
	Usage:       "",
	Description: "",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Action: runRoleList,
		},
		{
			Name:   "create",
			Action: runRoleCreate,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "uuid",
					Usage:    "app uuid",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Usage:    "role name",
					Required: true,
				},
			},
		},
	},
	Flags: []cli.Flag{},
}

func runRoleList(c *cli.Context) error {
	roles := make([]*models.Role, 0, 10)
	err := cfg.DB().Find(&roles).Error
	if err != nil {
		return err
	}
	for _, r := range roles {
		log.Info().Msgf("%d %s@%d", r.ID, r.Name, r.AppUUID)
	}
	return nil
}

func runRoleCreate(c *cli.Context) error {
	id := c.String("uuid")
	name := c.String("name")
	rl := &models.Role{}
	rl.AppUUID = id
	rl.Name = name
	err := cfg.DB().Where(rl).FirstOrCreate(rl).Error
	return err
}

var Resource = &cli.Command{
	Name:  "resource",
	Usage: "resource manual",
	Subcommands: []*cli.Command{
		{
			Name:   "list",
			Action: runResourceList,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "uuid",
					Usage: "app uuid",
				},
			},
		},
		{
			Name:   "create",
			Action: runResourceCreate,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "uuid",
					Usage:    "app uuid",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Usage:    "role name",
					Required: true,
				},
			},
		},
	},
}

func runResourceList(c *cli.Context) error {
	query := &models.Resource{}
	query.AppUUID = c.String("uuid")
	l := make([]*models.Resource, 0, 10)
	err := cfg.DB().Where(query).Find(&l).Error
	if err != nil {
		return nil
	}
	for _, r := range l {
		log.Info().Msgf("%d:  %s@%d", r.ID, r.Name, r.AppUUID)
	}
	return nil
}

func runResourceCreate(c *cli.Context) error {
	query := &models.Resource{}
	query.AppUUID = c.String("uuid")
	query.Name = c.String("name")
	err := cfg.DB().Where(query).FirstOrCreate(query).Error
	return err
}
