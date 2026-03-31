# Configuration

Default configuration is defined in `configs/default.toml`.

## Server
- `server.ws_port`: WebSocket port (default `7712`).
- `server.udp_port`: UDP port (default `7713`).
- `server.hostname`: bind host (empty string uses default behavior).

## UI
- `ui.theme`: UI theme identifier.
- `ui.log_max_lines`: retained log lines.
- `ui.history_max`: history capacity.

## Notifications
- `notifications.desktop`: desktop notifications toggle.
- `notifications.sound`: sound notifications toggle.

## Editor
- `editor.command`: external editor command.

## Plugin Sections
Default template includes:
- `plugins.slack-notify`
- `plugins.discord-notify`
- `plugins.git-detect`
- `plugins.sound-alert`

These keys are available as config schema today. Runtime plugin wiring is still in progress.
