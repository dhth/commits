package ui

import (
	"fmt"
	"regexp"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.message = ""

	m.ensureInvariantsHold()

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			switch m.activePane {
			case commitsList:
				if m.revEnd != nil {
					m.revEnd = nil
					m.updateCommitDelegate()
				} else if m.revStart != nil {
					m.revStart = nil
					m.revStartIndex = 0
					m.updateCommitDelegate()
				} else {
					return m, tea.Quit
				}
			case helpView:
				m.activePane = m.lastPane
			case branchList:
				if m.branchList.FilterState() != list.Filtering {
					m.branchList.ResetFilter()
					m.activePane = m.lastPane
				}
			case commitDetails:
				m.commitDetailsVP.GotoTop()
				m.activePane = commitsList
			}
		case "enter", "space":
			switch m.activePane {
			case commitsList, commitDetails:
				if len(m.commitsList.Items()) == 0 {
					m.message = "No commits to show"
					break
				}

				if m.revEnd != nil {
					cmds = append(cmds, showRange(m.config.ShowRangeCmd, *m.revStart, *m.revEnd))
				} else {
					hash := m.commitsList.SelectedItem().FilterValue()
					cmds = append(cmds, showCommit(m.config.ShowCommitCmd, hash))
				}
			case branchList:
				bItem, ok := m.branchList.SelectedItem().(branchItem)
				if ok {
					if bItem.branch.Name().String() != m.currentRef.Name().String() {
						m.revStart = nil
						m.revEnd = nil
						m.revStartIndex = 0
						m.updateCommitDelegate()
						cmds = append(cmds, getCommits(m.repo, bItem.branch))
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
		case "ctrl+e":
			switch m.activePane {
			case commitsList, commitDetails:
				if len(m.commitsList.Items()) == 0 {
					m.message = "No commits to open"
					break
				}

				if m.revEnd != nil {
					if len(m.config.OpenRangeCmd) == 0 {
						m.message = "open_range_command is not configured"
						break
					}

					cmds = append(cmds, openRange(m.config.OpenRangeCmd, *m.revStart, *m.revEnd))
				} else {
					if len(m.config.OpenCommitCmd) == 0 {
						m.message = "open_commit_command is not configured"
						break
					}

					hash := m.commitsList.SelectedItem().FilterValue()
					cmds = append(cmds, openCommit(m.config.OpenCommitCmd, hash))
				}
			}
		case "ctrl+t":
			switch m.activePane {
			case commitsList:
				if len(m.commitsList.Items()) == 0 {
					m.message = "No commits to select"
					break
				}

				if m.revStart == nil {
					hash := m.commitsList.SelectedItem().FilterValue()
					m.revStart = &hash
					m.revStartIndex = m.commitsList.Index()
					m.updateCommitDelegate()
				} else {
					if m.commitsList.Index() >= m.revStartIndex {
						m.message = "End revision cannot be before start revision"
					} else {
						hash := m.commitsList.SelectedItem().FilterValue()
						m.revEnd = &hash
						m.updateCommitDelegate()
					}
				}
			}
		case "ctrl+x":
			switch m.activePane {
			case commitsList, commitDetails:
				m.revStart = nil
				m.revEnd = nil
				m.updateCommitDelegate()
			}
		case "ctrl+r":
			switch m.activePane {
			case commitsList:
				cmds = append(cmds, getCommits(m.repo, m.currentRef))
			}
		case "ctrl+b":
			switch m.activePane {
			case commitsList, commitDetails:
				cmds = append(cmds, getBranches(m.repo))
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
			switch m.activePane {
			case commitsList:
				commit, ok := m.commitsList.SelectedItem().(Commit)
				if ok {
					m.commitDetailsVP.SetContent(commit.renderStats())
					m.activePane = commitDetails
				}
			case commitDetails:
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
			m.commitDetailsVP = viewport.New(viewport.WithWidth(msg.Width), viewport.WithHeight(msg.Height-7))
			m.commitStatsVPReady = true
		} else {
			m.commitDetailsVP.SetWidth(msg.Width)
			m.commitDetailsVP.SetHeight(msg.Height - 7)
		}

		if !m.helpVPReady {
			m.helpVP = viewport.New(viewport.WithWidth(msg.Width), viewport.WithHeight(msg.Height-7))
			m.helpVP.SetContent(helpText)
			m.helpVPReady = true
		} else {
			m.helpVP.SetWidth(msg.Width)
			m.helpVP.SetHeight(msg.Height - 7)
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
			cmds = append(cmds, m.commitsList.SetItems(commits))
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
			cmds = append(cmds, m.branchList.SetItems(branches))
			m.branchList.ResetSelected()
			m.activePane = branchList
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

func (m *model) updateCommitDelegate() {
	m.commitsListDel.revStart = m.revStart
	m.commitsListDel.revEnd = m.revEnd
	m.commitsList.SetDelegate(m.commitsListDel)
}

func (m *model) ensureInvariantsHold() {
	if m.revEnd != nil && m.revStart == nil {
		m.revEnd = nil
		m.revStart = nil
		m.revStartIndex = 0
		m.updateCommitDelegate()
	}
}
