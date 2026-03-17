Write-Host "Building VISI-ON for Windows..."
go mod tidy
$env:GOOS="windows"; $env:GOARCH="amd64"; $env:CGO_ENABLED="0"
go build -ldflags="-s -w" -o vision.exe ./cmd/vision
$dir = "$env:LOCALAPPDATA\Programs\vision"
New-Item -ItemType Directory -Force -Path $dir | Out-Null
Move-Item -Force vision.exe "$dir\vision.exe"
$p = [Environment]::GetEnvironmentVariable("PATH","User")
if ($p -notlike "*$dir*") {
    [Environment]::SetEnvironmentVariable("PATH","$p;$dir","User")
    Write-Host "Added $dir to PATH. Restart terminal."
}
Write-Host "Done. Type: vision"
