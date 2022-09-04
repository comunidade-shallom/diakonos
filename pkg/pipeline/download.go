package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
)

func (p Pipeline) runDownload(ctx context.Context, act ActionDefinition) (files.Output, collection.Params, error) {
	source, err := p.getSource(act.Source)
	if err != nil {
		return files.Output{}, collection.Params{}, err
	}

	params, err := p.cfg.Download.FromRaw(act.Params)
	if err != nil {
		return files.Output{}, collection.Params{}, err
	}

	params.Source = source.Value

	out, vid, err := download.YouTube(ctx, params)

	values := collection.Params{
		"source": params.Source,
	}

	if vid != nil {
		values["title"] = vid.Title
	}

	if err != nil {
		//nolint:errorlint
		if e, ok := err.(errors.BusinessError); ok && e.ErrorCode == download.ErrExist.ErrorCode {
			pterm.Warning.Println(e.Error())
		} else {
			return out, values, err
		}
	}

	return out, values, nil
}
