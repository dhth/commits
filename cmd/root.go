package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/dhth/commits/ui"
)

func die(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

var (
	path          = flag.String("path", "", "path to the repo")
	ignorePattern = flag.String("ignore-pattern", "", "ignore commit messages that match this regex")
)

func Execute() {
	flag.Parse()

	var repoPath string
	var err error

	if *path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			die("Couldn't get current working directory: %s", err.Error())
		}
		repoPath = cwd
	} else {
		repoPath, err = expandTilde(*path)
		if err != nil {
			die("Couldn't expand path: %s", err.Error())
		}
	}

	config := ui.Config{
		Path:          repoPath,
		IgnorePattern: *ignorePattern,
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
