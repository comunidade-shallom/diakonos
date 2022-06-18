package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/comunidade-shallom/diakonos/cmd"
	"github.com/comunidade-shallom/diakonos/cmd/video"
	"github.com/comunidade-shallom/diakonos/cmd/youtube"
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
				Usage:       "load configuration from",
				DefaultText: fmt.Sprintf("%s/diakonos.yml", support.GetBinDirPath()),
			},
		},
		Commands: []*cli.Command{youtube.Cmd, video.Cmd, cmd.CmdConfig},
		Before: func(c *cli.Context) error {
			pterm.Debug.Debugger = false

			pterm.DefaultHeader.
				WithMargin(5).
				Println("Diakonos CLI")

			appConfig, err := config.Load(c.String("config"))

			if err != nil {
				e, ok := err.(errors.BusinessError)
				if ok && e.ErrorCode == config.ConfigFileWasCreated.ErrorCode {
					pterm.Warning.Println(err.Error())
				} else {
					pterm.Fatal.PrintOnErrorf("Fail to load config (%s)", err)
					return err
				}
			}

			config.SetConfig(appConfig)

			c.Context = appConfig.WithContext(c.Context)

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

	err := app.Run(os.Args)

	if err != nil {
		pterm.Error.Println(err.Error())
		os.Exit(1)
	}
}
