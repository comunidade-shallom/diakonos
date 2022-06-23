package download

import "github.com/comunidade-shallom/diakonos/pkg/config"

type DownloadParams struct {
	From      string
	OutputDir string
	Quality   string
	MimeType  string
}

type DownloadedFile struct {
	DownloadParams
	Name string
}

func Params(raw map[string]string, def config.DownloadOptions) (DownloadParams, error) {
	p := DownloadParams{
		From:      raw["from"],
		Quality:   raw["quality"],
		OutputDir: raw["output_dir"],
		MimeType:  raw["mime_type"],
	}

	if p.OutputDir == "" {
		p.OutputDir = def.OutputDir
	}

	if p.Quality == "" {
		p.Quality = def.Quality
	}

	if p.MimeType == "" {
		p.MimeType = def.MimeType
	}

	return p, nil
}
