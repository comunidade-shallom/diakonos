package cmd

import (
	"os"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var ErrFailToMarshalConfig = errors.System(nil, "Fail to marshal config", "DCMD:001")

var CmdConfig = &cli.Command{
	Name:  "config",
	Usage: "Output the actual config",
	Action: func(ctx *cli.Context) error {
		cfg := config.Ctx(ctx.Context)

		data, err := yaml.Marshal(&cfg)
		if err != nil {
			return ErrFailToMarshalConfig.WithErr(err)
		}

		os.Stdout.WriteString("\n")
		os.Stdout.Write(data)
		os.Stdout.WriteString("\n")

		return nil
	},
}
