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

var (
	ErrFailToGenerateImage = errors.System(nil, "Fail to generate image", "C:001")
	ErrMissingText         = errors.Business("Text input must be defined", "C:002")
)

var CmdCover = &cli.Command{
	Name:  "cover",
	Usage: "generate covers",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "sizes",
		},
		&cli.IntFlag{
			Name:        "width",
			DefaultText: "1080",
			Value:       1080,
		},
		&cli.IntFlag{
			Name:        "height",
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
		outDir := cmd.String("out_dir")
		height := cmd.Int("height")
		width := cmd.Int("width")
		sizes := covers.ParseSizes(cmd.String("sizes"))
		text := cmd.Args().First()

		if text == "" {
			return ErrMissingText
		}

		if len(sizes) == 0 {
			sizes = append(sizes, covers.Size{
				Width:  width,
				Height: height,
			})
		}

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
		total := times * len(sizes)

		progressBar, err := pterm.DefaultProgressbar.
			WithTotal(total).
			WithTitle("Generating covers").
			WithRemoveWhenDone().
			Start()
		if err != nil {
			return err
		}

		generator := covers.GeneratorSource{
			Sources:  cfg.Sources,
			Width:    width,
			Height:   height,
			FontSize: cmd.Float64("font-size"),
			Text:     text,
		}

		for count := 0; count < times; count++ {
			builder, err := generator.Random()
			if err != nil {
				return err
			}

			for index, size := range sizes {
				name := filepath.Join(outDir, fmt.Sprintf("%s-%03d-%03d--%s.png", prefix, count+1, index+1, size.String()))

				progressBar.UpdateTitle("Generating " + files.GetRelative(name))

				err = gg.SavePNG(name, builder.WithSize(size).Build())
				if err != nil {
					return ErrFailToGenerateImage.WithErr(err)
				}

				pterm.Success.Printfln("Generated: %s", files.GetRelative(name))
				progressBar.Increment()
			}

		}

		return nil
	},
}
