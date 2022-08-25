package audio

import (
	"fmt"
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/urfave/cli/v2"
)

var CmdInfo = &cli.Command{
	Name:  "info",
	Usage: "show audio info",
	Action: func(cmd *cli.Context) error {
		source := cmd.Args().First()

		if source == "" {
			return ErrorMissingSourceArgument
		}

		if !path.IsAbs(source) {
			pwd, _ := os.Getwd()
			source = path.Join(pwd, source)
		}

		info, err := audios.Info(cmd.Context, source)

		fmt.Printf("info: %v\n", info)

		return err
	},
}
