package audio

import (
	"os"
	"path"
	"strconv"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/pterm/pterm"
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
		if err != nil {
			return err
		}

		err = pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
			{"Key", "Value"},
			{"Title", info.Title},
			{"Artist", info.Artist},
			{"Album", info.Album},
			{"Year", strconv.Itoa(info.Year)},
			{"URL", info.URL},
		}).Render()

		if err != nil {
			return err
		}

		if info.Comment != "" {
			pterm.Print(info.Comment)
		}

		return nil
	},
}
