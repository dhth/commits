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
  %s
%s
`,
		helpHeaderStyle.Render("commits Reference Manual"),
		helpSectionStyle.Render(`
  (scroll line by line with j/k/arrow keys or by half a page with <c-d>/<c-u>)

  commits has 4 views:
      - Commit List View
      - Commit Details View
      - Branch List View
      - Help View (this one)
`),
		helpHeaderStyle.Render("Keyboard Shortcuts"),
		helpHeaderStyle.Render("General"),
		helpSectionStyle.Render(`
      <tab>                           Switch focus between Commit List View and Commit Details View
      <ctrl+d>                        Open commit/revision range in your text editor (depends
                                      on editor_command in your config file)
      <ctrl+x>                        Clear revision range selection
      <ctrl+b>                        Change branch
      ?                               Show help view
`),
		helpHeaderStyle.Render("Commit List View"),
		helpSectionStyle.Render(`
      <enter>                         Show commit/revision range
      <ctrl+t>                        Choose revision range start/end
      <ctrl+p>                        Show git log
`),
		helpHeaderStyle.Render("Commit Details View"),
		helpSectionStyle.Render(`
      <enter>                         Show commit/revision range
      h/[                             Go to previous commit
      l/]                             Go to next commit
`),
		helpHeaderStyle.Render("Branch List View"),
		helpSectionStyle.Render(`
      <enter>                         Pick branch
      /                               Start filtering
`),
	)
)
