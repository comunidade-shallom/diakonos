package youtube

import (
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var CmdDownload = &cli.Command{
	Name:  "download",
	Usage: "Download a video from YouTube",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "output_dir",
		},
		&cli.StringFlag{
			Name: "quality",
		},
		&cli.StringFlag{
			Name: "mime_type",
		},
	},
	Action: func(ctx *cli.Context) error {
		params, err := getDownloadParams(ctx)
		if err != nil {
			return err
		}

		_, _, err = download.YouTube(ctx.Context, params)

		if err != nil {
			return err
		}

		pterm.Success.Println("Done")

		return err
	},
}
