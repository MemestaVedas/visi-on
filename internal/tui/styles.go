package tui

import "github.com/charmbracelet/lipgloss"

const (
    colorBg          = lipgloss.Color("#0d0f14")
    colorBgPanel     = lipgloss.Color("#12151c")
    colorBgPanelAlt  = lipgloss.Color("#161a24")
    colorBorder      = lipgloss.Color("#2a2f3d")
    colorBorderFocus = lipgloss.Color("#4a90d9")

    colorNormal  = lipgloss.Color("#4a90d9")
    colorInsert  = lipgloss.Color("#56d4a0")
    colorCommand = lipgloss.Color("#e8a245")

    colorTextPrimary   = lipgloss.Color("#d4d8e8")
    colorTextSecondary = lipgloss.Color("#6b7594")
    colorTextMuted     = lipgloss.Color("#3d4460")

    colorSuccess = lipgloss.Color("#56d4a0")
    colorError   = lipgloss.Color("#e05c6e")
    colorRunning = lipgloss.Color("#4a90d9")
)

var (
    PanelStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(colorBorder).
        Padding(0, 1)

    PanelFocusedStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(colorBorderFocus).
        Padding(0, 1)

    TabBarStyle      = lipgloss.NewStyle().Background(colorBg)
    TabActiveStyle   = lipgloss.NewStyle().Foreground(colorNormal).Bold(true).Padding(0, 2)
    TabInactiveStyle = lipgloss.NewStyle().Foreground(colorTextMuted).Padding(0, 2)
    TabNumberStyle   = lipgloss.NewStyle().Foreground(colorTextMuted)

    StatusBarStyle  = lipgloss.NewStyle().Background(colorBgPanel).Foreground(colorTextSecondary).Padding(0, 1)
    HintsBarStyle   = lipgloss.NewStyle().Background(colorBgPanelAlt).Foreground(colorTextMuted).Padding(0, 1)
    HintKeyStyle    = lipgloss.NewStyle().Foreground(colorTextSecondary).Bold(true)
    AppTitleStyle   = lipgloss.NewStyle().Foreground(colorNormal).Bold(true)
    SepStyle        = lipgloss.NewStyle().Foreground(colorBorder).SetString("│")

    TextPrimary   = lipgloss.NewStyle().Foreground(colorTextPrimary)
    TextSecondary = lipgloss.NewStyle().Foreground(colorTextSecondary)
    TextMuted     = lipgloss.NewStyle().Foreground(colorTextMuted)
    TextSuccess   = lipgloss.NewStyle().Foreground(colorSuccess)
    TextError     = lipgloss.NewStyle().Foreground(colorError)
    TextRunning   = lipgloss.NewStyle().Foreground(colorRunning)
    TextWarning   = lipgloss.NewStyle().Foreground(colorCommand)
)

func ModeStyle(m Mode) lipgloss.Style {
    switch m {
    case ModeInsert:  return lipgloss.NewStyle().Background(colorInsert).Foreground(colorBg).Bold(true).Padding(0, 1)
    case ModeCommand: return lipgloss.NewStyle().Background(colorCommand).Foreground(colorBg).Bold(true).Padding(0, 1)
    default:          return lipgloss.NewStyle().Background(colorNormal).Foreground(colorBg).Bold(true).Padding(0, 1)
    }
}

func ModeName(m Mode) string {
    switch m {
    case ModeInsert:  return " INSERT "
    case ModeCommand: return " COMMAND "
    default:          return " NORMAL "
    }
}
