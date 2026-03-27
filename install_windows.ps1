Write-Host "Building VISI-ON for Windows..."

# Clean up any potential locked files by killing the process
Stop-Process -Name "vision" -Force -ErrorAction SilentlyContinue

# Ensure dependencies
go mod tidy

# Build
$env:GOOS="windows"; $env:GOARCH="amd64"; $env:CGO_ENABLED="0"
Write-Host "Compiling executable..."
go build -ldflags="-s -w" -o vision.exe ./cmd/vision

if (-not (Test-Path "vision.exe")) {
    Write-Error "Build failed. vision.exe not found."
    exit 1
}

$dir = "$env:LOCALAPPDATA\Programs\vision"
if (-not (Test-Path $dir)) {
    New-Item -ItemType Directory -Force -Path $dir | Out-Null
}

$target = "$dir\vision.exe"
if (Test-Path $target) {
    Remove-Item -Force $target -ErrorAction SilentlyContinue
}

Write-Host "Installing to $dir..."
Move-Item -Force vision.exe $target

$p = [Environment]::GetEnvironmentVariable("PATH","User")
if ($p -notlike "*$dir*") {
    $newPath = "$p;$dir"
    [Environment]::SetEnvironmentVariable("PATH", $newPath, "User")
    Write-Host "Added $dir to your User PATH!"
    Write-Host "IMPORTANT: Please restart your terminal for the PATH changes to take effect."
}

Write-Host "Done! You can now launch the app by typing: vision"
