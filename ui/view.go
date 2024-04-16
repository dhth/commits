package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var content string
	var footer string

	var statusBar string
	if m.revStartChosen && !m.revEndChosen {
		statusBar += "picking revision range: " + revChoiceStyle.Render(fmt.Sprintf("%s..?", m.revStart))
	} else if m.revEndChosen {
		statusBar += "revision range picked:  " + revChoiceStyle.Render(fmt.Sprintf("%s..%s", m.revStart, m.revEnd))
	}
	if m.message != "" {
		statusBar += RightPadTrim(m.message, m.terminalWidth)
	}

	switch m.activePane {
	case commitsList:
		content = m.commitListStyle.Render(m.commitsList.View())
	case commitDetails:
		var commitStatsVP string
		if !m.commitStatsVPReady {
			commitStatsVP = "\n  Initializing..."
		} else {
			commitStatsVP = viewPortStyle.Render(fmt.Sprintf("  %s\n\n%s\n",
				commitStatsTitleStyle.Render("Commit Details"),
				commitDetailsStyle.Render(m.commitStatsVP.View())))
		}
		content = commitStatsVP
	}

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#282828")).
		Background(lipgloss.Color("#7c6f64"))

	headRef := fmt.Sprintf("  %s -> %s", headRefStyle.Render("HEAD"), headRefStyle.Render(m.currentRev))

	var helpMsg string
	if m.showHelp {
		helpMsg = " " + helpMsgStyle.Render("ctrl+d: open in editor; enter: show commit; tab: commit details; ctrl+t: pick range; ctrl+x: clear range")
	}

	footerStr := fmt.Sprintf("%s%s%s",
		modeStyle.Render("commits"),
		headRef,
		helpMsg,
	)
	footer = footerStyle.Render(footerStr)

	return lipgloss.JoinVertical(lipgloss.Left,
		content,
		statusBar,
		footer,
	)
}
