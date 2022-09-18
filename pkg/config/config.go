package config

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/convert"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/download"
	"github.com/comunidade-shallom/diakonos/pkg/merge"
	"github.com/comunidade-shallom/diakonos/pkg/sources"
)

type ctxKey struct{}

type AppConfig struct {
	BaseOutputDir string `fig:"base_output_dir" yaml:"base_output_dir" default:"outputs"`
	Download      download.Config
	Cut           cut.Config
	Convert       convert.Config
	Audio         audios.Config
	Merge         merge.Config
	Sources       sources.Sources
}

func Ctx(ctx context.Context) *AppConfig {
	cf, _ := ctx.Value(ctxKey{}).(*AppConfig)

	return cf
}

func (c *AppConfig) WithContext(ctx context.Context) context.Context {
	if cf, ok := ctx.Value(ctxKey{}).(*AppConfig); ok {
		if cf == c {
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, c)
}

func (c AppConfig) WithOutput(dir string) AppConfig {
	c.Download.OutputDir = dir
	c.Cut.OutputDir = dir
	c.Audio.OutputDir = dir
	c.Merge.OutputDir = dir

	return c
}
