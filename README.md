<p align="center">
  <h1 align="center">commits</h1>
  <p align="center">
    <a href="https://github.com/dhth/commits/actions/workflows/main.yml"><img alt="build status" src="https://img.shields.io/github/actions/workflow/status/dhth/commits/main.yml?style=flat-square"></a>
    <a href="https://github.com/dhth/commits/actions/workflows/vulncheck.yml"><img alt="vuln check" src="https://img.shields.io/github/actions/workflow/status/dhth/commits/vulncheck.yml?style=flat-square&label=vulncheck"></a>
  </p>
</p>

`commits` lets you glance at git commits through a simple TUI.

<p align="center">
  <img src="https://tools.dhruvs.space/images/commits/commits.gif" alt="Usage" />
</p>

🛠️ Pre-requisites
---

- `git` (only if you want to use `commits` to see diffs)

💾 Installation
---

**go**:

```sh
go install github.com/dhth/commits@latest
```

⚡️ Usage
---

`commits` can receive its configuration via command line flags, and/or a TOML
config file. The default location for this config file is OS-specific:
`$XDG_CONFIG_HOME/commits/commits.toml` on Linux, `~/Library/Application
Support/commits/commits.toml` on macOS.

```toml
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
```

```bash
commits -ignore-pattern='^\[regex\]'
commits -config-file-path='/path/to/config/file.toml'
```

Reference Manual
---

```
commits Reference Manual

commits has 4 views:
    - Commit List View
    - Commit Details View
    - Branch List View
    - Help View

Keyboard Shortcuts

   General

       <tab>                           Switch focus between Commit List View and Commit Details View
       <ctrl+e>                        Open commit/revision range
       <ctrl+x>                        Clear revision range selection
       <ctrl+b>                        Change branch
       ?                               Show help view

   Commit List View

       <enter>/<space>                 Show commit/revision range
       <ctrl+t>                        Choose revision range start/end
       <ctrl+p>                        Show git log

   Commit Details View

       <enter>/<space>                 Show commit/revision range
       h/[                             Go to previous commit
       l/]                             Go to next commit

   Branch List View

       <enter>                         Pick branch
       /                               Start filtering

```

Screenshots
---

![Screen 1](https://tools.dhruvs.space/images/commits/commits-1.png)

![Screen 2](https://tools.dhruvs.space/images/commits/commits-2.png)

Acknowledgements
---

`commits` is built using [bubbletea][1].

[1]: https://github.com/charmbracelet/bubbletea
