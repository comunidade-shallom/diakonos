package cut

import (
	"fmt"
	"path"

	"github.com/mowshon/moviego"
	"github.com/pterm/pterm"
)

func CutFile(options CutParams) (CroppedFile, error) {
	out := CroppedFile{
		CutParams: options,
	}

	video, err := moviego.Load(options.Source)

	if err != nil {
		return out, err
	}

	name := fmt.Sprintf(
		"%v-%v--%s",
		options.Start.Seconds(),
		options.Finish.Seconds(),
		path.Base(options.Source),
	)

	out.Name = path.Join(options.OutputDir, name)

	pterm.Info.Printfln("Cropping: %s", path.Base(options.Source))
	pterm.Info.Printfln("Start: %s", options.Start)
	pterm.Info.Printfln("Finish: %s", options.Finish)

	err = video.
		SubClip(options.Start.Seconds(), options.Finish.Seconds()).
		Output(out.Name).
		Run()

	return out, err
}
