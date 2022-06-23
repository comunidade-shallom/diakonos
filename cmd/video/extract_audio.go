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
	Subcommands: []*cli.Command{CmdExtractAudio},
}

var CmdExtractAudio = &cli.Command{
	Name: "audio",
	Action: func(c *cli.Context) error {
		source := c.Args().First()

		if source == "" {
			return ErrorMissingSourceArgument
		}

		if !path.IsAbs(source) {
			pwd, _ := os.Getwd()
			source = path.Join(pwd, source)
		}

		cfg := config.Ctx(c.Context)

		params, err := cfg.Audio.Apply(audios.Params{
			Source: source,
		})
		if err != nil {
			return err
		}

		file, err := audios.Extract(params)

		if err == nil {
			pterm.Success.Printfln("Done: %s", file.NameRelative())
		}

		return err
	},
}
