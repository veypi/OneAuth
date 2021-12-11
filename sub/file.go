package sub

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/urfave/cli/v2"
	"github.com/veypi/OneAuth/cfg"
	"github.com/veypi/OneAuth/models"
	"github.com/veypi/OneAuth/models/index"
	"github.com/veypi/utils/log"
	"reflect"
)

var File = &cli.Command{
	Name: "file",
	Subcommands: []*cli.Command{
		{
			Name: "list",
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name: "id",
				},
			},
			Action: runFileList,
		},
		{
			Name: "history",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "id",
				},
				&cli.UintFlag{Name: "from"},
				&cli.UintFlag{Name: "size"},
				&cli.UintFlag{Name: "uid", Usage: "user_id"},
			},
			Action: runFileHistory,
		},
	},
}

func runFileList(c *cli.Context) error {
	id := c.Uint("id")
	log.Info().Msgf("try to get file of %d", id)
	res, err := index.File.Search().
		Query(elastic.NewTermQuery("OwnerID", id)).
		Do(cfg.Ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%-32s %-40s  %-6s %6s\n", "file_id", "path", "count", "size")
	for _, c := range res.Each(reflect.TypeOf(&models.File{})) {
		f := c.(*models.File)
		fmt.Printf("%-32s %-40s  %-6d %6d\n", f.ID(), f.Path, f.Count, f.Size)
	}
	return nil
}
func runFileHistory(c *cli.Context) error {
	id := c.String("id")
	uid := c.Uint("uid")
	from := int(c.Uint("from"))
	size := int(c.Uint("size"))
	if size <= from {
		size = from + 10
	}
	log.Info().Msgf("try to get history[%d, %d] of %s", from, size, id)
	query := index.History.Search().Sort("CreatedAt", false).From(from).Size(size)
	if id != "" {
		query = query.Query(elastic.NewTermQuery("FileID", id))
	}
	if uid > 0 {
		query = query.Query(elastic.NewTermQuery("UserID", uid))
	}
	res, err := query.Do(cfg.Ctx)
	if err != nil {
		return err
	}
	fmt.Printf("      %-24s %-10s  %-20s %5s %s\n", "time", "action", "ip", "user", "note")
	for index, c := range res.Each(reflect.TypeOf(&models.History{})) {
		f := c.(*models.History)
		fmt.Printf("%-5d %-24s %-10s  %-20s %5s %s\n", index, f.CreatedAt.Format(cfg.CFG.TimeFormat), f.Action, f.IP, f.ActorID, f.Tag)
	}
	return nil
}
