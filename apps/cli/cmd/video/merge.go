package video

import (
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/merge"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var CmdMerge = &cli.Command{
	Name:  "merge",
	Usage: "merge a list of videos in a single video",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "define the name of generated video",
			DefaultText: "hash from input",
			Required:    false,
		},
	},
	Action: func(ctx *cli.Context) error {
		sources := ctx.Args().Slice()

		if len(sources) == 0 {
			return ErrorMissingSourceArgument
		}

		cfg := config.Ctx(ctx.Context)

		params, err := cfg.Merge.Apply(merge.Params{
			Sources: sources,
			Name:    ctx.String("name"),
		})
		if err != nil {
			return err
		}

		out, err := merge.Files(params)

		if err == nil {
			pterm.Success.Printfln("Done: %s", out.NameRelative())
		}

		return err
	},
}
