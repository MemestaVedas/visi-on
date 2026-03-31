# VISI-ON Wiki

## Product Summary
VISI-ON is a build monitoring experience for developers who live in the terminal. It combines a keyboard-first interface, visual feedback, and a modular architecture for future integrations.

The product currently delivers:
- A modern Bubble Tea + Lip Gloss TUI.
- Multi-mode navigation (Normal, Insert, Command).
- A tabbed workspace with focused views.
- Styled status, hints, and toast notifications.
- Core build domain models and manager logic.
- Cross-platform install scripts for Windows and Linux.

## Who This Is For
- Individual developers who want fast build visibility.
- Teams building tooling around command-line workflows.
- Contributors who want to extend monitoring, parser, or plugin capabilities.

## Core Concepts
- Mode system:
  - Normal mode for navigation.
  - Insert mode for input workflows (Launcher-targeted path).
  - Command mode (`:`) for quick commands.
- Tabs:
  - Dashboard
  - Launcher
  - History
  - Plugins
  - Help
- Notifications:
  - Success, error, warning, and info toasts.
- Animation:
  - Tick-driven animation state for subtle, live terminal feedback.

## Start Here
1. [Installation](Installation.md)
2. [User Guide](User-Guide.md)
3. [Configuration](Configuration.md)
4. [Architecture](Architecture.md)
5. [Development](Development.md)

## Current Scope and Gaps
VISI-ON has strong UI foundations and domain structure. Several backend integrations are intentionally scaffolded and documented as next implementation targets.
