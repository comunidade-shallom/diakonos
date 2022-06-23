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

func CutFile(options CutParams) (CroppedFile, error) {
	out := CroppedFile{
		CutParams: options,
	}

	name := fmt.Sprintf(
		"%v-%v--%s",
		options.Start.Seconds(),
		options.Finish.Seconds(),
		path.Base(options.Source),
	)

	out.Name = path.Join(options.OutputDir, name)

	if files.FileExists(out.Name) {
		return out, ErrExist.Msgf(files.GetRelative(out.Name))
	}

	pterm.Info.Printfln("Cropping: %s", path.Base(options.Source))
	pterm.Info.Printfln("Start: %s", options.Start)
	pterm.Info.Printfln("Finish: %s", options.Finish)

	err := ffmpeg.
		Input(options.Source).
		Output(out.Name, ffmpeg.KwArgs{
			"ss":  fmt.Sprintf("%f", options.Start.Seconds()),
			"to":  fmt.Sprintf("%f", options.Finish.Seconds()),
			"c:v": "copy",
			"c:a": "copy",
		}).
		Run()

	return out, err
}

// ffmpeg -i input.mp4 -ss 00:05:10 -to 00:15:30 -c:v copy -c:a copy output2.mp4
