package audio

import (
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var ErrorMissingSourceArgument = errors.Business("Missing 'source' param", "DV:001")

var CmdNormalize = &cli.Command{
	Name:  "normalize",
	Usage: "normalize audio",
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

		file, err := audios.Normalize(ctx.Context, params)

		if err == nil {
			pterm.Success.Printfln("Done: %s", file.NameRelative())
		}

		return err
	},
}
