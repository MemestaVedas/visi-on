//go:build linux

package linux

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strconv"
    "strings"

    "github.com/visi-on/visi-on/internal/platform"
)

type LinuxPlatform struct{}

func New() *LinuxPlatform { return &LinuxPlatform{} }

func (l *LinuxPlatform) Name() string { return "linux" }

func (l *LinuxPlatform) ScanBuildProcesses() ([]platform.ProcessInfo, error) {
    entries, err := os.ReadDir("/proc")
    if err != nil {
        return nil, fmt.Errorf("read /proc: %w", err)
    }
    var results []platform.ProcessInfo
    for _, e := range entries {
        pid, err := strconv.Atoi(e.Name())
        if err != nil {
            continue
        }
        data, err := os.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
        if err != nil {
            continue
        }
        cmdline := strings.ReplaceAll(string(data), "\x00", " ")
        if !isBuildCommand(cmdline) {
            continue
        }
        cwd, _ := os.Readlink(fmt.Sprintf("/proc/%d/cwd", pid))
        results = append(results, platform.ProcessInfo{
            PID:     pid,
            Name:    filepath.Base(strings.Fields(cmdline)[0]),
            CmdLine: strings.TrimSpace(cmdline),
            CWD:     cwd,
            Source:  "native",
        })
    }
    return results, nil
}

func isBuildCommand(c string) bool {
    c = strings.ToLower(c)
    for _, p := range []string{
        "cargo build", "cargo test", "npm run", "npm build",
        "make ", "go build", "go test", "gradle", "cmake --build",
    } {
        if strings.Contains(c, p) {
            return true
        }
    }
    return false
}

func (l *LinuxPlatform) WatchProcess(pid int, onChange func(platform.ProcessInfo)) error { return nil }
func (l *LinuxPlatform) WatchDirectory(path string, onChange func(platform.FileEvent)) error { return nil }

func (l *LinuxPlatform) DetectBuildTool(dir string) (platform.BuildTool, error) {
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
    return platform.BuildToolUnknown, nil
}

func (l *LinuxPlatform) SendNotification(title, body string) error {
    return exec.Command("notify-send", "--app-name=GoBuild", title, body).Run()
}

func (l *LinuxPlatform) PlaySound(t platform.SoundType) error    { return nil }
func (l *LinuxPlatform) GetCPUPercent() (float64, error)         { return 0, nil }
func (l *LinuxPlatform) GetRAMUsage() (uint64, uint64, error)    { return 0, 0, nil }
func (l *LinuxPlatform) GetNetworkIO() (uint64, uint64, error)   { return 0, 0, nil }
