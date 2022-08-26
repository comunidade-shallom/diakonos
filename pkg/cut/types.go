package cut

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
)

type Config struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"cuts"`
}

type Params struct {
	Source    string
	OutputDir string
	Start     time.Duration
	Finish    time.Duration
}

func (c Config) Params(raw collection.Params) (Params, error) {
	params := Params{
		Source:    raw.String("source"),
		OutputDir: raw.String("output_dir"),
	}

	start, err := time.ParseDuration(raw.String("start"))
	if err != nil {
		return params, err
	}

	params.Start = start

	finish, err := time.ParseDuration(raw.String("finish"))
	if err != nil {
		return params, err
	}

	params.Finish = finish

	return c.Apply(params)
}

func (c Config) Apply(params Params) (Params, error) {
	if params.OutputDir == "" {
		params.OutputDir = c.OutputDir
	}

	if !path.IsAbs(params.Source) {
		pwd, _ := os.Getwd()
		params.Source = path.Join(pwd, params.Source)
	}

	return params, nil
}

func (p Params) Filename() string {
	name := fmt.Sprintf(
		"%v-%v--%s",
		p.Start.Seconds(),
		p.Finish.Seconds(),
		path.Base(p.Source),
	)

	return path.Join(p.OutputDir, name)
}
