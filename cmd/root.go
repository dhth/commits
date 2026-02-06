package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/dhth/commits/ui"
	"github.com/go-git/go-git/v5"
)

func die(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

var ignorePattern = flag.String("ignore-pattern", "", "ignore commit messages that match this regex")

func Execute() {
	currentUser, err := user.Current()
	if err != nil {
		die("Something went horribly wrong. Let @dhth know about this error on github: ", err.Error())
	}

	var defaultConfigFP string
	if err == nil {
		defaultConfigFP = fmt.Sprintf("%s/.config/commits/commits.toml", currentUser.HomeDir)
	}

	configFilePath := flag.String("config-file-path", defaultConfigFP, "location of commits' config file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\nFlags:\n", helpText)
		flag.PrintDefaults()
	}
	flag.Parse()

	if *configFilePath == "" {
		die("config-file-path cannot be empty")
	}
	var configFPExpanded string
	if strings.Contains(*configFilePath, "~") {
		configFPExpanded, err = expandTilde(*configFilePath)
		if err != nil {
			die("Something went horribly wrong. Let @dhth know about this error on github: ", err.Error())
		}
	} else {
		configFPExpanded = *configFilePath
	}

	_, err = os.Stat(configFPExpanded)
	if os.IsNotExist(err) {
		die(cfgErrSuggestion(fmt.Sprintf("Error: file doesn't exist at %q", configFPExpanded)))
	}

	repoPath, err := os.Getwd()
	if err != nil {
		die("Couldn't get current working directory: %s", err.Error())
	}

	cfg, err := readConfig(configFPExpanded)
	if err != nil {
		die(cfgErrSuggestion(fmt.Sprintf("Error reading config: %s", err.Error())))
	}

	var ig string

	if *ignorePattern != "" {
		ig = *ignorePattern
	} else if cfg.IgnorePattern != nil {
		ig = *cfg.IgnorePattern
	}

	r, err := git.PlainOpen(repoPath)
	if err != nil {
		die("Couldn't fetch git repo: %s", err.Error())
	}

	config := ui.Config{
		IgnorePattern:   ig,
		OpenInEditorCmd: cfg.EditorCmd,
		Repo:            r,
	}

	ui.RenderUI(config)
}

func expandTilde(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		return strings.Replace(path, "~", usr.HomeDir, 1), nil
	}
	return path, nil
}
