package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
)

type Pipeline struct {
	Name    string             `yml:"name"`
	Actions []ActionDefinition `yml:"actions"`
	outputs map[string]files.Output
	cfg     config.AppConfig
}

var (
	ErrInvalidActionType          = errors.Business("invalid action type [%s/%s]", "DP:002")
	ErrSourceIsEmpty              = errors.Business("source is empty [%s/%s]", "DP:003")
	ErrActionSourceIsNotAvailable = errors.Business("action source (%s) is not available [%s/%s]", "DP:004")
)

func (p *Pipeline) Run(ctx context.Context, cfg config.AppConfig) (map[string]files.Output, error) {
	p.cfg = cfg

	p.outputs = make(map[string]files.Output)

	pterm.Info.Printfln("Starting pipeline: %s", p.Name)

	for _, act := range p.Actions {
		pterm.Debug.Printfln("Running %s", act.ID)

		out, err := p.runAction(ctx, act)
		if err != nil {
			pterm.Warning.Printfln("Finish pipeline with error: %s", err.Error())

			return p.outputs, err
		}

		pterm.Success.Printfln("Generated: %s", files.GetRelative(out.Filename))

		if act.ID == "" {
			pterm.Warning.Printf("Missing action id (%s)", act.Type)
		} else {
			p.outputs[act.ID] = out
		}
	}

	pterm.Success.Println("Finish pipeline")

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

func (p Pipeline) runAction(ctx context.Context, act ActionDefinition) (files.Output, error) {
	switch act.Type {
	case Download:
		return p.runDownload(ctx, act)
	case CutVideo:
		return p.runCutVideo(ctx, act)
	case ExtractAudio:
		return p.runExtractAudio(ctx, act)
	default:
		return files.Output{}, ErrInvalidActionType.Msgf(act.ID, act.Type)
	}
}
