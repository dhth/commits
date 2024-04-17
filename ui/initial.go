package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func InitialModel(config Config) model {

	commitListDel := defaultDelegate(lipgloss.Color(commitsListColor))

	branchListDel := defaultDelegate(lipgloss.Color(branchListColor))

	baseStyle = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		Foreground(lipgloss.Color(defaultBackgroundColor))

	tableListStyle := baseStyle.Copy().
		PaddingTop(1).
		PaddingRight(2).
		PaddingLeft(1).
		PaddingBottom(1)

	m := model{
		config:          config,
		commitsList:     list.New(nil, commitListDel, 0, 0),
		branchList:      list.New(nil, branchListDel, 0, 0),
		commitListStyle: tableListStyle,
		showHelp:        true,
	}

	m.commitsList.Title = "Commits"
	m.commitsList.SetStatusBarItemName("commit", "commits")
	m.commitsList.DisableQuitKeybindings()
	m.commitsList.SetShowHelp(false)
	m.commitsList.SetFilteringEnabled(false)
	m.commitsList.Styles.Title.Foreground(lipgloss.Color(defaultBackgroundColor))
	m.commitsList.Styles.Title.Background(lipgloss.Color(commitsListColor))
	m.commitsList.Styles.Title.Bold(true)

	m.branchList.Title = "Branches"
	m.branchList.SetStatusBarItemName("branch", "branches")
	m.branchList.DisableQuitKeybindings()
	m.branchList.SetShowHelp(false)
	m.branchList.Styles.Title.Foreground(lipgloss.Color(defaultBackgroundColor))
	m.branchList.Styles.Title.Background(lipgloss.Color(branchListColor))
	m.branchList.Styles.Title.Bold(true)

	return m
}
