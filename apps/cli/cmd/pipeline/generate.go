package pipeline

import (
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/pipeline"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/urfave/cli/v2"
)

var ErrorMissingSourceArgument = errors.Business("Missing 'source' param", "DP:001")

var Cmd = &cli.Command{
	Name:  "pipeline",
	Usage: "Process a .yml file and dynamic generate multimedia assets",
	Action: func(ctx *cli.Context) error {
		source := ctx.Args().First()

		if source == "" {
			return ErrorMissingSourceArgument
		}

		if !path.IsAbs(source) {
			pwd, _ := os.Getwd()
			source = path.Join(pwd, source)
		}

		content, err := os.ReadFile(source)
		if err != nil {
			return err
		}

		collection, err := pipeline.Parse(content)
		if err != nil {
			return err
		}

		cfg := config.Ctx(ctx.Context)

		_, err = collection.Run(ctx.Context, *cfg)

		return err
	},
}
