#!/usr/bin/env bash
set -euo pipefail
echo "→ Building VISI-ON for Linux..."
go mod tidy
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o vision ./cmd/vision
sudo mv vision /usr/local/bin/vision
echo "✓ Done. Type: vision"
