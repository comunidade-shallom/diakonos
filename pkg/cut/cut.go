package cut

import (
	"fmt"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var ErrExist = errors.Business("file already exist (%s)", "DC:001")

func CutFile(options Params) (files.Output, error) {
	out := files.Output{
		Filename: options.Filename(),
	}

	if out.Exists() {
		return out, ErrExist.Msgf(out.NameRelative())
	}

	pterm.Info.Printfln("Cropping: %s", path.Base(options.Source))
	pterm.Info.Printfln("Start: %s", options.Start)
	pterm.Info.Printfln("Finish: %s", options.Finish)
	pterm.Debug.Printfln("Target: %s", options.Source)
	pterm.Debug.Printfln("OutputDir: %s", options.OutputDir)
	pterm.Debug.Printfln("Target: %s", out.Filename)

	err := ffmpeg.
		Input(options.Source).
		Output(out.Filename, ffmpeg.KwArgs{
			"ss":  fmt.Sprintf("%f", options.Start.Seconds()),
			"to":  fmt.Sprintf("%f", options.Finish.Seconds()),
			"c:v": "copy",
			"c:a": "copy",
		}).
		Run()

	return out, err
}

// ffmpeg -i input.mp4 -ss 00:05:10 -to 00:15:30 -c:v copy -c:a copy output2.mp4
