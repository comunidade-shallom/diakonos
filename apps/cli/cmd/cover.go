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
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gopkg.in/fogleman/gg.v1"
)

var CmdCover = &cli.Command{
	Name:  "cover",
	Usage: "generate cover",
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
		&cli.StringFlag{
			Name:        "out_dir",
			DefaultText: "covers/" + time.Now().Format("20060201"),
			Value:       "covers/" + time.Now().Format("20060201"),
		},
	},
	Action: func(cmd *cli.Context) error {
		cfg := config.Ctx(cmd.Context)

		outDir := cmd.String("out_dir")

		if !filepath.IsAbs(outDir) {
			outDir = filepath.Join(cfg.BaseOutputDir, outDir)
		}

		if err := support.EnsureDir(outDir); err != nil {
			return nil
		}

		prefix := cmd.String("prefix")
		times := cmd.Int("times")

		fmt.Printf("outDir: %v\n", outDir)
		fmt.Printf("prefix: %v\n", prefix)
		fmt.Printf("times: %v\n", times)

		for i := 0; i < times; i++ {
			img, err := covers.Generator{
				Sources: cfg.Sources,
				Size:    cmd.Int("size"),
				Text:    cmd.String("text"),
			}.Generate()
			if err != nil {
				return err
			}

			name := filepath.Join(outDir, fmt.Sprintf("%s (%v).png", prefix, i))

			err = gg.SavePNG(name, img)

			if err != nil {
				return err
			}

			pterm.Success.Printfln("Generated: %s", files.GetRelative(name))
		}

		return nil
	},
}
