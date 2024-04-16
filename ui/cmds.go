package ui

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func chooseTableEntry(tableName string) tea.Cmd {
	return func() tea.Msg {
		return tableChosenMsg{tableName}
	}
}

func hideHelp(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(time.Time) tea.Msg {
		return hideHelpMsg{}
	})
}

func getRepoInfo(path string) tea.Cmd {
	return func() tea.Msg {
		r, err := git.PlainOpen(path)
		if err != nil {
			return repoInfoFetchedMsg{err: err}
		}

		var remoteURLs []string
		remotes, err := r.Remotes()
		if err == nil {
			for _, remote := range remotes {
				for _, url := range remote.Config().URLs {
					remoteURLs = append(remoteURLs, url)
				}
			}
		}

		info := repoInfo{remoteURLs: remoteURLs}

		return repoInfoFetchedMsg{info: info}
	}
}

func getCommits(path string) tea.Cmd {
	return func() tea.Msg {
		r, err := git.PlainOpen(path)
		if err != nil {
			return commitsFetched{err: fmt.Errorf("Couldn't fetch git repo: %s", err.Error())}
		}
		ref, err := r.Head()
		if err != nil {
			return commitsFetched{err: fmt.Errorf("Couldn't get HEAD: %s", err.Error())}
		}

		since := time.Now().Add(-time.Hour * 24 * 7 * 6)
		cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, All: false})

		if err != nil {
			return commitsFetched{err: err}
		}

		var commits []*object.Commit
		for {
			commit, iterErr := cIter.Next()
			if iterErr != nil {
				break
			}
			commits = append(commits, commit)
		}
		return commitsFetched{commits: commits}
	}
}

func getCurrentRev(path string) tea.Cmd {
	return func() tea.Msg {
		r, err := git.PlainOpen(path)
		if err != nil {
			return currentRevFetchedMsg{err: err}
		}
		ref, err := r.Head()
		if err != nil {
			return currentRevFetchedMsg{err: err}
		}

		return currentRevFetchedMsg{rev: ref.Name().String()}
	}
}

func showRevisionRange(path string, revisionRange string) tea.Cmd {
	c := exec.Command("git", "-C", path, "diff", revisionRange)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		if err != nil {
			return showDiffFinished{err}
		}
		return tea.Msg(showDiffFinished{})
	})
}

func showCommit(path string, hash string) tea.Cmd {
	c := exec.Command("git", "-C", path, "show", hash)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		if err != nil {
			return showDiffFinished{err}
		}
		return tea.Msg(showDiffFinished{})
	})
}

func openURLInBrowser(url string) tea.Cmd {
	return func() tea.Msg {
		var openCmd string
		switch runtime.GOOS {
		case "darwin":
			openCmd = "open"
		default:
			openCmd = "xdg-open"
		}
		c := exec.Command(openCmd, url)
		err := c.Run()

		return urlOpenedinBrowserMsg{url: url, err: err}
	}
}

func (m model) showGitLog() tea.Cmd {

	hashWidth := m.terminalWidth / 10
	messageWidth := m.terminalWidth / 2
	dateWidth := m.terminalWidth / 6
	authorWidth := m.terminalWidth / 6
	prettyFormat := fmt.Sprintf("%%<(%d,trunc)%%Cred%%h%%Creset%%<(%d,trunc)%%s %%Cgreen%%<(%d,trunc)%%cr%%C(bold blue)%%<(%d,trunc)%%an%%Creset", hashWidth, messageWidth, dateWidth, authorWidth)

	c := exec.Command("git",
		"-C",
		m.config.Path,
		"log",
		"--color",
		"--since=\"3 months ago\"",
		fmt.Sprintf("--pretty=format:%s", prettyFormat),
		"--abbrev-commit",
	)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		if err != nil {
			return showDiffFinished{err}
		}
		return tea.Msg(showDiffFinished{})
	})
}

func openRevisionRangeInEditor(command []string, revisionRange string) tea.Cmd {
	cmdRep := make([]string, 0)
	for _, word := range command {
		if strings.Contains(word, "{{revision}}") {
			cmdRep = append(cmdRep, strings.Replace(word, "{{revision}}", revisionRange, 1))
		} else {
			cmdRep = append(cmdRep, word)
		}
	}

	log.Printf("%#v", cmdRep)
	c := exec.Command(cmdRep[0], cmdRep[1:]...)

	return tea.ExecProcess(c, func(err error) tea.Msg {
		if err != nil {
			return showCommitInEditorFinished{err: err}
		}
		return tea.Msg(showCommitInEditorFinished{hash: revisionRange})
	})
}
