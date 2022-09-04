package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/comunidade-shallom/diakonos/apps/cli/cmd"
	"github.com/comunidade-shallom/diakonos/apps/cli/cmd/audio"
	"github.com/comunidade-shallom/diakonos/apps/cli/cmd/pipeline"
	"github.com/comunidade-shallom/diakonos/apps/cli/cmd/video"
	"github.com/comunidade-shallom/diakonos/apps/cli/cmd/youtube"
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/support"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Description:          "Diakonos - Tools to speed media content development",
		Usage:                "Diakonos CLI",
		Version:              config.Version(),
		Copyright:            "https://github.com/comunidade-shallom",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Load configuration from",
				DefaultText: fmt.Sprintf("%s/diakonos.yml", support.GetBinDirPath()),
			},
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug mode",
			},
		},
		Commands: []*cli.Command{youtube.Cmd, video.Cmd, pipeline.Cmd, cmd.CmdConfig, cmd.CmdCover, audio.Cmd},
		Before: func(ctx *cli.Context) error {
			pterm.Debug.Debugger = !ctx.Bool("debug")

			pterm.DefaultHeader.
				WithMargin(5). //nolint:gomnd
				Println("Diakonos CLI")

			appConfig, err := config.Load(ctx.String("config"))
			if err != nil {
				e, ok := err.(errors.BusinessError) //nolint:errorlint
				if ok && e.ErrorCode == config.ConfigFileWasCreated.ErrorCode {
					pterm.Warning.Println(err.Error())
				} else {
					pterm.Fatal.PrintOnErrorf("Fail to load config (%s)", err)

					return err
				}
			}

			ctx.Context = appConfig.WithContext(ctx.Context)

			pterm.Debug.Printfln("Config loaded")

			return nil
		},
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println("Diakonos - Tools to speed media content development")
		fmt.Println("")
		fmt.Println(config.VersionVerbose())
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		pterm.Error.Println(err.Error())
		os.Exit(1)
	}
}
