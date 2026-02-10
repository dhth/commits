package cmd

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dhth/commits/ui"
	"github.com/go-git/go-git/v5"
)

var (
	configDirName  = "commits"
	configFileName = "commits.toml"
	ignorePattern  = flag.String("ignore-pattern", "", "ignore commit messages that match this regex")
)

var (
	errCouldntGetUserHomeDir   = errors.New("couldn't get your home directory")
	errCouldntGetUserConfigDir = errors.New("couldn't get your config directory")
	errConfigFileDoesntExist   = errors.New("config file doesn't exist")
	errCouldntReadConfigFile   = errors.New("couldn't read config file")
	errCouldntParseConfig      = errors.New("couldn't parse config")
	errConfigIsInvalid         = errors.New("config is invalid")
	errCouldntGetRepo          = errors.New("couldn't open git repo")
	errCouldntGetCwd           = errors.New("couldn't get current working directory")
)

func Execute() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("%w: %w", errCouldntGetUserHomeDir, err)
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("%w: %w", errCouldntGetUserConfigDir, err)
	}
	defaultConfigPath := filepath.Join(configDir, configDirName, configFileName)

	configFilePath := flag.String("config-file-path", defaultConfigPath, "location of commits' config file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `commits lets you glance at git commits through a simple TUI

Usage: commits [flags]

Flags:
`)
		flag.PrintDefaults()
	}
	flag.Parse()

	if *configFilePath == "" {
		return fmt.Errorf("config-file-path cannot be empty")
	}

	configPathToUse := expandPath(*configFilePath, homeDir)

	repoPath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("%w: %w", errCouldntGetCwd, err)
	}

	configBytes, err := os.ReadFile(configPathToUse)
	if errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("%w at %q", errConfigFileDoesntExist, configPathToUse)
	} else if err != nil {
		return fmt.Errorf("%w: %w", errCouldntReadConfigFile, err)
	}

	rawCfg, err := parseConfig(string(configBytes))
	if err != nil {
		return fmt.Errorf("%w: %w", errCouldntParseConfig, err)
	}

	if *ignorePattern != "" {
		rawCfg.IgnorePattern = ignorePattern
	}

	config, err := ui.NewConfig(rawCfg)
	if err != nil {
		return fmt.Errorf("%w: %w", errConfigIsInvalid, err)
	}

	repo, err := git.PlainOpenWithOptions(repoPath, &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return fmt.Errorf("%w: %w", errCouldntGetRepo, err)
	}

	return ui.RenderUI(repo, config)
}

func expandPath(path string, homeDir string) string {
	pathWithoutTilde, found := strings.CutPrefix(path, "~/")
	if !found {
		return path
	}

	return filepath.Join(homeDir, pathWithoutTilde)
}
