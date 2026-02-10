package cmd

import (
	"github.com/BurntSushi/toml"
	"github.com/dhth/commits/ui"
)

func parseConfig(contents string) (ui.RawConfig, error) {
	var config ui.RawConfig
	_, err := toml.Decode(contents, &config)

	return config, err
}
