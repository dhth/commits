package ui

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
)

func RenderUI(repo *git.Repository, config Config) {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer func() {
			_ = f.Close()
		}()
	}
	p := tea.NewProgram(InitialModel(repo, config), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Something went wrong %s", err)
	}
}
