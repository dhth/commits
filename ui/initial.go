package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func InitialModel(config Config) model {

	appDelegate := commitHashDelegate()

	baseStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		Foreground(lipgloss.Color(DefaultBackgroundColor))

	tableListStyle := baseStyle.Copy().
		PaddingTop(1).
		PaddingRight(2).
		PaddingLeft(1).
		PaddingBottom(1)

	m := model{
		config:          config,
		commitsList:     list.New(nil, appDelegate, 0, 0),
		commitListStyle: tableListStyle,
		showHelp:        true,
	}

	m.commitsList.Title = "Commits"
	m.commitsList.SetStatusBarItemName("commit", "commits")
	m.commitsList.DisableQuitKeybindings()
	m.commitsList.SetShowHelp(false)
	m.commitsList.SetFilteringEnabled(false)
	m.commitsList.Styles.Title.Foreground(lipgloss.Color(DefaultBackgroundColor))
	m.commitsList.Styles.Title.Background(lipgloss.Color(CommitsListColor))
	m.commitsList.Styles.Title.Bold(true)

	return m
}
