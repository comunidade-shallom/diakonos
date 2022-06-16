package cut

import (
	"fmt"
	"path"

	"github.com/mowshon/moviego"
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

	err = video.
		SubClip(options.Start.Seconds(), options.Finish.Seconds()).
		Output(path.Join(options.OutputDir, name)).
		Run()

	return out, err
}
