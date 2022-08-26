package audios

import (
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
)

type Config struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"audios"`
}

type Params struct {
	Source    string
	OutputDir string
}

var ErrExist = errors.Business("file already exist (%s)", "DE:001")

func (c Config) FromRaw(raw collection.Params) (Params, error) {
	return c.Apply(Params{
		Source:    raw.String("source"),
		OutputDir: raw.String("output_dir"),
	})
}

func (c Config) Apply(p Params) (Params, error) {
	if p.OutputDir == "" {
		p.OutputDir = c.OutputDir
	}

	return p, nil
}

func (p Params) Filename() string {
	return files.ChangeLocation(p.Source, p.OutputDir, "mp3")
}

func (p Params) WithPrefix(prefix string) string {
	return files.AddPrefix(p.Filename(), prefix)
}

func (p Params) SouceRelative() string {
	return files.GetRelative(p.Source)
}
