package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/extract"
	"github.com/comunidade-shallom/diakonos/pkg/fileutils"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
)

type Pipeline struct {
	Name    string             `yml:"name"`
	Actions []ActionDefinition `yml:"actions"`
	outputs map[string]Output
	cfg     config.AppConfig
	ctx     context.Context
}

var (
	ErrInvalidActionType          = errors.Business("invalid action type [%s/%s]", "DP:002")
	ErrSourceIsEmpty              = errors.Business("source is empty [%s/%s]", "DP:003")
	ErrActionSourceIsNotAvailable = errors.Business("action source (%s) is not available [%s/%s]", "DP:004")
)

func (p *Pipeline) Run(ctx context.Context, cfg config.AppConfig) (map[string]Output, error) {
	p.cfg = cfg
	p.ctx = ctx

	p.outputs = make(map[string]Output)

	for _, act := range p.Actions {
		out, err := p.runAction(act)

		if err != nil {
			return p.outputs, err
		}

		pterm.Success.Printfln("Generated: %s", fileutils.GetRelative(out.Filename))

		if act.ID == "" {
			pterm.Warning.Printf("Missing action id (%s)", act.Type)
		} else {
			p.outputs[act.ID] = out
		}
	}

	return p.outputs, nil
}

func (p Pipeline) getSource(act ActionDefinition) (Source, error) {
	s := act.Source

	if s.Value != "" {
		return Source{
			Value: s.Value,
		}, nil
	}

	if s.Action == "" {
		return Source{}, ErrSourceIsEmpty.Msgf(act.ID, act.Type)
	}

	out, ok := p.outputs[s.Action]

	if !ok {
		return Source{}, ErrActionSourceIsNotAvailable.Msgf(s.Action, act.ID, act.Type)
	}

	return Source{
		Value: out.Filename,
	}, nil
}

func (p Pipeline) runAction(act ActionDefinition) (Output, error) {
	switch act.Type {
	case Download:
		return p.runDownload(act)
	case CutVideo:
		return p.runCutVideo(act)
	case ExtractAudio:
		return p.runExtractAudio(act)
	default:
		return Output{}, ErrInvalidActionType.Msgf(act.ID, act.Type)
	}
}

func (p Pipeline) runDownload(act ActionDefinition) (Output, error) {
	source, err := p.getSource(act)

	if err != nil {
		return Output{}, err
	}

	params, err := download.Params(act.Params, p.cfg.Download)

	if err != nil {
		return Output{}, err
	}

	params.From = source.Value

	f, _, err := download.YouTube(p.ctx, params)

	if err != nil {
		if e, ok := err.(errors.BusinessError); ok && e.ErrorCode == download.ErrExist.ErrorCode {
			pterm.Warning.Println(e.Error())
		} else {
			return Output{}, err
		}
	}

	return Output{
		Filename: f.Name,
	}, nil
}

func (p Pipeline) runCutVideo(act ActionDefinition) (Output, error) {
	source, err := p.getSource(act)

	if err != nil {
		return Output{}, err
	}

	params, err := cut.Params(act.Params, p.cfg.Cut)

	if err != nil {
		return Output{}, err
	}

	params.Source = source.Value

	o, err := cut.CutFile(params)

	return Output{
		Filename: o.Name,
	}, err
}

func (p Pipeline) runExtractAudio(act ActionDefinition) (Output, error) {
	source, err := p.getSource(act)

	if err != nil {
		return Output{}, err
	}

	params, err := extract.Params(act.Params, p.cfg.Audio)

	if err != nil {
		return Output{}, err
	}

	params.Source = source.Value

	o, err := extract.Audio(params)

	return Output{
		Filename: o.Name,
	}, err
}
