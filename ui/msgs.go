package ui

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type hideHelpMsg struct{}

type repoInfoFetchedMsg struct {
	info repoInfo
	err  error
}

type commitsFetched struct {
	commits []*object.Commit
	ref     *plumbing.Reference
	err     error
}

type branchesFetched struct {
	branches []*plumbing.Reference
	err      error
}

type showDiffFinished struct {
	err error
}

type showCommitInEditorFinished struct {
	hash string
	err  error
}
