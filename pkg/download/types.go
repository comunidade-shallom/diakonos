package download

import (
	"fmt"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/gosimple/slug"
)

type Config struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/downloads"`
	Quality   string `fig:"quality" yaml:"quality" default:"hd1080"`
	MimeType  string `fig:"mime_type" yaml:"mime_type" default:"mp4"`
}

type Params struct {
	OutputDir string
	Quality   string
	MimeType  string
	Source    string
}

type Output struct {
	files.Output
}

var ErrExist = errors.Business("file already exist (%s)", "DC:001")

func (c Config) FromRaw(raw map[string]string) (Params, error) {
	p := Params{
		Source:    raw["source"],
		Quality:   raw["quality"],
		OutputDir: raw["output_dir"],
		MimeType:  raw["mime_type"],
	}

	return c.Apply(p)
}

func (c Config) Apply(p Params) (Params, error) {
	if p.OutputDir == "" {
		p.OutputDir = c.OutputDir
	}

	if p.Quality == "" {
		p.Quality = c.Quality
	}

	if p.MimeType == "" {
		p.MimeType = c.MimeType
	}

	return p, nil
}

func (p Params) Filename(title string) string {
	name := fmt.Sprintf(
		"%s--%s.%s",
		slug.Make(title), p.Quality, p.MimeType,
	)

	return path.Join(p.OutputDir, name)
}
