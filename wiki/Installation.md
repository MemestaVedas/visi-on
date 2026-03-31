# Installation

## Prerequisites
- Go 1.22+
- Git (for cloning and updates)

## Windows (Recommended)
From the project root:

```powershell
powershell -ExecutionPolicy Bypass -File .\install_windows.ps1
```

What the script does:
1. Stops a running `vision` process if present.
2. Runs `go mod tidy`.
3. Builds `vision.exe` from `./cmd/vision`.
4. Installs to:
   - `C:\Users\<you>\AppData\Local\Programs\vision\vision.exe`
5. Adds that directory to User `PATH` if missing.

Then run:

```powershell
vision
```

If `vision` is not found in an already-open terminal, open a new terminal session.

## Linux
From the project root:

```bash
chmod +x install_linux.sh
./install_linux.sh
```

What the script does:
1. Runs `go mod tidy`.
2. Builds Linux binary `vision`.
3. Moves binary to `/usr/local/bin/vision`.

Then run:

```bash
vision
```

## Verify Installation
```powershell
vision --help
```

You should see commands including `run` and `serve`.

## Update / Reinstall
Run the same install script again. It rebuilds and replaces the installed binary.
