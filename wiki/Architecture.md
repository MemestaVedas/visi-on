# Architecture

## Stack
- Language: Go 1.22+
- CLI: Cobra
- TUI: Bubble Tea
- Styling: Lip Gloss

## Entry Point
- `cmd/vision/main.go`
  - Defines root command and subcommands.
  - Launches the Bubble Tea app model.

## Core Domain
- `internal/core/build.go`
  - Build state machine and build model.
  - Build tool types, logs, parsed entries, and errors.
- `internal/core/manager.go`
  - Concurrent build manager with add/get/remove/list/active helpers.

## TUI Layer
- `internal/tui/app.go`
  - App model, mode transitions, key handling, and main render composition.
- `internal/tui/screens/*`
  - Screen-level rendering for Dashboard and placeholder tabs.
- `internal/tui/statusbar.go`, `tabs.go`, `hints.go`, `styles.go`
  - Shared UI components and style system.
- `internal/tui/animations.go`, `transitions.go`
  - Frame-driven animation and transition helpers.
- `internal/tui/notifications.go`
  - Toast notification queue and rendering.

## Platform Abstraction
- `internal/platform/platform.go`
  - Interface for process scan/watch, directory watch, notifications, telemetry.
- `internal/platform/linux/proc.go`
  - Partial Linux implementation; process scanning and build-tool detection exist.
- `internal/platform/windows/toolhelp.go`
  - Build-tool detection exists; several runtime hooks are scaffolded.

## Scaffolded Areas
These directories exist but are intentionally minimal right now:
- `internal/ipc`
- `internal/parser`
- `internal/config`
- `internal/plugin/builtin`

## Data Flow (Current)
1. CLI starts app model.
2. TUI receives keyboard and window-size events.
3. App model updates mode/tab/command buffer.
4. Animation ticks advance frame state.
5. View composition renders workspace + notifications + status + hints.
