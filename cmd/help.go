package cmd

import "fmt"

var (
	configSampleFormat = `
# commit messages that match "ignore_pattern" will not be shown in the TUI list
ignore_pattern = '^\[regex\]'

# show_commit_command is run when you press enter/space on a single commit;
# {{hash}} is replaced at runtime with the commit hash
# (defaults to "git show {{hash}}" if not set)
show_commit_command = [ "git", "show", "{{hash}}" ]

# show_range_command is run when you press enter/space on a revision range;
# {{base}} and {{head}} are replaced at runtime
# (defaults to "git diff {{base}}..{{head}}" if not set)
show_range_command = [ "git", "diff", "{{base}}..{{head}}" ]

# open_commit_command is run when you press ctrl+e on a single commit;
# {{hash}} is replaced at runtime with the commit hash
open_commit_command = [ "nvim", "-c", ":DiffviewOpen {{hash}}~1..{{hash}}" ]

# open_range_command is run when you press ctrl+e on a revision range;
# {{base}} and {{head}} are replaced at runtime
open_range_command = [ "nvim", "-c", ":DiffviewOpen {{base}}..{{head}}" ]
`
	helpText = `commits lets you glance at git commits through a simple TUI

Usage: commits [flags]
`
)

func cfgErrSuggestion(msg string) string {
	return fmt.Sprintf(`%s

Make sure to structure the toml config file as follows:

------
%s
------

Use "commits --help" for more information`,
		msg,
		configSampleFormat,
	)
}
