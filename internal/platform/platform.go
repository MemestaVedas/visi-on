package platform

type SoundType int

const (
    SoundSuccess SoundType = iota
    SoundFailure
    SoundWarning
)

type BuildTool int

const (
    BuildToolUnknown BuildTool = iota
    BuildToolCargo
    BuildToolNPM
    BuildToolMake
    BuildToolGo
    BuildToolMSBuild
    BuildToolGradle
    BuildToolCMake
)

type FileEvent struct {
    Path string
    Op   string
}

type ProcessInfo struct {
    PID      int
    PPID     int
    Name     string
    CmdLine  string
    CWD      string
    CPUPct   float64
    MemBytes uint64
    Source   string
}

type Platform interface {
    ScanBuildProcesses() ([]ProcessInfo, error)
    WatchProcess(pid int, onChange func(ProcessInfo)) error
    WatchDirectory(path string, onChange func(FileEvent)) error
    DetectBuildTool(dir string) (BuildTool, error)
    SendNotification(title, body string) error
    PlaySound(soundType SoundType) error
    GetCPUPercent() (float64, error)
    GetRAMUsage() (uint64, uint64, error)
    GetNetworkIO() (uint64, uint64, error)
    Name() string
}
