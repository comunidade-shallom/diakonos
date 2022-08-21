package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/merge"
)

func (p Pipeline) runMergeVideo(ctx context.Context, act ActionDefinition) (files.Output, error) {
	sources, err := p.getSources(act.Sources)
	if err != nil {
		return files.Output{}, err
	}

	srcs := make([]string, len(sources))

	for index, src := range sources {
		srcs[index] = src.Value
	}

	params, err := p.cfg.Merge.Apply(merge.Params{
		Sources: srcs,
		Name:    act.Params.String("name"),
	})
	if err != nil {
		return files.Output{}, err
	}

	return merge.Files(ctx, params)
}
