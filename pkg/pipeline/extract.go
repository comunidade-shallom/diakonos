package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
)

func (p Pipeline) runExtractAudio(ctx context.Context, act ActionDefinition) (files.Output, collection.Params, error) {
	source, err := p.getSource(act.Source)
	if err != nil {
		return files.Output{}, collection.Params{}, err
	}

	params, err := p.cfg.Audio.FromRaw(act.Params)
	if err != nil {
		return files.Output{}, collection.Params{}, err
	}

	params.Source = source.Value

	out, err := audios.Extract(ctx, params)

	return out, collection.Params{}, err
}
