package extract

import (
	"github.com/comunidade-shallom/diakonos/pkg/fileutils"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var ErrExist = errors.Business("file already exist (%s)", "DE:001")

func Audio(options ExtractParams) (ExtractedFile, error) {
	out := ExtractedFile{
		ExtractParams: options,
	}

	out.Name = fileutils.ChangeLocation(options.Source, options.OutputDir, "mp3")

	if fileutils.FileExists(out.Name) {
		return out, ErrExist.Msgf(fileutils.GetRelative(out.Name))
	}

	pterm.Info.Printfln("Extraction audio from: %s", fileutils.GetRelative(options.Source))
	pterm.Info.Printfln("Generating: %s", fileutils.GetRelative(out.Name))

	err := ffmpeg.Input(options.Source).
		Audio().
		Output(out.Name, ffmpeg.KwArgs{"f": "mp3", "ab": 192000, "vn": ""}).
		Run()

	return out, err
}

// ffmpeg -i video.mp4 -f mp3 -ab 192000 -vn music.mp3
