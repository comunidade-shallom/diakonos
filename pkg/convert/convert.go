package convert

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ToMP4(_ context.Context, params Params) (files.Output, error) {
	out := files.Output{
		Filename: params.Filename(params.Prefix, extMP4),
	}

	if out.Exists() {
		return out, errors.ErrExist.Msgf(out.NameRelative())
	}

	pterm.Info.Printfln("Extraction audio from: %s", params.SouceRelative())
	pterm.Info.Printfln("Generating: %s", out.NameRelative())

	// ffmpeg -y -i $path -acodec libmp3lame -ar 44100 -ac 1 -vcodec libx264 $dirname/$basename.mp4

	err := ffmpeg.Input(params.Source).
		Output(out.Filename, ffmpeg.KwArgs{
			"movflags":  "faststart",
			"acodec":    "libmp3lame",
			"ar":        "44100",
			"ac":        "2",
			"vcodec":    "libx264",
			"profile:v": "high",
			"level":     "4.0",
			"preset":    params.Preset,
		}).
		Run()

	return out, err
}
