package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func commitHashDelegate() list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.
		SelectedTitle.
		Foreground(lipgloss.Color(commitsListColor)).
		BorderLeftForeground(lipgloss.Color(commitsListColor))
	d.Styles.SelectedDesc = d.Styles.
		SelectedTitle.
		Copy()

	return d
}
