package cmd

import "fmt"

var (
	configSampleFormat = `
# commit messages that match "ignore_pattern" will not be shown in the TUI list
ignore_pattern = '^\[regex\]'

# editor_command is run when you press ctrl+d; {{revision}} is replaced at
# runtime with a revision range
editor_command = [ "nvim", "-c", ":DiffviewOpen {{revision}}" ]
`
	helpText = `commits lets you glance at git commits through a simple TUI.

Keyboard shortcuts:
- tab:    commit details
- enter:  show commit/revision range
- ctrl+d: open in editor
- ctrl+t: pick range
- ctrl+x: clear range
- ctrl+p: show log

Usage: commits [flags]
`
)

func cfgErrSuggestion(msg string) string {
	return fmt.Sprintf(`%s

Make sure to structure the toml config file as follows:

------
%s
------

Use "commits -help" for more information`,
		msg,
		configSampleFormat,
	)
}
