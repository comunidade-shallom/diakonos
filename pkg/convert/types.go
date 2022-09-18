package convert

import (
	"os"
	"path"
	"strings"

	"github.com/comunidade-shallom/diakonos/pkg/files"
)

type (
	Preset  string
	Presets []Preset
)

const (
	UltraFast Preset = "ultrafast"
	SuperFast Preset = "superfast"
	VeryFast  Preset = "veryfast"
	Faster    Preset = "faster"
	Fast      Preset = "fast"
	Medium    Preset = "medium"
	Slow      Preset = "slow"
	Slower    Preset = "slower"
	VerySlow  Preset = "veryslow"
	Placebo   Preset = "placebo"
)

var PresetsAvailable = Presets{
	UltraFast,
	SuperFast,
	VeryFast,
	Faster,
	Fast,
	Medium,
	Slow,
	Slower,
	VerySlow,
	Placebo,
}

const extMP4 = "mp4"

type Config struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"conversions"`
}

type Params struct {
	Preset    Preset
	Source    string
	Prefix    string
	OutputDir string
}

func (c Config) Apply(params Params) (Params, error) {
	if params.OutputDir == "" {
		params.OutputDir = c.OutputDir
	}

	if !path.IsAbs(params.Source) {
		pwd, _ := os.Getwd()
		params.Source = path.Join(pwd, params.Source)
	}

	if params.Preset == "" {
		params.Preset = Fast
	}

	return params, nil
}

func (p Params) Filename(prefix, ext string) string {
	return files.AddPrefix(files.ChangeLocation(p.Source, p.OutputDir, ext), prefix)
}

func (p Params) SouceRelative() string {
	return files.GetRelative(p.Source)
}

func (p Presets) String() string {
	list := make([]string, len(p))

	for index, value := range p {
		list[index] = string(value)
	}

	return strings.Join(list, "|")
}

func (p Preset) String() string {
	return string(p)
}
