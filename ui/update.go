package ui

import (
	"fmt"
	"regexp"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.message = ""

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.activePane == commitsList {
				return m, tea.Quit
			} else {
				m.activePane = commitsList
			}
		case "enter":
			switch m.activePane {
			case commitsList, commitDetails:
				hash := m.commitsList.SelectedItem().FilterValue()
				cmds = append(cmds, showCommit(m.config.Path, hash))
			}
		case "ctrl+p":
			switch m.activePane {
			case commitsList:
				cmds = append(cmds, m.showGitLog())
			}
		case "ctrl+d":
			if m.config.OpenInEditorCmd != nil {
				switch m.activePane {
				case commitsList, commitDetails:
					hash := m.commitsList.SelectedItem().FilterValue()
					cmds = append(cmds, openCommitInEditor(m.config.OpenInEditorCmd, hash))
				}
			}
		case "[", "h":
			switch m.activePane {
			case commitDetails:
				m.commitsList.CursorUp()
				commit, ok := m.commitsList.SelectedItem().(Commit)
				if ok {
					m.commitStatsVP.SetContent(commit.renderStats())
					m.activePane = commitDetails
				}
			}
		case "]", "l":
			switch m.activePane {
			case commitDetails:
				m.commitsList.CursorDown()
				commit, ok := m.commitsList.SelectedItem().(Commit)
				if ok {
					m.commitStatsVP.SetContent(commit.renderStats())
					m.activePane = commitDetails
				}
			}
		case "tab", "shift+tab":
			if m.activePane == commitsList {
				commit, ok := m.commitsList.SelectedItem().(Commit)
				if ok {
					m.commitStatsVP.SetContent(commit.renderStats())
					m.activePane = commitDetails
				}
			} else {
				m.activePane = commitsList
			}
		}
	case hideHelpMsg:
		m.showHelp = false
	case tea.WindowSizeMsg:
		_, h1 := m.commitListStyle.GetFrameSize()
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.commitsList.SetHeight(msg.Height - h1 - 2)
		m.commitsList.SetWidth(msg.Width)

		if !m.commitStatsVPReady {
			m.commitStatsVP = viewport.New(msg.Width, msg.Height-7)
			m.commitStatsVP.HighPerformanceRendering = false
			m.commitStatsVPReady = true
		} else {
			m.commitStatsVP.Width = msg.Width
			m.commitStatsVP.Height = msg.Height - 7
		}
	case repoInfoFetchedMsg:
		if msg.err == nil {
			m.repoInfo = msg.info
		} else {
			m.message = fmt.Sprintf("%v", msg.info.remoteURLs)
		}
	case currentRevFetchedMsg:
		if msg.err != nil {
			m.message = msg.err.Error()
		} else {
			m.currentRev = msg.rev
		}
	case commitsFetched:
		if msg.err != nil {
			m.message = msg.err.Error()
		} else {
			var commits []list.Item
			if m.config.IgnorePattern != "" {
				re, regexErr := regexp.Compile(m.config.IgnorePattern)

				for _, commit := range msg.commits {
					var matched bool
					if regexErr == nil {
						matched = re.MatchString(commit.Message)
					}
					if !matched {
						commits = append(commits, Commit{
							Commit: commit,
						})
					}
				}
			} else {
				for _, commit := range msg.commits {
					commits = append(commits, Commit{
						Commit: commit,
					})
				}

			}
			m.commitsList.SetItems(commits)
		}
	case urlOpenedinBrowserMsg:
		if msg.err == nil {
			m.message = msg.err.Error()
		}
	}

	switch m.activePane {
	case commitsList:
		m.commitsList, cmd = m.commitsList.Update(msg)
		cmds = append(cmds, cmd)
	case commitDetails:
		m.commitStatsVP, cmd = m.commitStatsVP.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
