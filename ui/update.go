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
			} else if m.activePane == helpView {
				m.activePane = m.lastPane
			} else if m.activePane == branchList {
				if m.branchList.FilterState() != list.Filtering {
					m.branchList.ResetFilter()
					m.activePane = m.lastPane
				}
			} else if m.activePane == commitDetails {
				m.commitDetailsVP.GotoTop()
				m.activePane = commitsList
			}
		case "enter":
			switch m.activePane {
			case commitsList, commitDetails:
				if m.revEndChosen {
					cmds = append(cmds, showRevisionRange(m.config.Path, fmt.Sprintf("%s..%s", m.revStart, m.revEnd)))
				} else {
					hash := m.commitsList.SelectedItem().FilterValue()
					cmds = append(cmds, showCommit(m.config.Path, hash))
				}
			case branchList:
				bItem, ok := m.branchList.SelectedItem().(branchItem)
				if ok {
					if bItem.branch.Name().String() != m.currentRef.Name().String() {
						cmds = append(cmds, getCommits(m.config.Repo, bItem.branch))
					} else {
						m.activePane = commitsList
					}
				}
				m.branchList.ResetFilter()
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
					if m.revEndChosen {
						cmds = append(cmds, openRevisionRangeInEditor(m.config.OpenInEditorCmd, fmt.Sprintf("%s..%s", m.revStart, m.revEnd)))
					} else {
						hash := m.commitsList.SelectedItem().FilterValue()
						cmds = append(cmds, openRevisionRangeInEditor(m.config.OpenInEditorCmd, fmt.Sprintf("%s~1..%s", hash, hash)))
					}
				}
			} else {
				m.message = "editor_command is not configured"
			}
		case "ctrl+t":
			switch m.activePane {
			case commitsList:
				if !m.revStartChosen {
					m.revStart = m.commitsList.SelectedItem().FilterValue()[:10]
					m.revStartChosen = true
					m.revStartIndex = m.commitsList.Index()
				} else if !m.revEndChosen {
					if m.commitsList.Index() >= m.revStartIndex {
						m.message = "End revision cannot be before start revision"
					} else {
						m.revEnd = m.commitsList.SelectedItem().FilterValue()[:10]
						m.revEndChosen = true
					}
				}
			}
		case "ctrl+x":
			switch m.activePane {
			case commitsList, commitDetails:
				m.revStartChosen = false
				m.revEndChosen = false
			}
		case "ctrl+r":
			switch m.activePane {
			case commitsList:
				cmds = append(cmds, getCommits(m.config.Repo, m.currentRef))
			}
		case "ctrl+b":
			switch m.activePane {
			case commitsList, commitDetails:
				cmds = append(cmds, getBranches(m.config.Repo))
			}
		case "[", "h":
			switch m.activePane {
			case commitDetails:
				m.commitsList.CursorUp()
				commit, ok := m.commitsList.SelectedItem().(Commit)
				if ok {
					m.commitDetailsVP.SetContent(commit.renderStats())
					m.activePane = commitDetails
				}
			}
		case "]", "l":
			switch m.activePane {
			case commitDetails:
				m.commitsList.CursorDown()
				commit, ok := m.commitsList.SelectedItem().(Commit)
				if ok {
					m.commitDetailsVP.SetContent(commit.renderStats())
					m.activePane = commitDetails
				}
			}
		case "tab", "shift+tab":
			if m.activePane == commitsList {
				commit, ok := m.commitsList.SelectedItem().(Commit)
				if ok {
					m.commitDetailsVP.SetContent(commit.renderStats())
					m.activePane = commitDetails
				}
			} else if m.activePane == commitDetails {
				m.activePane = commitsList
			}
		case "?":
			if m.activePane == helpView {
				break
			}
			m.lastPane = m.activePane
			m.activePane = helpView
		}
	case hideHelpMsg:
		m.showHelp = false
	case tea.WindowSizeMsg:
		_, h1 := m.commitListStyle.GetFrameSize()
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.commitsList.SetHeight(msg.Height - h1 - 2)
		m.commitsList.SetWidth(msg.Width)

		m.branchList.SetHeight(msg.Height - h1 - 2)
		m.branchList.SetWidth(msg.Width)

		if !m.commitStatsVPReady {
			m.commitDetailsVP = viewport.New(msg.Width, msg.Height-7)
			m.commitDetailsVP.HighPerformanceRendering = false
			m.commitStatsVPReady = true
		} else {
			m.commitDetailsVP.Width = msg.Width
			m.commitDetailsVP.Height = msg.Height - 7
		}

		if !m.helpVPReady {
			m.helpVP = viewport.New(msg.Width, msg.Height-7)
			m.helpVP.HighPerformanceRendering = false
			m.helpVP.SetContent(helpText)
			m.helpVPReady = true
		} else {
			m.helpVP.Width = msg.Width
			m.helpVP.Height = msg.Height - 7
		}
	case repoInfoFetchedMsg:
		if msg.err == nil {
			m.repoInfo = msg.info
		} else {
			m.message = fmt.Sprintf("%v", msg.info.remoteURLs)
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
			m.commitsList.ResetSelected()

			if msg.ref != nil {
				m.currentRef = msg.ref
			}

			m.activePane = commitsList
		}
	case branchesFetched:
		if msg.err != nil {
			m.message = msg.err.Error()
		} else {
			var branches []list.Item
			for _, branch := range msg.branches {
				branches = append(branches, branchItem{
					branch: branch,
				})
			}
			m.branchList.SetItems(branches)
			m.branchList.ResetSelected()
			m.activePane = branchList
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
		m.commitDetailsVP, cmd = m.commitDetailsVP.Update(msg)
		cmds = append(cmds, cmd)
	case branchList:
		m.branchList, cmd = m.branchList.Update(msg)
		cmds = append(cmds, cmd)
	case helpView:
		m.helpVP, cmd = m.helpVP.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
