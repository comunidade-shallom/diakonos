package merge

import (
	"os"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var ErrExist = errors.Business("file already exist (%s)", "DM:001")

func Files(options Params) (files.Output, error) {
	out := files.Output{
		Filename: options.Filename(),
	}

	if out.Exists() {
		return out, ErrExist.Msgf(out.NameRelative())
	}

	tmp, err := options.tempFile()
	if err != nil {
		return out, err
	}

	defer os.Remove(tmp.Name())

	pterm.Info.Printfln("Generating: %s", out.NameRelative())
	pterm.Debug.Printfln("Target: %s", out.Filename)
	pterm.Debug.Printfln("Temp file: %s", tmp.Name())

	err = ffmpeg.
		Input(tmp.Name(), ffmpeg.KwArgs{"f": "concat", "safe": "0"}).
		Output(out.Filename, ffmpeg.KwArgs{"c": "copy"}).
		Run()

	return out, err
}

// ffmpeg -f concat -safe 0 -i videos.txt -c copy output9.mp4
