# CLAUDE.md

This file provides guidance to AI agents when working with code in this repository.

## Project Overview

`commits` is a terminal user interface (TUI) application that lets users browse
git commits interactively. Built in Go using the Bubble Tea framework
(Charmbracelet) for TUI components and go-git for Git operations.

## Common Commands

```bash
# Build and run
just run              # go run .
just build            # go build -ldflags "-w -s" .
just install          # go install -ldflags "-w -s" .

# Quality
just lint             # golangci-lint run
just fmt              # gofumpt -l -w .
just vuln             # govulncheck ./...
just all              # fmt + lint

# Dependencies
just tidy             # go mod tidy
just upgrade          # go get -u ./...
```

Always run go commands via `just`.

## Debugging

Set `DEBUG=1` to enable Bubble Tea debug logging to `debug.log` (see `ui/ui.go`).

## Architecture

The application follows the Bubble Tea pattern (Model-Update-View) with a pane-based navigation system.

### Pane System

The UI has 4 panes (`commitsList`, `commitDetails`, `branchList`, `helpView`), tracked by `model.activePane`. Only one pane is active at a time. Keyboard input is routed first through global handlers (quit, navigation) in `update.go`, then delegated to the active pane's bubbles component.

### Data Flow

1. **Startup**: `cmd/root.go` parses flags + TOML config (`~/.config/commits/commits.toml`) into `ui.Config`, opens the git repo with go-git, calls `ui.RenderUI`
2. **Init**: `model.Init()` fires `getCommits` (async tea.Cmd) which reads commits from the last ~12 weeks via go-git
3. **User interaction**: Key events in `update.go` trigger tea.Cmds defined in `cmds.go`; external commands (`git diff`, `git show`, `git log`, editor) use `tea.ExecProcess` to temporarily yield the terminal
4. **Revision ranges**: Users can select start/end commits (`ctrl+t`) tracked by `revStart`/`revEnd` pointers on the model; the `commitDelegate` in `delegate.go` highlights selected commits with different colors

### Key Files

- `ui/model.go` - Central state: panes, viewport readiness, revision range selection
- `ui/update.go` - All keyboard handling and message dispatch
- `ui/cmds.go` - Async tea.Cmds: git operations (go-git), shell-outs (`git diff/show/log`), editor launch
- `ui/delegate.go` - Custom list delegate that renders commit items with revision range highlighting
- `ui/initial.go` - Model construction, bubbles component setup
- `ui/view.go` - Renders the active pane + footer (branch, help hint, revision range)
- `cmd/root.go` - CLI entrypoint: flag parsing, config loading, repo opening
- `cmd/config.go` - TOML config struct and parser

## Key Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/bubbles` - Pre-built TUI components (lists, viewports)
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/go-git/go-git/v5` - Pure Go git implementation
- `github.com/BurntSushi/toml` - Config parsing

## Writing Code

- Follow existing conventions set up in the codebase; if an existing convention
    can be improved, let the user know about it
- Only add code comments when they genuinely add more clarity
