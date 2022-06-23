package video

import (
	"os"
	"path"

	"github.com/comunidade-shallom/diakonos/pkg/config"
	"github.com/comunidade-shallom/diakonos/pkg/files"
	"github.com/comunidade-shallom/diakonos/pkg/merge"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

var CmdMerge = &cli.Command{
	Name:  "merge",
	Usage: "merge a list of videos in a single video",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "define the name of generated video",
			DefaultText: "hash from input",
			Required:    false,
		},
	},
	Action: func(c *cli.Context) error {
		sources, err := buildSources(c.Args().Slice())

		if err != nil {
			return err
		}

		for _, v := range sources {
			pterm.Debug.Println(v)
		}

		cfg := config.Ctx(c.Context)

		out, err := merge.MergeFiles(merge.MergeParams{
			OutputDir: cfg.Merge.OutputDir,
			Sources:   sources,
			Name:      c.String("name"),
		})

		if err == nil {
			pterm.Success.Printfln("Done: %s", files.GetRelative(out.Name))
		}

		return err
	},
}

func buildSources(sources []string) ([]string, error) {
	if len(sources) == 0 {
		return sources, ErrorMissingSourceArgument
	}

	pwd, _ := os.Getwd()

	for i, v := range sources {
		if path.IsAbs(v) {
			continue
		}

		sources[i] = path.Join(pwd, v)
	}

	return sources, nil
}
