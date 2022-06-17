package video

import (
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/extract"
	"github.com/comunidade-shallom/diakonos/pkg/fileutils"
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

		file, err := extract.Audio(extract.ExtractParams{
			Source:    source,
			OutputDir: cfg.Audio.OutputDir,
		})

		pterm.Success.Printfln("Done: %s", fileutils.GetRelative(file.Name))

		return err
	},
}
