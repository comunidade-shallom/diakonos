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

func SetConfig(cfg AppConfig) {
	current = cfg
}

func Load(file string) (cfg AppConfig, err error) {
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

func applyDefaults(cfg AppConfig) (AppConfig, error) {
	pwd, _ := os.Getwd()

	var err error

	if cfg.Download.OutputDir == "" {
		cfg.Download.OutputDir = path.Join(pwd, "downloads")
	}

	if !path.IsAbs(cfg.Download.OutputDir) {
		cfg.Download.OutputDir = path.Join(pwd, cfg.Download.OutputDir)
	}

	if cfg.Download.MimeType == "" {
		cfg.Download.MimeType = "mp4"
	}

	if cfg.Download.Quality == "" {
		cfg.Download.Quality = "hd1080"
	}

	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func ensureConfig() (cfg AppConfig, err error) {
	err = defaults.Set(&cfg)

	if err != nil {
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