package youtube

import (
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/extract"
	"github.com/comunidade-shallom/diakonos/pkg/fileutils"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var cutFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "extract-audio",
		Usage: "Extract file audio",
	},
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

		videoFile, _, err := download.YouTube(c.Context, params)

		if err != nil {
			if e, ok := err.(errors.BusinessError); ok && e.ErrorCode == download.ErrExist.ErrorCode {
				pterm.Warning.Println(e.Error())
			} else {
				return err
			}
		}

		pterm.Success.Printfln("Done: %s", fileutils.GetRelative(videoFile.Name))

		cfg := config.Ctx(c.Context)

		start := c.Duration("start")
		finish := c.Duration("finish")

		croppedFile, err := cut.CutFile(cut.CutParams{
			OutputDir: cfg.Cut.OutputDir,
			Source:    videoFile.Name,
			Start:     start,
			Finish:    finish,
		})

		if err != nil {
			return err
		}

		pterm.Success.Printfln("Done: %s", fileutils.GetRelative(croppedFile.Name))

		if c.Bool("extract-audio") {
			audioFile, err := extract.Audio(extract.ExtractParams{
				Source:    croppedFile.Name,
				OutputDir: cfg.Audio.OutputDir,
			})

			if err != nil {
				return err
			}

			pterm.Success.Printfln("Done: %s", fileutils.GetRelative(audioFile.Name))
		}

		pterm.Success.Printfln("Done")

		return err
	},
}
