package ui

import (
	"github.com/go-git/go-git/v5"
)

type Config struct {
	IgnorePattern   string
	OpenInEditorCmd []string
	Repo            *git.Repository
}
