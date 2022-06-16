package youtube

import (
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:        "youtube",
	Usage:       "Intect with YouTube",
	Subcommands: []*cli.Command{CmdDownload},
}
