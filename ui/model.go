package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Pane uint

const (
	commitsList Pane = iota
	commitStats
)

type model struct {
	config             Config
	repoInfo           repoInfo
	commitsList        list.Model
	currentRev         string
	message            string
	commitListStyle    lipgloss.Style
	terminalHeight     int
	terminalWidth      int
	commitStatsVP      viewport.Model
	activePane         Pane
	commitStatsVPReady bool
	showHelp           bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		getCurrentRev(m.config.Path),
		getCommits(m.config.Path),
	)
}
