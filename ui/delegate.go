package ui

import (
	"fmt"
	"image/color"
	"io"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type commitDelegate struct {
	base     list.DefaultDelegate
	revStart *string
	revEnd   *string
}

func newCommitDelegate(color color.Color) commitDelegate {
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

	if isStart {
		titleStyle = titleStyle.Foreground(lipgloss.Color(revisionBaseColor)).Bold(true)
	} else if isEnd {
		titleStyle = titleStyle.Foreground(lipgloss.Color(revisionHeadColor)).Bold(true)
	}

	title := titleStyle.Render(commit.Title())
	desc := descStyle.Render(commit.Description())

	_, _ = fmt.Fprintf(w, "%s\n%s", title, desc)
}

func defaultDelegate(color color.Color) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.Styles.SelectedTitle = d.Styles.
		SelectedTitle.
		Foreground(color).
		BorderLeftForeground(color)
	d.Styles.SelectedDesc = d.Styles.
		SelectedTitle

	return d
}
