package ui

import (
	"hash/fnv"

	"github.com/charmbracelet/lipgloss"
)

const (
	defaultBackgroundColor = "#282828"
	commitsListColor       = "#fe8019"
	branchListColor        = "#8ec07c"
	modeColor              = "#b8bb26"
	helpMsgColor           = "#83a598"
	commitStatsTitleColor  = "#d3869b"
	hashColor              = "#d3869b"
	dateColor              = "#928374"
	commitMsgColor         = "#ffb703"
	commitStatsColor       = "#a2d2ff"
	headRefColor           = "#d3869b"
	revChoiceColor         = "#8ec07c"
	helpViewTitleColor     = "#83a598"
	helpHeaderColor        = "#83a598"
	helpSectionColor       = "#fabd2f"
)

var (
	baseStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color(defaultBackgroundColor))

	modeStyle = baseStyle.Copy().
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color(modeColor))

	helpMsgStyle = baseStyle.Copy().
			Bold(true).
			Foreground(lipgloss.Color(helpMsgColor))

	viewPortStyle = lipgloss.NewStyle().
			PaddingTop(1).
			PaddingRight(2).
			PaddingLeft(1).
			PaddingBottom(1)

	commitDetailsStyle = lipgloss.NewStyle().
				PaddingLeft(2)

	commitStatsTitleStyle = baseStyle.Copy().
				Bold(true).
				Background(lipgloss.Color(commitStatsTitleColor)).
				Align(lipgloss.Left)

	authorColors = []string{
		"#00a5cf",
		"#00bbf9",
		"#00f5d4",
		"#25a18e",
		"#84bcda",
		"#9f86c0",
		"#d56062",
		"#df7373",
		"#ecc30b",
		"#f37748",
		"#f694c1",
		"#fee440",
		"#a89984",
		"#ffafcc",
		"#ff8fab",
		"#ffd166",
		"#a7c957",
		"#c8b6ff",
	}
	authorStyle = func(author string) lipgloss.Style {
		h := fnv.New32()
		h.Write([]byte(author))
		hash := h.Sum32()

		color := authorColors[int(hash)%len(authorColors)]

		st := lipgloss.NewStyle().
			Foreground(lipgloss.Color(color))

		return st
	}

	hashStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(hashColor))

	dateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dateColor))

	commitStatsStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(commitStatsColor))

	headRefStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(headRefColor))

	revChoiceStyle = lipgloss.NewStyle().
			PaddingRight(1).
			Foreground(lipgloss.Color(revChoiceColor))

	helpVPTitleStyle = baseStyle.Copy().
				Bold(true).
				Background(lipgloss.Color(helpViewTitleColor)).
				Align(lipgloss.Left)

	helpHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(helpHeaderColor))

	helpSectionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(helpSectionColor))
)
