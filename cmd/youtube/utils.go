package youtube

import (
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/urfave/cli/v2"
)

var (
	ErrorMissingFromArgument = errors.Business("Missing 'from' param (eg.: https://www.youtube.com/watch?v=8yAbX8W3Caw)", "DCD:001")
)

func getDownloadParams(c *cli.Context) (download.DownloadParams, error) {
	conf := config.Ctx(c.Context)

	from := c.Args().First()

	if from == "" {
		return download.DownloadParams{}, ErrorMissingFromArgument
	}

	params := download.DownloadParams{
		From:      from,
		OutputDir: c.String("output_dir"),
		Quality:   c.String("quality"),
		MimeType:  c.String("mime_type"),
	}

	if params.OutputDir == "" {
		params.OutputDir = conf.Download.OutputDir
	}

	if params.Quality == "" {
		params.Quality = conf.Download.Quality
	}

	if params.MimeType == "" {
		params.MimeType = conf.Download.MimeType
	}

	return params, nil
}
