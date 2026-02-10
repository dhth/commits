package cmd

import (
	"errors"
	"fmt"
)

var configSampleFormat = `
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

func GetErrorFollowUp(err error) (string, bool) {
	if errors.Is(err, errCouldntParseConfig) {
		return fmt.Sprintf(`
Make sure to structure the TOML config file as follows:

------%s------
`, configSampleFormat), true
	}

	if errors.Is(err, errConfigIsInvalid) {
		return fmt.Sprintf(`
A valid config looks like this:

------%s------
`, configSampleFormat), true
	}

	return "", false
}
