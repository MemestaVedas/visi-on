//go:build windows

package windows

import (
    "os"
    "path/filepath"
    "strings"

    "github.com/visi-on/visi-on/internal/platform"
)

type WindowsPlatform struct{}

func New() *WindowsPlatform { return &WindowsPlatform{} }

func (w *WindowsPlatform) Name() string { return "windows" }

// ScanBuildProcesses — Phase 2: ToolHelp32 + NtQueryInformationProcess
func (w *WindowsPlatform) ScanBuildProcesses() ([]platform.ProcessInfo, error) { return nil, nil }

func (w *WindowsPlatform) WatchProcess(pid int, onChange func(platform.ProcessInfo)) error { return nil }
func (w *WindowsPlatform) WatchDirectory(path string, onChange func(platform.FileEvent)) error { return nil }

func (w *WindowsPlatform) DetectBuildTool(dir string) (platform.BuildTool, error) {
    checks := []struct{ file string; tool platform.BuildTool }{
        {"Cargo.toml", platform.BuildToolCargo},
        {"package.json", platform.BuildToolNPM},
        {"Makefile", platform.BuildToolMake},
        {"go.mod", platform.BuildToolGo},
        {"build.gradle", platform.BuildToolGradle},
        {"CMakeLists.txt", platform.BuildToolCMake},
    }
    for _, c := range checks {
        if _, err := os.Stat(filepath.Join(dir, c.file)); err == nil {
            return c.tool, nil
        }
    }
    // Also check for .sln (MSBuild)
    if matches, _ := filepath.Glob(filepath.Join(dir, "*.sln")); len(matches) > 0 {
        return platform.BuildToolMSBuild, nil
    }
    return platform.BuildToolUnknown, nil
}

// SendNotification — Phase 2: WinRT ToastNotificationManager
func (w *WindowsPlatform) SendNotification(title, body string) error { return nil }
func (w *WindowsPlatform) PlaySound(t platform.SoundType) error      { return nil }
func (w *WindowsPlatform) GetCPUPercent() (float64, error)           { return 0, nil }
func (w *WindowsPlatform) GetRAMUsage() (uint64, uint64, error)      { return 0, 0, nil }
func (w *WindowsPlatform) GetNetworkIO() (uint64, uint64, error)     { return 0, 0, nil }

func isBuildCommand(c string) bool {
    c = strings.ToLower(c)
    for _, p := range []string{"cargo", "npm", "make", "go build", "msbuild", "dotnet build", "gradle"} {
        if strings.Contains(c, p) {
            return true
        }
    }
    return false
}
