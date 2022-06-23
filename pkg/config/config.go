package config

import (
	"context"

	"github.com/comunidade-shallom/diakonos/pkg/audios"
	"github.com/comunidade-shallom/diakonos/pkg/cut"
	"github.com/comunidade-shallom/diakonos/pkg/download"
)

type ctxKey struct{}

type AudioOptions struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/audios"`
}

type MergeOptions struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/merges"`
}

type AppConfig struct {
	Download download.Config
	Cut      cut.Config
	Audio    audios.Config
	Merge    MergeOptions
}

var current AppConfig

func Ctx(ctx context.Context) AppConfig {
	cf, _ := ctx.Value(ctxKey{}).(AppConfig)

	return cf
}

func (c AppConfig) WithContext(ctx context.Context) context.Context {
	if cf, ok := ctx.Value(ctxKey{}).(AppConfig); ok {
		if cf == c {
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey{}, c)
}
