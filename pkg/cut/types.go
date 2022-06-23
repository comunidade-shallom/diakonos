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
	p := Params{
		Source:    raw["source"],
		OutputDir: raw["output_dir"],
	}

	start, err := time.ParseDuration(raw["start"])
	if err != nil {
		return p, err
	}

	p.Start = start

	finish, err := time.ParseDuration(raw["finish"])
	if err != nil {
		return p, err
	}

	p.Finish = finish

	return c.Apply(p)
}

func (c Config) Apply(p Params) (Params, error) {
	if p.OutputDir == "" {
		p.OutputDir = c.OutputDir
	}

	if !path.IsAbs(p.Source) {
		pwd, _ := os.Getwd()
		p.Source = path.Join(pwd, p.Source)
	}

	return p, nil
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
