package download

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/comunidade-shallom/diakonos/pkg/fileutils"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/gosimple/slug"
	youtube "github.com/kkdai/youtube/v2"
	ytdl "github.com/kkdai/youtube/v2/downloader"
	"github.com/pterm/pterm"
	"golang.org/x/net/http/httpproxy"
)

var ErrExist = errors.Business("file already exist (%s)", "DC:001")

func YouTube(ctx context.Context, options DownloadParams) (DownloadedFile, *youtube.Video, error) {
	out := DownloadedFile{
		DownloadParams: options,
	}
	client := youtubeClient(options.OutputDir)

	video, err := client.GetVideo(options.From)

	if err != nil {
		return out, video, err
	}

	name := fmt.Sprintf(
		"%s--%s.%s",
		slug.Make(video.Title), options.Quality, options.MimeType,
	)

	out.Name = path.Join(options.OutputDir, name)

	if fileutils.FileExists(out.Name) {
		return out, video, ErrExist.Msgf(fileutils.GetRelative(out.Name))
	}

	pterm.Info.Printfln("Downloading: %s", video.Title)
	pterm.Info.Printfln("Quality: %s", options.Quality)
	pterm.Info.Printfln("MimeType: %s", options.MimeType)
	pterm.Info.Printfln("Target: %s", fileutils.GetRelative(out.Name))

	err = client.DownloadComposite(ctx, name, video, options.Quality, options.MimeType)

	if err != nil {
		return out, video, err
	}

	return out, video, err
}

func youtubeClient(outputDir string) *ytdl.Downloader {
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
