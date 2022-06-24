package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/files"
)

func (p Pipeline) runCutVideo(ctx context.Context, act ActionDefinition) (files.Output, error) {
	source, err := p.getSource(act)
	if err != nil {
		return files.Output{}, err
	}

	params, err := p.cfg.Cut.Params(act.Params)
	if err != nil {
		return files.Output{}, err
	}

	params.Source = source.Value

	return cut.Video(params)
}
