package ui

import "fmt"

var (
	helpText = fmt.Sprintf(`
  %s
%s
  %s

  %s
%s
  %s
%s
  %s
%s
`,
		helpHeaderStyle.Render("commits Reference Manual"),
		helpSectionStyle.Render(`
  commits has 3 views:
      - Commit List View
      - Commit Details View
      - Help View (this one)
`),
		helpHeaderStyle.Render("Keyboard Shortcuts"),
		helpHeaderStyle.Render("General"),
		helpSectionStyle.Render(`
      <tab>                           Switch focus between Commit List View and Commit Details View
      <enter>                         Show commit/revision range
      <ctrl+d>                        Open commit/revision range in your text editor (depends
                                      on editor_command in your config file)
      <ctrl+x>                        Clear revision range selection
      ?                               Show help view
`),
		helpHeaderStyle.Render("Commit List View"),
		helpSectionStyle.Render(`
      <ctrl+t>                        Choose revision range start/end
      <ctrl+p>                        Show git log
`),
		helpHeaderStyle.Render("Commit Details View"),
		helpSectionStyle.Render(`
      h/[                             Go to previous commit
      l/]                             Go to next commit
`),
	)
)
