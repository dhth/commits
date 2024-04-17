package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-git/go-git/v5/plumbing"
)

type Pane uint

const (
	commitsList Pane = iota
	commitDetails
	branchList
	helpView
)

type model struct {
	config             Config
	repoInfo           repoInfo
	commitsList        list.Model
	currentRef         *plumbing.Reference
	message            string
	branchList         list.Model
	commitListStyle    lipgloss.Style
	terminalHeight     int
	terminalWidth      int
	commitDetailsVP    viewport.Model
	revStartChosen     bool
	revEndChosen       bool
	revStartIndex      int
	revStart           string
	revEndIndex        int
	revEnd             string
	activePane         Pane
	lastPane           Pane
	commitStatsVPReady bool
	showHelp           bool
	helpVP             viewport.Model
	helpVPReady        bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		hideHelp(time.Minute*2),
		getCommits(m.config.Repo, nil),
	)
}
