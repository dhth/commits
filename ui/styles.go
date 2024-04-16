package ui

import (
	"hash/fnv"

	"github.com/charmbracelet/lipgloss"
)

const (
	DefaultBackgroundColor = "#282828"
	CommitsListColor       = "#fe8019"
	ModeColor              = "#b8bb26"
	HelpMsgColor           = "#83a598"
	CommitStatsTitleColor  = "#83a598"
	HashColor              = "#d3869b"
	DateColor              = "#928374"
	CommitMsgColor         = "#ffb703"
	CommitStatsColor       = "#a2d2ff"
	headRefColor           = "#d3869b"
)

var (
	baseStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color(DefaultBackgroundColor))

	modeStyle = baseStyle.Copy().
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color(ModeColor))

	helpMsgStyle = baseStyle.Copy().
			Bold(true).
			Foreground(lipgloss.Color(HelpMsgColor))

	viewPortStyle = lipgloss.NewStyle().
			PaddingTop(1).
			PaddingRight(2).
			PaddingLeft(1).
			PaddingBottom(1)

	commitDetailsStyle = lipgloss.NewStyle().
				PaddingLeft(2)

	commitStatsTitleStyle = baseStyle.Copy().
				Bold(true).
				Background(lipgloss.Color(CommitStatsTitleColor)).
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
			Foreground(lipgloss.Color(HashColor))

	dateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(DateColor))

	commitStatsStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(CommitStatsColor))

	headRefStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(headRefColor))
)
