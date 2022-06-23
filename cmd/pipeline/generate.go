package pipeline

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/pipeline"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var (
	ErrorMissingSourceArgument = errors.Business("Missing 'source' param", "DP:001")
)

var Cmd = &cli.Command{
	Name:  "pipeline",
	Usage: "Process a .yml file and dynamic generate multimedia assets",
	Action: func(c *cli.Context) error {

		source := c.Args().First()

		if source == "" {
			return ErrorMissingSourceArgument
		}

		if !path.IsAbs(source) {
			pwd, _ := os.Getwd()
			source = path.Join(pwd, source)
		}

		content, err := ioutil.ReadFile(source)

		if err != nil {
			return err
		}

		var collection pipeline.Pipeline

		err = yaml.Unmarshal(content, &collection)

		if err != nil {
			return err
		}

		cfg := config.Ctx(c.Context)

		_, err = collection.Run(c.Context, cfg)

		return err
	},
}
