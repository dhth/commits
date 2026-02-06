package cmd

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/dhth/commits/ui"
)

var errCouldntReadConfigFile = errors.New("couldn't read config file")

func readConfig(filePath string) (ui.RawConfig, error) {
	var config ui.RawConfig
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		return config, fmt.Errorf("%w: %w", errCouldntReadConfigFile, err)
	}

	return config, nil
}
