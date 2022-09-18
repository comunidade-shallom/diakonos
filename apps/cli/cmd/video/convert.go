package video

import (
	"fmt"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/convert"
	"github.com/urfave/cli/v2"
)

var CmdConvert = &cli.Command{
	Name:  "convert",
	Usage: "convert video file to MP4",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "preset",
			Usage:    fmt.Sprintf("quality preset (%s)", convert.PresetsAvailable),
			Value:    string(convert.Fast),
			Required: false,
		},
	},
	Action: func(cmd *cli.Context) error {
		source := cmd.Args().First()

		if source == "" {
			return ErrorMissingSourceArgument
		}

		cfg := config.Ctx(cmd.Context)

		params, err := cfg.Convert.Apply(convert.Params{
			Source: source,
			Preset: convert.Preset(cmd.String("preset")),
		})
		if err != nil {
			return err
		}

		_, err = convert.ToMP4(cmd.Context, params)

		return err
	},
}
