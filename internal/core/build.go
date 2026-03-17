package core

import (
	"time"

	"github.com/google/uuid"
)

type BuildState int

const (
	StateIdle BuildState = iota
	StateQueued
	StateBuilding
	StateSuccess
	StateFailed
	StateCancelled
)

func (s BuildState) String() string {
	switch s {
	case StateIdle:
		return "idle"
	case StateQueued:
		return "queued"
	case StateBuilding:
		return "building"
	case StateSuccess:
		return "success"
	case StateFailed:
		return "failed"
	case StateCancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

type BuildTool int

const (
	ToolUnknown BuildTool = iota
	ToolCargo
	ToolNPM
	ToolMake
	ToolGo
	ToolMSBuild
	ToolGradle
	ToolCMake
	ToolGeneric
)

func (t BuildTool) String() string {
	switch t {
	case ToolCargo:
		return "cargo"
	case ToolNPM:
		return "npm"
	case ToolMake:
		return "make"
	case ToolGo:
		return "go"
	case ToolMSBuild:
		return "msbuild"
	case ToolGradle:
		return "gradle"
	case ToolCMake:
		return "cmake"
	default:
		return "generic"
	}
}

type LogLevel int

const (
	LogInfo LogLevel = iota
	LogWarning
	LogError
	LogNote
)

type ParsedEntry struct {
	File    string
	Line    int
	Column  int
	Code    string
	Message string
}

type LogLine struct {
	Timestamp time.Time
	Level     LogLevel
	Raw       string
	Parsed    *ParsedEntry
}

type BuildError struct {
	File    string
	Line    int
	Column  int
	Code    string
	Message string
	Level   LogLevel
}

type Build struct {
	ID          string
	Name        string
	Tool        BuildTool
	Command     string
	WorkDir     string
	State       BuildState
	Progress    float64
	StartTime   time.Time
	EndTime     *time.Time
	Duration    time.Duration
	PID         int
	LogLines    []LogLine
	Errors      []BuildError
	Tags        []string
	ProfileName string
}

func NewBuild(name, command, workDir string, tool BuildTool) *Build {
	return &Build{
		ID:      uuid.New().String(),
		Name:    name,
		Command: command,
		WorkDir: workDir,
		Tool:    tool,
		State:   StateQueued,
	}
}

type BuildProfile struct {
	Name    string
	WorkDir string
	Command string
	Tags    []string
	Env     map[string]string
}
