# Development

## Local Build
From project root:

```powershell
go build -o vision.exe ./cmd/vision
```

Linux equivalent:

```bash
go build -o vision ./cmd/vision
```

## Run in Development
```powershell
go run ./cmd/vision
```

## Format and Verify
```powershell
gofmt -w ./cmd ./internal
go test ./...
```

## Installation Scripts
- Windows: `install_windows.ps1`
- Linux: `install_linux.sh`
- Convenience launcher (Windows): `build_install.bat`

## Contributing Checklist
1. Keep behavior and style changes isolated.
2. Run `gofmt` on changed files.
3. Run `go test ./...` before commit.
4. Document user-visible behavior in the wiki.

## Near-Term Implementation Targets
- Parser and IPC integration.
- Platform telemetry completion on Windows/Linux.
- Launcher input flow and profile execution pipeline.
- Plugin runtime and configuration loading.
- History persistence and searchable log views.
