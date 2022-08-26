package config

import (
	goErrors "errors"
	"os"
	"path"
	"path/filepath"

	"github.com/comunidade-shallom/diakonos/pkg/support"
	"github.com/comunidade-shallom/diakonos/pkg/support/errors"
	"github.com/creasty/defaults"
	"github.com/kkyr/fig"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

var (
	ErrFailToLoadConfig  = errors.System(nil, "fail to load config", "DCONF:001")
	ErrFailEnsureConfig  = errors.System(nil, "fail to ensure config", "DCONF:002")
	ConfigFileWasCreated = errors.Business("a new config file was created (%s)", "DCONF:003")
)

func Load(file string) (AppConfig, error) {
	var err error

	cfg := AppConfig{}

	if file != "" {
		err = fig.Load(&cfg,
			fig.File(filepath.Base(file)),
			fig.Dirs(filepath.Dir(file)),
		)

		if err != nil {
			return cfg, ErrFailToLoadConfig.WithErr(err)
		}

		return applyDefaults(cfg)
	}

	home, err := homedir.Dir()
	if err != nil {
		return cfg, ErrFailToLoadConfig.WithErr(err)
	}

	err = fig.Load(&cfg,
		fig.File("diakonos.yml"),
		fig.Dirs(
			".",
			path.Join(home, ".diakonos"),
			path.Join(home, ".config"),
			path.Join(home, ".config/diakonos"),
			home,
			"/etc/diakonos",
			support.GetBinDirPath(),
		),
	)

	if goErrors.Is(err, fig.ErrFileNotFound) {
		return ensureConfig()
	}

	if err != nil {
		return cfg, err
	}

	return applyDefaults(cfg)
}

//nolint:cyclop
func applyDefaults(cfg AppConfig) (AppConfig, error) {
	pwd, _ := os.Getwd()

	if cfg.BaseOutputDir == "" {
		cfg.BaseOutputDir = path.Join(pwd, "outputs")
	}

	if !path.IsAbs(cfg.BaseOutputDir) {
		cfg.BaseOutputDir = path.Join(pwd, cfg.BaseOutputDir)
	}

	base := cfg.BaseOutputDir

	// download
	if cfg.Download.OutputDir == "" {
		cfg.Download.OutputDir = path.Join(base, "downloads")
	}

	if !path.IsAbs(cfg.Download.OutputDir) {
		cfg.Download.OutputDir = path.Join(base, cfg.Download.OutputDir)
	}

	if cfg.Download.MimeType == "" {
		cfg.Download.MimeType = "mp4"
	}

	if cfg.Download.Quality == "" {
		cfg.Download.Quality = "hd1080"
	}

	// cut
	if cfg.Cut.OutputDir == "" {
		cfg.Cut.OutputDir = path.Join(base, "cuts")
	}

	if !path.IsAbs(cfg.Cut.OutputDir) {
		cfg.Cut.OutputDir = path.Join(base, cfg.Cut.OutputDir)
	}

	// audios
	if cfg.Audio.OutputDir == "" {
		cfg.Audio.OutputDir = path.Join(base, "audios")
	}

	if !path.IsAbs(cfg.Audio.OutputDir) {
		cfg.Audio.OutputDir = path.Join(base, cfg.Audio.OutputDir)
	}

	// merged
	if cfg.Merge.OutputDir == "" {
		cfg.Merge.OutputDir = path.Join(base, "merges")
	}

	if !path.IsAbs(cfg.Merge.OutputDir) {
		cfg.Merge.OutputDir = path.Join(base, cfg.Merge.OutputDir)
	}

	return cfg, nil
}

func ensureConfig() (AppConfig, error) {
	var err error

	cfg := AppConfig{}

	if err = defaults.Set(&cfg); err != nil {
		return cfg, ErrFailEnsureConfig.WithErr(err)
	}

	cfg, err = applyDefaults(cfg)

	if err != nil {
		return cfg, ErrFailEnsureConfig.WithErr(err)
	}

	buf, err := yaml.Marshal(cfg)
	if err != nil {
		return cfg, ErrFailEnsureConfig.WithErr(err)
	}

	pwd, _ := os.Getwd()

	configFile := path.Join(pwd, "diakonos.yml")

	err = os.WriteFile(configFile, buf, os.ModePerm)

	if err != nil {
		return cfg, ErrFailEnsureConfig.WithErr(err)
	}

	return cfg, ConfigFileWasCreated.Msgf(configFile)
}
