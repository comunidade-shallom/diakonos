package extract

import "github.com/comunidade-shallom/diakonos/pkg/config"

type ExtractParams struct {
	Source    string
	OutputDir string
}

type ExtractedFile struct {
	ExtractParams
	Name string
}

func Params(raw map[string]string, def config.AudioOptions) (ExtractParams, error) {
	p := ExtractParams{
		Source:    raw["source"],
		OutputDir: raw["output_dir"],
	}

	if p.OutputDir == "" {
		p.OutputDir = def.OutputDir
	}

	return p, nil
}
