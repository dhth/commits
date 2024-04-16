package ui

import "github.com/go-git/go-git/v5/plumbing/object"

type hideHelpMsg struct{}

type tableChosenMsg struct {
	tableName string
}

type repoInfoFetchedMsg struct {
	info repoInfo
	err  error
}

type commitsFetched struct {
	commits []*object.Commit
	err     error
}

type showDiffFinished struct {
	err error
}

type currentRevFetchedMsg struct {
	rev string
	err error
}

type urlOpenedinBrowserMsg struct {
	url string
	err error
}
