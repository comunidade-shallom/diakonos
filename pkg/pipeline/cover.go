package pipeline

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/comunidade-shallom/diakonos/pkg/covers"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support/collection"
	"github.com/pterm/pterm"
	"gopkg.in/fogleman/gg.v1"
)

//nolint:funlen
func (p Pipeline) runCoverGenerate(
	_ context.Context,
	act ActionDefinition,
) (files.Output, collection.Params, error) {
	text := act.Params.String("text")
	times := act.Params.Int("times")
	prefix := act.Params.String("prefix")
	sizes := covers.BuildSizes(act.Params.Strings("sizes"))

	if len(sizes) == 0 {
		sizes = append(sizes, covers.Size{
			//nolint:gomnd
			Width: 1080,
			//nolint:gomnd
			Height: 1080,
		})
	}

	if times == 0 {
		times = 5
	}

	if prefix == "" {
		prefix = strconv.FormatInt(time.Now().Unix(), 16)
	}

	generator := covers.GeneratorSource{
		Sources: p.cfg.Sources,
		Text:    text,
	}

	pterm.Info.Printfln("Generating cover to: %s", text)

	targetDir := p.targetDir

	progressBar, err := pterm.DefaultProgressbar.
		WithTotal(times * len(sizes)).
		WithTitle("Generating covers").
		WithRemoveWhenDone().
		Start()
	if err != nil {
		return files.Output{}, collection.Params{}, err
	}

	for count := 0; count < times; count++ {
		builder, err := generator.Random()
		if err != nil {
			return files.Output{}, collection.Params{}, err
		}

		for index, size := range sizes {
			name := filepath.Join(targetDir, fmt.Sprintf("%s-%03d-%03d--%s.png", prefix, count+1, index+1, size.String()))
			progressBar.UpdateTitle("Generating " + files.GetRelative(name))

			err = gg.SavePNG(name, builder.WithSize(size).Build())

			if err != nil {
				return files.Output{}, collection.Params{}, err
			}

			pterm.Success.Printfln("Saved: %s", files.GetRelative(name))
			progressBar.Increment()
		}
	}

	return files.Output{}, collection.Params{}, nil
}
