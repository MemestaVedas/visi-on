package tui

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/visi-on/visi-on/internal/tui/screens"
)

type AppModel struct {
    mode      Mode
    activeTab Tab
    w, h      int
    statusBar StatusBarModel
    commandBuf string
}

func NewAppModel() AppModel { return AppModel{mode: ModeNormal, activeTab: TabDashboard} }

func (m AppModel) Init() tea.Cmd { return nil }

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.w, m.h = msg.Width, msg.Height
    case tea.KeyMsg:
        return m.key(msg)
    }
    return m, nil
}

func (m AppModel) key(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    // Esc: always back to Normal — checked first, no exceptions
    if msg.String() == "esc" {
        m.mode = ModeNormal
        m.commandBuf = ""
        return m, nil
    }
    switch m.mode {
    case ModeNormal:  return m.normalKey(msg)
    case ModeInsert:  return m.insertKey(msg)
    case ModeCommand: return m.commandKey(msg)
    }
    return m, nil
}

func (m AppModel) normalKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.String() {
    case "1": m.activeTab = TabDashboard
    case "2": m.activeTab = TabLauncher
    case "3": m.activeTab = TabHistory
    case "4": m.activeTab = TabPlugins
    case "5", "?": m.activeTab = TabHelp
    case "q":  return m, tea.Quit
    case ":":  m.mode = ModeCommand
    case "i":
        if m.activeTab == TabLauncher { m.mode = ModeInsert }
    }
    return m, nil
}

func (m AppModel) insertKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    // Phase 2: wire input fields
    return m, nil
}

func (m AppModel) commandKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
    switch msg.Type {
    case tea.KeyEnter:
        cmd := m.commandBuf
        m.mode = ModeNormal
        m.commandBuf = ""
        return m, m.execCmd(cmd)
    case tea.KeyBackspace:
        if len(m.commandBuf) > 0 { m.commandBuf = m.commandBuf[:len(m.commandBuf)-1] }
    case tea.KeyRunes:
        m.commandBuf += string(msg.Runes)
    }
    return m, nil
}

func (m AppModel) execCmd(cmd string) tea.Cmd {
    if cmd == "q" || cmd == "quit" { return tea.Quit }
    return nil
}

func (m AppModel) View() string {
    if m.w == 0 { return "" }
    tabBar    := TabBarView(m.activeTab, m.w)
    statusBar := m.statusBar.View(m.mode, m.w)
    hintsBar  := HintsBarView(m.activeTab, m.mode, m.commandBuf, m.w)
    chromeH   := lipgloss.Height(tabBar) + lipgloss.Height(statusBar) + lipgloss.Height(hintsBar) + 2
    wsH       := m.h - chromeH
    if wsH < 1 { wsH = 1 }
    return lipgloss.JoinVertical(lipgloss.Left,
        tabBar, m.workspace(m.w, wsH), statusBar, hintsBar)
}

func (m AppModel) workspace(w, h int) string {
    switch m.activeTab {
    case TabDashboard:
        d := screens.NewDashboard()
        d.Width, d.Height = w, h
        return d.View()
    case TabLauncher: return screens.LauncherView(w, h)
    case TabHistory:  return screens.HistoryView(w, h)
    case TabPlugins:  return screens.PluginsView(w, h)
    case TabHelp:     return screens.HelpView(w, h)
    }
    return ""
}
