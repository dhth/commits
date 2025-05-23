package cmd

import (
	"github.com/BurntSushi/toml"
)

type cliConfig struct {
	IgnorePattern *string  `toml:"ignore_pattern"`
	EditorCmd     []string `toml:"editor_command"`
}

func readConfig(filePath string) (cliConfig, error) {
	var config cliConfig
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
