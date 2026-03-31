# User Guide

## CLI Commands
- `vision`: start the TUI.
- `vision run [command]`: enters the TUI flow intended for command monitoring.
- `vision serve`: placeholder command for WebSocket server functionality.

## Interaction Model
VISI-ON is modal.

### Normal Mode
- `1`: Dashboard
- `2`: Launcher
- `3`: History
- `4`: Plugins
- `5` or `?`: Help
- `:`: enter Command mode
- `q`: quit
- `Esc`: return to Normal mode from anywhere

### Insert Mode
- Entered with `i` when Launcher tab is active.
- Intended for typed input workflows (currently scaffolded).

### Command Mode
- Start with `:` from Normal mode.
- Supported commands:
  - `help`
  - `plugins`
  - `dashboard`
  - `launcher`
  - `history`
  - `q` or `quit`
- `Tab` autocomplete is supported for known commands.
- `Enter` executes command.

## Tabs
- Dashboard: current pipeline and build status visuals.
- Launcher: build profile launcher view (UI-first, input path scaffolded).
- History: historical run space (currently placeholder content).
- Plugins: plugin management surface (currently placeholder content).
- Help: usage hints and key references.

## Notifications
Transient toast notifications can display success, error, warning, and info states.

## Tips
- Resize your terminal for best panel layout.
- Use command mode for quick tab jumps.
- Keep workflows in Normal mode and use `Esc` as a universal reset.
