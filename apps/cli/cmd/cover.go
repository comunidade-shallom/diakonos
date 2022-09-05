//nolint:gomnd
package cmd

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/covers"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/support"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/gosimple/slug"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gopkg.in/fogleman/gg.v1"
)

var ErrFailToGenerateImage = errors.System(nil, "Fail to generate image", "C:001")

var CmdCover = &cli.Command{
	Name:  "cover",
	Usage: "generate covers",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "text",
			Required: true,
		},
		&cli.IntFlag{
			Name:        "size",
			DefaultText: "1080",
			Value:       1080,
		},
		&cli.StringFlag{
			Name:        "prefix",
			DefaultText: strconv.FormatInt(time.Now().Unix(), 16),
			Value:       strconv.FormatInt(time.Now().Unix(), 16),
		},
		&cli.IntFlag{
			Name:        "times",
			DefaultText: "5",
			Value:       5,
		},
		&cli.Float64Flag{
			Name:        "font-size",
			DefaultText: "130",
			Value:       130,
		},
		&cli.StringFlag{
			Name: "out_dir",
		},
	},
	Action: func(cmd *cli.Context) error {
		cfg := config.Ctx(cmd.Context)
		text := cmd.String("text")
		outDir := cmd.String("out_dir")

		if outDir == "" {
			outDir = "covers/" + slug.Make(text)
		}

		if !filepath.IsAbs(outDir) {
			outDir = filepath.Join(cfg.BaseOutputDir, outDir)
		}

		if err := support.EnsureDir(outDir); err != nil {
			return nil
		}

		prefix := cmd.String("prefix")
		times := cmd.Int("times")

		progressBar, err := pterm.DefaultProgressbar.
			WithTotal(times).
			WithTitle("Generating covers").
			WithRemoveWhenDone().
			Start()
		if err != nil {
			return err
		}

		generator := covers.Generator{
			Sources:  cfg.Sources,
			Size:     cmd.Int("size"),
			FontSize: cmd.Float64("font-size"),
			Text:     text,
		}

		for count := 0; count < times; count++ {
			name := filepath.Join(outDir, fmt.Sprintf("%s (%03d).png", prefix, count+1))
			progressBar.UpdateTitle("Generating " + files.GetRelative(name))

			img, err := generator.Generate()
			if err != nil {
				return err
			}

			err = gg.SavePNG(name, img)

			if err != nil {
				return ErrFailToGenerateImage.WithErr(err)
			}

			pterm.Success.Printfln("Generated: %s", files.GetRelative(name))

			progressBar.Increment()
		}

		return nil
	},
}
