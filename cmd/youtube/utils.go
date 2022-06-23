package youtube

import (
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/urfave/cli/v2"
)

var ErrorMissingSourceArgument = errors.Business("Missing 'source' arg (eg.: https://www.youtube.com/watch?v=8yAbX8W3Caw)", "DCD:001")

func getDownloadParams(c *cli.Context) (download.Params, error) {
	conf := config.Ctx(c.Context)

	source := c.Args().First()

	if source == "" {
		return download.Params{}, ErrorMissingSourceArgument
	}

	return conf.Download.Apply(download.Params{
		Source:    source,
		OutputDir: c.String("output_dir"),
		Quality:   c.String("quality"),
		MimeType:  c.String("mime_type"),
	})
}
