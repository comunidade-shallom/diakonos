package video

import (
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/urfave/cli/v2"
)

var (
	ErrorMissingSourceArgument = errors.Business("Missing 'source' param", "DV:001")
)

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
	Action: func(c *cli.Context) error {
		source := c.Args().First()

		if source == "" {
			return ErrorMissingSourceArgument
		}

		cfg := config.Ctx(c.Context)

		start := c.Duration("start")
		finish := c.Duration("finish")

		if !path.IsAbs(source) {
			pwd, _ := os.Getwd()
			source = path.Join(pwd, source)
		}

		_, err := cut.CutFile(cut.CutParams{
			OutputDir: cfg.Cut.OutputDir,
			Source:    source,
			Start:     start,
			Finish:    finish,
		})

		return err
	},
}