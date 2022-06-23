package cut

import (
	"fmt"
	"os"
	"path"
	"time"
)

type Config struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/cuts"`
}

type Params struct {
	Source    string
	OutputDir string
	Start     time.Duration
	Finish    time.Duration
}

func (c Config) Params(raw map[string]string) (Params, error) {
	params := Params{
		Source:    raw["source"],
		OutputDir: raw["output_dir"],
	}

	start, err := time.ParseDuration(raw["start"])
	if err != nil {
		return params, err
	}

	params.Start = start

	finish, err := time.ParseDuration(raw["finish"])
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
