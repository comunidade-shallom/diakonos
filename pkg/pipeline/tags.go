package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
)

func (p Pipeline) runDefineAudioTags(
	ctx context.Context,
	act ActionDefinition,
) (files.Output, collection.Params, error) {
	source, err := p.getSource(act.Source)
	if err != nil {
		return files.Output{}, collection.Params{}, err
	}

	params := audios.AudioTags{}.FromRaw(act.Params)

	err = audios.WriteTags(ctx, params, source.Value)

	return files.Output{
		Filename: source.Value,
	}, collection.Params{}, err
}
