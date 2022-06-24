//nolint:errorlint,containedctx
package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
)

type Pipeline struct {
	Name    string             `yml:"name"`
	Actions []ActionDefinition `yml:"actions"`
	outputs map[string]files.Output
	cfg     config.AppConfig
	ctx     context.Context
}

var (
	ErrInvalidActionType          = errors.Business("invalid action type [%s/%s]", "DP:002")
	ErrSourceIsEmpty              = errors.Business("source is empty [%s/%s]", "DP:003")
	ErrActionSourceIsNotAvailable = errors.Business("action source (%s) is not available [%s/%s]", "DP:004")
)

func (p *Pipeline) Run(ctx context.Context, cfg config.AppConfig) (map[string]files.Output, error) {
	p.cfg = cfg
	p.ctx = ctx

	p.outputs = make(map[string]files.Output)

	for _, act := range p.Actions {
		out, err := p.runAction(act)
		if err != nil {
			return p.outputs, err
		}

		pterm.Success.Printfln("Generated: %s", files.GetRelative(out.Filename))

		if act.ID == "" {
			pterm.Warning.Printf("Missing action id (%s)", act.Type)
		} else {
			p.outputs[act.ID] = out
		}
	}

	return p.outputs, nil
}

func (p Pipeline) getSource(act ActionDefinition) (Source, error) {
	source := act.Source

	if source.Value != "" {
		return Source{
			Value: source.Value,
		}, nil
	}

	if source.Action == "" {
		return Source{}, ErrSourceIsEmpty.Msgf(act.ID, act.Type)
	}

	out, ok := p.outputs[source.Action]

	if !ok {
		return Source{}, ErrActionSourceIsNotAvailable.Msgf(source.Action, act.ID, act.Type)
	}

	return Source{
		Value: out.Filename,
	}, nil
}

func (p Pipeline) runAction(act ActionDefinition) (files.Output, error) {
	switch act.Type {
	case Download:
		return p.runDownload(act)
	case CutVideo:
		return p.runCutVideo(act)
	case ExtractAudio:
		return p.runExtractAudio(act)
	default:
		return files.Output{}, ErrInvalidActionType.Msgf(act.ID, act.Type)
	}
}

func (p Pipeline) runDownload(act ActionDefinition) (files.Output, error) {
	source, err := p.getSource(act)
	if err != nil {
		return files.Output{}, err
	}

	params, err := p.cfg.Download.FromRaw(act.Params)
	if err != nil {
		return files.Output{}, err
	}

	params.Source = source.Value

	out, _, err := download.YouTube(p.ctx, params)
	if err != nil {
		if e, ok := err.(errors.BusinessError); ok && e.ErrorCode == download.ErrExist.ErrorCode {
			pterm.Warning.Println(e.Error())
		} else {
			return out, err
		}
	}

	return out, nil
}

func (p Pipeline) runCutVideo(act ActionDefinition) (files.Output, error) {
	source, err := p.getSource(act)
	if err != nil {
		return files.Output{}, err
	}

	params, err := p.cfg.Cut.Params(act.Params)
	if err != nil {
		return files.Output{}, err
	}

	params.Source = source.Value

	return cut.CutFile(params)
}

func (p Pipeline) runExtractAudio(act ActionDefinition) (files.Output, error) {
	source, err := p.getSource(act)
	if err != nil {
		return files.Output{}, err
	}

	params, err := p.cfg.Audio.FromRaw(act.Params)
	if err != nil {
		return files.Output{}, err
	}

	params.Source = source.Value

	return audios.Extract(params)
}
