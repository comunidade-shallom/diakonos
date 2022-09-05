package audios

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/pterm/pterm"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Normalize(_ context.Context, params Params) (files.Output, error) {
	out := files.Output{
		Filename: params.WithPrefix("normalized-"),
	}

	if out.Exists() {
		return out, ErrExist.Msgf(out.NameRelative())
	}

	pterm.Debug.Printfln("Input: %s", params.Source)
	pterm.Info.Printfln("Normalizing audio from: %s", params.SouceRelative())
	pterm.Info.Printfln("Generating: %s", out.NameRelative())

	err := ffmpeg.Input(params.Source).
		Filter("loudnorm", ffmpeg.Args{
			// "I=-16:TP=-1.5:LRA=11",
		}).
		//nolint:contextcheck
		Output(out.Filename).
		Run()

	return out, err
}

// ./ffmpeg -i /path/to/input.wav -af loudnorm=I=-16:TP=-1.5:LRA=11:print_format=summary -f null -
