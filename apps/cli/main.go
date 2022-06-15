package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Description:          "Diakonos - Tools to speed media content development",
		Usage:                "Diakonos CLI",
		Copyright:            "https://github.com/comunidade-shallom",
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
		panic(err)
	}
}
