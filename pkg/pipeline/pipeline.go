package pipeline

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
)

type (
	Action   string
	Pipeline struct {
		Name    string             `yml:"name"`
		Actions []ActionDefinition `yml:"actions"`
		outputs map[string]files.Output
		cfg     config.AppConfig
	}
)

const (
	YoutubeDownload   Action = "youtube-download"
	VideoCut          Action = "video-cut"
	VideoExtractAudio Action = "video-extract-audio"
	VideoMerge        Action = "video-merge"
	AudioNormalize    Action = "audio-normalize"
)

var (
	ErrInvalidActionType          = errors.Business("invalid action type [%s/%s]", "DP:002")
	ErrSourceIsEmpty              = errors.Business("source is empty", "DP:003")
	ErrSourceListIsEmpty          = errors.Business("source list is empty", "DP:004")
	ErrActionSourceIsNotAvailable = errors.Business("action source (%s) is not available [%s/%s]", "DP:005")
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

func (p Pipeline) getSource(act ActionSource) (Source, error) {
	if act.Value != "" {
		return Source{
			Value: act.Value,
		}, nil
	}

	if act.Action == "" {
		return Source{}, ErrSourceIsEmpty
	}

	out, ok := p.outputs[act.Action]

	if !ok {
		return Source{}, ErrActionSourceIsNotAvailable.Msgf(act.Action)
	}

	return Source{
		Value: out.Filename,
	}, nil
}

func (p Pipeline) getSources(sources []ActionSource) ([]Source, error) {
	list := []Source{}

	if len(sources) == 0 {
		return list, ErrSourceListIsEmpty
	}

	for _, sourceInput := range sources {
		src, err := p.getSource(sourceInput)
		if err != nil {
			return list, err
		}

		list = append(list, src)
	}

	return list, nil
}

func (p Pipeline) runAction(ctx context.Context, act ActionDefinition) (files.Output, error) {
	switch act.Type {
	case YoutubeDownload:
		return p.runDownload(ctx, act)
	case VideoCut:
		return p.runCutVideo(ctx, act)
	case VideoExtractAudio:
		return p.runExtractAudio(ctx, act)
	case AudioNormalize:
		return p.runAudioNormalize(ctx, act)
	case VideoMerge:
		return p.runMergeVideo(ctx, act)
	default:
		return files.Output{}, ErrInvalidActionType.Msgf(act.ID, act.Type)
	}
}
