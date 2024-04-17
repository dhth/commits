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
				commitDetailsStyle.Render(m.commitDetailsVP.View())))
		}
		content = commitStatsVP
	case branchList:
		content = m.commitListStyle.Render(m.branchList.View())
	case helpView:
		var helpVP string
		if !m.helpVPReady {
			helpVP = "\n  Initializing..."
		} else {
			helpVP = viewPortStyle.Render(fmt.Sprintf("  %s\n\n%s\n",
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
		headRef = fmt.Sprintf("  %s -> %s", headRefStyle.Render("HEAD"), headRefStyle.Render(m.currentRef.Name().Short()))
	}

	var helpMsg string
	if m.showHelp {
		helpMsg = " " + helpMsgStyle.Render("press ? for help")
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
