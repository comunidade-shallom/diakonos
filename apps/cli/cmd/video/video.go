package video

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:        "video",
	Usage:       "Interact with video files",
	Subcommands: []*cli.Command{CmdCut, CmdExtract, CmdMerge},
}
