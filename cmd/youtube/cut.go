package youtube

import (
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var cutFlags = []cli.Flag{
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
	&cli.StringFlag{
		Name: "output_dir",
	},
	&cli.StringFlag{
		Name: "quality",
	},
	&cli.StringFlag{
		Name: "mime_type",
	},
}

var CmdCut = &cli.Command{
	Name:  "cut",
	Usage: "Crop YouTube video",
	Flags: cutFlags,
	Action: func(c *cli.Context) error {

		params, err := getDownloadParams(c)

		if err != nil {
			return err
		}

		file, _, err := download.YouTube(c.Context, params)

		if err != nil {
			return err
		}

		cfg := config.Ctx(c.Context)

		start := c.Duration("start")
		finish := c.Duration("finish")

		_, err = cut.CutFile(cut.CutParams{
			OutputDir: cfg.Cut.OutputDir,
			Source:    file.Name,
			Start:     start,
			Finish:    finish,
		})

		if err != nil {
			return err
		}

		pterm.Success.Printfln("Done")

		return err
	},
}