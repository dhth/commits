package ui

import (
	"fmt"
	"strings"
)

type RawConfig struct {
	IgnorePattern *string  `toml:"ignore_pattern"`
	ShowCommitCmd []string `toml:"show_commit_command"`
	ShowRangeCmd  []string `toml:"show_range_command"`
	OpenCommitCmd []string `toml:"open_commit_command"`
	OpenRangeCmd  []string `toml:"open_range_command"`
}

type Config struct {
	IgnorePattern string
	ShowCommitCmd []string
	ShowRangeCmd  []string
	OpenCommitCmd []string
	OpenRangeCmd  []string
}

func NewConfig(raw RawConfig) (Config, error) {
	var cfg Config

	err := validateCommitCmd(raw.ShowCommitCmd)
	if err != nil {
		return cfg, fmt.Errorf("invalid value for show_commit_command: %w", err)
	}

	err = validateRangeCmd(raw.ShowRangeCmd)
	if err != nil {
		return cfg, fmt.Errorf("invalid value for show_range_command: %w", err)
	}

	err = validateCommitCmd(raw.OpenCommitCmd)
	if err != nil {
		return cfg, fmt.Errorf("invalid value for open_commit_command: %w", err)
	}

	err = validateRangeCmd(raw.OpenRangeCmd)
	if err != nil {
		return cfg, fmt.Errorf("invalid value for open_range_command: %w", err)
	}

	if raw.IgnorePattern != nil {
		cfg.IgnorePattern = *raw.IgnorePattern
	}
	cfg.ShowCommitCmd = raw.ShowCommitCmd
	cfg.ShowRangeCmd = raw.ShowRangeCmd
	cfg.OpenCommitCmd = raw.OpenCommitCmd
	cfg.OpenRangeCmd = raw.OpenRangeCmd

	return cfg, nil
}

func validateCommitCmd(cmd []string) error {
	if len(cmd) == 0 {
		return nil
	}

	for _, arg := range cmd {
		if strings.Contains(arg, "{{hash}}") {
			return nil
		}
	}

	return fmt.Errorf("must contain {{hash}} placeholder")
}

func validateRangeCmd(cmd []string) error {
	if len(cmd) == 0 {
		return nil
	}

	hasBase := false
	hasHead := false
	for _, arg := range cmd {
		if strings.Contains(arg, "{{base}}") {
			hasBase = true
		}
		if strings.Contains(arg, "{{head}}") {
			hasHead = true
		}
	}

	if !hasBase || !hasHead {
		return fmt.Errorf("must contain both {{base}} and {{head}} placeholders")
	}

	return nil
}
