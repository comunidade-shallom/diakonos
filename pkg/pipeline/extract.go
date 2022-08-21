package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/files"
)

func (p Pipeline) runExtractAudio(ctx context.Context, act ActionDefinition) (files.Output, error) {
	source, err := p.getSource(act.Source)
	if err != nil {
		return files.Output{}, err
	}

	params, err := p.cfg.Audio.FromRaw(act.Params)
	if err != nil {
		return files.Output{}, err
	}

	params.Source = source.Value

	return audios.Extract(ctx, params)
}
