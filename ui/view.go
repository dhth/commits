package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/go-git/go-git/v5/plumbing"
)

func (m model) View() tea.View {
	var content string
	var footer string

	var statusBar string
	if m.message != "" {
		statusBar += RightPadTrim(m.message, m.terminalWidth)
	}

	switch m.activePane {
	case commitsList:
		content = m.commitListStyle.Render(m.commitsList.View())
	case commitDetails:
		var commitStatsVP string
		if !m.commitStatsVPReady {
			commitStatsVP = "\n Initializing..."
		} else {
			commitStatsVP = viewPortStyle.Render(fmt.Sprintf(" %s\n\n%s\n",
				commitStatsTitleStyle.Render("Commit Details"),
				commitDetailsStyle.Render(m.commitDetailsVP.View())))
		}
		content = commitStatsVP
	case branchList:
		content = m.commitListStyle.Render(m.branchList.View())
	case helpView:
		var helpVP string
		if !m.helpVPReady {
			helpVP = "\n Initializing..."
		} else {
			helpVP = viewPortStyle.Render(fmt.Sprintf(" %s\n\n%s\n",
				helpVPTitleStyle.Render("Help"),
				m.helpVP.View()))
		}
		content = helpVP
	}

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#282828")).
		Background(lipgloss.Color("#7c6f64"))

	var headRef string
	if m.currentRef != nil {
		headRef = "  " + headRefLabel(m.currentRef)
	}

	var helpMsg string
	if m.showHelp {
		helpMsg = " " + helpMsgStyle.Render("press ? for help")
	}

	var revisionRange string
	if m.revStart != nil && m.revEnd == nil {
		revisionRange = " " + revChoiceStyle.Render(fmt.Sprintf("%s..?", shortenHash(*m.revStart)))
	} else if m.revEnd != nil {
		revisionRange = " " + revChoiceStyle.Render(fmt.Sprintf("%s..%s", shortenHash(*m.revStart), shortenHash(*m.revEnd)))
	}

	footerStr := fmt.Sprintf("%s%s%s%s",
		modeStyle.Render("commits"),
		headRef,
		helpMsg,
		revisionRange,
	)
	footer = footerStyle.Render(footerStr)

	v := tea.NewView(lipgloss.JoinVertical(lipgloss.Left,
		content,
		statusBar,
		footer,
	))
	v.AltScreen = true

	return v
}

func headRefLabel(ref *plumbing.Reference) string {
	if ref.Name() == plumbing.HEAD {
		return fmt.Sprintf("%s detached at %s",
			headRefStyle.Render("HEAD"), headRefStyle.Render(shortenHash(ref.Hash().String())))
	}

	return fmt.Sprintf("%s -> %s", headRefStyle.Render("HEAD"), headRefStyle.Render(ref.Name().Short()))
}

func shortenHash(hash string) string {
	if len(hash) >= 10 {
		return hash[:10]
	}

	return hash
}
