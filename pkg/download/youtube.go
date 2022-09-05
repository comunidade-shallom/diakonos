//nolint:gomnd
package download

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	youtube "github.com/kkdai/youtube/v2"
	ytdl "github.com/kkdai/youtube/v2/downloader"
	"github.com/pterm/pterm"
	"golang.org/x/net/http/httpproxy"
)

func YouTube(ctx context.Context, params Params) (files.Output, *youtube.Video, error) {
	out := files.Output{}

	client := youtubeClient(params.OutputDir)

	//nolint:contextcheck
	video, err := client.GetVideo(params.Source)
	if err != nil {
		return out, video, err
	}

	out.Filename = params.Filename(video.Title)

	if out.Exists() {
		return out, video, ErrExist.Msgf(out.NameRelative())
	}

	pterm.Info.Printfln("Downloading: %s", video.Title)
	pterm.Info.Printfln("Quality: %s", params.Quality)
	pterm.Info.Printfln("MimeType: %s", params.MimeType)
	pterm.Info.Printfln("Target: %s", out.NameRelative())
	pterm.Debug.Printfln("Target: %s", out.Filename)
	pterm.Debug.Printfln("OutputDir: %s", params.OutputDir)

	err = client.DownloadComposite(ctx, out.Name(), video, params.Quality, params.MimeType)

	return out, video, err
}

func youtubeClient(outputDir string) *ytdl.Downloader {
	proxyFunc := httpproxy.FromEnvironment().ProxyFunc()

	httpTransport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return proxyFunc(req.URL)
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
