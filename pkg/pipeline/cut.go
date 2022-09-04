package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
)

func (p Pipeline) runCutVideo(ctx context.Context, act ActionDefinition) (files.Output, collection.Params, error) {
	source, err := p.getSource(act.Source)
	if err != nil {
		return files.Output{}, collection.Params{}, err
	}

	params, err := p.cfg.Cut.Params(act.Params)
	if err != nil {
		return files.Output{}, collection.Params{}, err
	}

	params.Source = source.Value

	out, err := cut.Video(ctx, params)

	return out, collection.Params{}, err
}
