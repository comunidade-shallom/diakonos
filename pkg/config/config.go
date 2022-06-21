package config

import "context"

type ctxKey struct{}

type DownloadOptions struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/downloads"`
	Quality   string `fig:"quality" yaml:"quality" default:"hd1080"`
	MimeType  string `fig:"mime_type" yaml:"mime_type" default:"mp4"`
}

type CutOptions struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/cuts"`
}

type AudioOptions struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/audios"`
}

type MergeOptions struct {
	OutputDir string `fig:"output_dir" yaml:"output_dir" default:"outputs/merges"`
}

type AppConfig struct {
	Download DownloadOptions
	Cut      CutOptions
	Audio    AudioOptions
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
