package ui

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Commit struct {
	Commit *object.Commit
}

type branchItem struct {
	branch *plumbing.Reference
}

type repoInfo struct {
	remoteURLs []string
}

func (commit Commit) Title() string {
	return strings.Split(commit.Commit.Message, "\n")[0]
}

func (c Commit) Description() string {
	return fmt.Sprintf("%s %s", authorStyle(c.Commit.Author.Name).Render(RightPadTrim(c.Commit.Author.Name, 80)), dateStyle.Render(RightPadTrim(humanize.Time(c.Commit.Author.When), 15)))
}

func (commit Commit) FilterValue() string {
	return commit.Commit.Hash.String()
}

func (c Commit) renderStats() string {
	var details string
	details += "\n"
	details += fmt.Sprintf("commit: %s", hashStyle.Render(c.Commit.Hash.String()))
	details += "\n"
	details += fmt.Sprintf("Author: %s <%s>", authorStyle(c.Commit.Author.Name).Render(c.Commit.Author.Name), authorStyle(c.Commit.Author.Name).Render(c.Commit.Author.Email))
	details += "\n"
	details += fmt.Sprintf("Date: %s (%s)", dateStyle.Render(c.Commit.Author.When.String()), dateStyle.Render(humanize.Time(c.Commit.Author.When)))
	details += "\n\n"
	details += c.Commit.Message
	details += "\n\n\n"

	stats, err := c.Commit.Stats()
	if err != nil {
		return fmt.Sprintf("Couldn't get commit stats: %s", err.Error())
	}
	details += commitStatsStyle.Render(stats.String())
	return details
}

func (b branchItem) Title() string {
	return b.branch.Name().Short()
}

func (b branchItem) Description() string {
	return b.branch.Hash().String()[:8]
}

func (b branchItem) FilterValue() string {
	return b.branch.Name().Short()
}
