package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
)

func (p Pipeline) runDownload(ctx context.Context, act ActionDefinition) (files.Output, error) {
	source, err := p.getSource(act)
	if err != nil {
		return files.Output{}, err
	}

	params, err := p.cfg.Download.FromRaw(act.Params)
	if err != nil {
		return files.Output{}, err
	}

	params.Source = source.Value

	out, _, err := download.YouTube(ctx, params)
	if err != nil {
		//nolint:errorlint
		if e, ok := err.(errors.BusinessError); ok && e.ErrorCode == download.ErrExist.ErrorCode {
			pterm.Warning.Println(e.Error())
		} else {
			return out, err
		}
	}

	return out, nil
}
