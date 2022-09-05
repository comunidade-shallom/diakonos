package video

import (
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var CmdExtract = &cli.Command{
	Name:        "extract",
	Usage:       "Extract audio from a video file",
	Subcommands: []*cli.Command{CmdExtractAudio},
}

var CmdExtractAudio = &cli.Command{
	Name: "audio",
	Action: func(ctx *cli.Context) error {
		source := ctx.Args().First()

		if source == "" {
			return ErrorMissingSourceArgument
		}

		if !path.IsAbs(source) {
			pwd, _ := os.Getwd()
			source = path.Join(pwd, source)
		}

		cfg := config.Ctx(ctx.Context)

		params, err := cfg.Audio.Apply(audios.Params{
			Source: source,
		})
		if err != nil {
			return err
		}

		file, err := audios.Extract(ctx.Context, params)

		if err == nil {
			pterm.Success.Printfln("Done: %s", file.NameRelative())
		}

		return err
	},
}
