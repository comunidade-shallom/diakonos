package cut

import (
	"time"

	"github.com/comunidade-shallom/diakonos/pkg/config"
)

type CutParams struct {
	Source    string
	OutputDir string
	Start     time.Duration
	Finish    time.Duration
}

type CroppedFile struct {
	CutParams
	Name string
}

func Params(raw map[string]string, def config.CutOptions) (CutParams, error) {
	p := CutParams{
		Source:    raw["source"],
		OutputDir: raw["output_dir"],
	}

	if p.OutputDir == "" {
		p.OutputDir = def.OutputDir
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

	return p, nil
}
