package youtube

import (
	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/download"
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
	Action: func(ctx *cli.Context) error {
		params, err := getDownloadParams(ctx)
		if err != nil {
			return err
		}

		videoFile, _, err := download.YouTube(ctx.Context, params)
		if err != nil {
			//nolint:errorlint
			if e, ok := err.(errors.BusinessError); ok && e.ErrorCode == download.ErrExist.ErrorCode {
				pterm.Warning.Println(e.Error())
			} else {
				return err
			}
		}

		pterm.Success.Printfln("Done: %s", videoFile.NameRelative())

		cfg := config.Ctx(ctx.Context)

		start := ctx.Duration("start")
		finish := ctx.Duration("finish")

		cutParams, err := cfg.Cut.Apply(cut.Params{
			Source: videoFile.Filename,
			Start:  start,
			Finish: finish,
		})
		if err != nil {
			return err
		}

		croppedFile, err := cut.Video(ctx.Context, cutParams)
		if err != nil {
			return err
		}

		pterm.Success.Printfln("Done: %s", croppedFile.NameRelative())

		if ctx.Bool("extract-audio") {
			extractParams, err := cfg.Audio.Apply(audios.Params{
				Source: croppedFile.Filename,
			})
			if err != nil {
				return err
			}

			audioFile, err := audios.Extract(ctx.Context, extractParams)
			if err != nil {
				return err
			}

			pterm.Success.Printfln("Done: %s", audioFile.NameRelative())
		}

		if err == nil {
			pterm.Success.Printfln("All done")
		}

		return err
	},
}
