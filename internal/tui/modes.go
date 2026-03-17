package tui

type Mode int

const (
    ModeNormal  Mode = iota
    ModeInsert
    ModeCommand
)

type Tab int

const (
    TabDashboard Tab = iota
    TabLauncher
    TabHistory
    TabPlugins
    TabHelp
    tabCount
)

var tabLabels = []string{"Dashboard", "Launcher", "History", "Plugins", "Help"}
