package video

import (
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/urfave/cli/v2"
)

var ErrorMissingSourceArgument = errors.Business("Missing 'source' param", "DV:001")

var CmdCut = &cli.Command{
	Name:  "cut",
	Usage: "cut video file",
	Flags: []cli.Flag{
		&cli.DurationFlag{
			Name:     "start",
			Usage:    "begin of video",
			Required: true,
		},
		&cli.DurationFlag{
			Name:     "finish",
			Usage:    "end of video",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		source := ctx.Args().First()

		if source == "" {
			return ErrorMissingSourceArgument
		}

		cfg := config.Ctx(ctx.Context)

		start := ctx.Duration("start")
		finish := ctx.Duration("finish")

		params, err := cfg.Cut.Apply(cut.Params{
			Source: source,
			Start:  start,
			Finish: finish,
		})
		if err != nil {
			return err
		}

		_, err = cut.CutFile(params)

		return err
	},
}
