package download

import (
	"fmt"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/gosimple/slug"
)

type Config struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"downloads"`
	Quality   string `fig:"quality" yaml:"quality" default:"hd1080"`
	MimeType  string `fig:"mime_type" yaml:"mime_type" default:"mp4"`
}

type Params struct {
	OutputDir string
	Quality   string
	MimeType  string
	Source    string
}

var ErrExist = errors.Business("file already exist (%s)", "DC:001")

func (c Config) FromRaw(raw collection.Params) (Params, error) {
	return c.Apply(Params{
		Source:    raw.String("source"),
		Quality:   raw.String("quality"),
		OutputDir: raw.String("output_dir"),
		MimeType:  raw.String("mime_type"),
	})
}

func (c Config) Apply(params Params) (Params, error) {
	if params.OutputDir == "" {
		params.OutputDir = c.OutputDir
	}

	if params.Quality == "" {
		params.Quality = c.Quality
	}

	if params.MimeType == "" {
		params.MimeType = c.MimeType
	}

	return params, nil
}

func (p Params) Filename(title string) string {
	name := fmt.Sprintf(
		"%s--%s.%s",
		slug.Make(title), p.Quality, p.MimeType,
	)

	return path.Join(p.OutputDir, name)
}
