# VISI-ON

VISI-ON is a terminal-first developer build monitor with a modern TUI and cross-platform foundations.

## What It Does
- Runs as a keyboard-driven terminal UI (`vision`).
- Organizes workflow into tabs: Dashboard, Launcher, History, Plugins, and Help.
- Supports command mode (`:`) for quick navigation.
- Includes an animation and notification system for real-time feedback.
- Provides platform abstractions for process scanning, build-tool detection, and notifications.

## Quick Start
### Windows
```powershell
powershell -ExecutionPolicy Bypass -File .\install_windows.ps1
vision
```

### Linux
```bash
./install_linux.sh
vision
```

## Documentation (Wiki)
- [Home](wiki/Home.md)
- [Installation](wiki/Installation.md)
- [User Guide](wiki/User-Guide.md)
- [Configuration](wiki/Configuration.md)
- [Architecture](wiki/Architecture.md)
- [Development](wiki/Development.md)

## Current Status
VISI-ON has a polished TUI shell and core domain models. Some integration layers are scaffolded and ready for implementation (IPC, parser, config loader, plugin runtime, and parts of platform telemetry).
