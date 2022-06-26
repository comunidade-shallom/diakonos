package audios

import (
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/pterm/pterm"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const bitrate = 192000

func Extract(params Params) (files.Output, error) {
	out := files.Output{
		Filename: params.Filename(),
	}

	if out.Exists() {
		return out, ErrExist.Msgf(out.NameRelative())
	}

	pterm.Info.Printfln("Extraction audio from: %s", params.SouceRelative())
	pterm.Info.Printfln("Generating: %s", out.NameRelative())

	err := ffmpeg.Input(params.Source).
		Audio().
		Output(out.Filename, ffmpeg.KwArgs{"f": "mp3", "b:a": bitrate, "vn": ""}).
		Run()

	return out, err
}

// ffmpeg -i video.mp4 -f mp3 -ab 192000 -vn music.mp3
