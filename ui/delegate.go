package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type commitDelegate struct {
	base     list.DefaultDelegate
	revStart *string
	revEnd   *string
}

func newCommitDelegate(color lipgloss.Color) commitDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.
		SelectedTitle.
		Foreground(color).
		BorderLeftForeground(color)
	d.Styles.SelectedDesc = d.Styles.
		SelectedTitle

	return commitDelegate{
		base: d,
	}
}

func (d commitDelegate) Height() int {
	return d.base.Height()
}

func (d commitDelegate) Spacing() int {
	return d.base.Spacing()
}

func (d commitDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return d.base.Update(msg, m)
}

func (d commitDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	commit, ok := item.(Commit)
	if !ok {
		d.base.Render(w, m, index, item)
		return
	}

	hash := commit.FilterValue()

	var titleStyle, descStyle lipgloss.Style

	isSelected := index == m.Index()
	isStart := d.revStart != nil && hash == *d.revStart
	isEnd := d.revEnd != nil && hash == *d.revEnd

	if isSelected {
		titleStyle = d.base.Styles.SelectedTitle
		descStyle = d.base.Styles.SelectedDesc
	} else {
		titleStyle = d.base.Styles.NormalTitle
		descStyle = d.base.Styles.NormalDesc
	}

	if isStart && isEnd {
		titleStyle = titleStyle.Foreground(lipgloss.Color("6")).Bold(true)
	} else if isStart {
		titleStyle = titleStyle.Foreground(lipgloss.Color("2")).Bold(true)
	} else if isEnd {
		titleStyle = titleStyle.Foreground(lipgloss.Color("5")).Bold(true)
	}

	title := titleStyle.Render(commit.Title())
	desc := descStyle.Render(commit.Description())

	_, _ = fmt.Fprintf(w, "%s\n%s", title, desc)
}

func defaultDelegate(color lipgloss.Color) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.
		SelectedTitle.
		Foreground(color).
		BorderLeftForeground(color)
	d.Styles.SelectedDesc = d.Styles.
		SelectedTitle

	return d
}
