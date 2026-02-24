package ui

import (
	"errors"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/go-git/go-git/v5"
)

var errCouldntSetupDebugLogging = errors.New("couldn't set up debug logging")

func RenderUI(repo *git.Repository, config Config) error {
	if os.Getenv("DEBUG") == "1" {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			return fmt.Errorf("%w: %w", errCouldntSetupDebugLogging, err)
		}
		defer func() { _ = f.Close() }()
	}

	p := tea.NewProgram(InitialModel(repo, config))
	_, err := p.Run()

	return err
}
