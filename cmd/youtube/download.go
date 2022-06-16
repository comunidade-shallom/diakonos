package youtube

import (
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	ytdl "github.com/kkdai/youtube/v2/downloader"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/http/httpproxy"
)

var (
	ErrorMissingFromArgument = errors.Business("Missing 'from' param (eg.: https://www.youtube.com/watch?v=8yAbX8W3Caw)", "DCD:001")
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
	Action: func(c *cli.Context) error {
		config := config.Ctx(c.Context)

		from := c.Args().First()

		if from == "" {
			return ErrorMissingFromArgument
		}

		params := download.DownloadParams{
			From:      from,
			OutputDir: c.String("output_dir"),
			Quality:   c.String("quality"),
			MimeType:  c.String("mime_type"),
		}

		if params.OutputDir == "" {
			params.OutputDir = config.Download.OutputDir
		}

		if params.Quality == "" {
			params.Quality = config.Download.Quality
		}

		if params.MimeType == "" {
			params.MimeType = config.Download.MimeType
		}

		_, _, err := download.YouTube(c.Context, params)

		pterm.Success.Println("Done")

		return err
	},
}

func getClient(outputDir string) *ytdl.Downloader {

	proxyFunc := httpproxy.FromEnvironment().ProxyFunc()
	httpTransport := &http.Transport{
		Proxy: func(r *http.Request) (uri *url.URL, err error) {
			return proxyFunc(r.URL)
		},
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2:     true,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	downloader := &ytdl.Downloader{
		OutputDir: outputDir,
	}

	downloader.Client.Debug = false
	downloader.HTTPClient = &http.Client{Transport: httpTransport}

	return downloader
}
