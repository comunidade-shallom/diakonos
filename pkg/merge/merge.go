package merge

import (
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/pterm/pterm"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var ErrExist = errors.Business("file already exist (%s)", "DM:001")

func MergeFiles(options MergeParams) (MergedFile, error) {
	out := MergedFile{
		MergeParams: options,
	}

	out.Name = options.filename()

	if out.fileExists() {
		return out, ErrExist.Msgf(files.GetRelative(out.Name))
	}

	tmp, err := options.tempFile()
	if err != nil {
		return out, err
	}

	defer os.Remove(tmp.Name())

	pterm.Info.Printfln("Generating: %s", path.Base(out.Name))

	err = ffmpeg.
		Input(tmp.Name(), ffmpeg.KwArgs{"f": "concat", "safe": "0"}).
		Output(out.Name, ffmpeg.KwArgs{"c": "copy"}).
		Run()

	return out, err
}

// ffmpeg -f concat -safe 0 -i videos.txt -c copy output9.mp4
