package audio

import (
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:        "audio",
	Usage:       "Interact with audio files",
	Subcommands: []*cli.Command{CmdNormalize, CmdInfo},
}
