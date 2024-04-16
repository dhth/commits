# commits

✨ Overview
---

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
config file. The default location for this config file is
`~/.config/commits/commits.toml`.

```toml
# commit messages that match "ignore_pattern" will not be shown in the TUI list
ignore_pattern = '^\[regex\]'

# editor_command is run when you press ctrl+d; {{revision}} is replaced at
# runtime with a revision range
editor_command = [ "nvim", "-c", ":DiffviewOpen {{revision}}" ]
```

```bash
commits -path='/path/to/git/repo'
commits -ignore_pattern='^\[regex\]'
commits -config-file-path='/path/to/config/file.toml'
```

Screenshots
---

![Screen 1](https://tools.dhruvs.space/images/commits/commits-1.png)

![Screen 2](https://tools.dhruvs.space/images/commits/commits-2.png)

Acknowledgements
---

`commits` is built using [bubbletea][1].

[1]: https://github.com/charmbracelet/bubbletea
