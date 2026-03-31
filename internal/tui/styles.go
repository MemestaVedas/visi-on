package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/visi-on/visi-on/internal/tui/colors"
)

var (
	colorBg          = colors.Bg
	colorBgPanel     = colors.BgPanel
	colorBgPanelAlt  = colors.BgPanelAlt
	colorBorder      = colors.Border
	colorBorderFocus = colors.BorderFocus

	colorNormal  = colors.Normal
	colorInsert  = colors.Insert
	colorCommand = colors.Command

	colorTextPrimary   = colors.TextPrimary
	colorTextSecondary = colors.TextSecondary
	colorTextMuted     = colors.TextMuted

	colorSuccess = colors.Success
	colorError   = colors.Error
	colorRunning = colors.Running
)

var (
	PanelStyle = lipgloss.NewStyle().
			Background(colorBgPanel).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(0, 1)

	PanelFocusedStyle = lipgloss.NewStyle().
				Background(colorBgPanel).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(colorBorderFocus).
				Padding(0, 1)

	TabBarStyle      = lipgloss.NewStyle().Background(colorBg).Padding(0, 1)
	TabActiveStyle   = lipgloss.NewStyle().Foreground(colorTextPrimary).Background(colorBgPanel).Bold(true).Padding(0, 2).Border(lipgloss.NormalBorder()).BorderForeground(colorBorderFocus)
	TabInactiveStyle = lipgloss.NewStyle().Foreground(colorTextSecondary).Background(colorBgPanelAlt).Padding(0, 2).Border(lipgloss.NormalBorder()).BorderForeground(colorBorder)
	TabNumberStyle   = lipgloss.NewStyle().Foreground(colorTextMuted)
	BrandStyle       = lipgloss.NewStyle().Foreground(colorTextPrimary).Background(colorNormal).Bold(true).Padding(0, 1).MarginRight(1)

	StatusBarStyle = lipgloss.NewStyle().Background(colorBgPanel).Foreground(colorTextSecondary).Padding(0, 1)
	HintsBarStyle  = lipgloss.NewStyle().Background(colorBgPanelAlt).Foreground(colorTextMuted).Padding(0, 1)
	HintKeyStyle   = lipgloss.NewStyle().Foreground(colorTextPrimary).Bold(true)
	AppTitleStyle  = lipgloss.NewStyle().Foreground(colorNormal).Bold(true)
	SepStyle       = lipgloss.NewStyle().Foreground(colorBorder).SetString("•")

	TextPrimary   = lipgloss.NewStyle().Foreground(colorTextPrimary)
	TextSecondary = lipgloss.NewStyle().Foreground(colorTextSecondary)
	TextMuted     = lipgloss.NewStyle().Foreground(colorTextMuted)
	TextSuccess   = lipgloss.NewStyle().Foreground(colorSuccess)
	TextError     = lipgloss.NewStyle().Foreground(colorError)
	TextRunning   = lipgloss.NewStyle().Foreground(colorRunning)
	TextWarning   = lipgloss.NewStyle().Foreground(colorCommand)

	MetricOKStyle   = lipgloss.NewStyle().Foreground(colorTextPrimary).Background(colorBgPanelAlt).Padding(0, 1)
	MetricWarnStyle = lipgloss.NewStyle().Foreground(colorCommand).Background(colorBgPanelAlt).Padding(0, 1)
	MetricErrStyle  = lipgloss.NewStyle().Foreground(colorError).Background(colorBgPanelAlt).Padding(0, 1)
)

func ModeStyle(m Mode) lipgloss.Style {
	switch m {
	case ModeInsert:
		return lipgloss.NewStyle().Background(colorInsert).Foreground(colorBg).Bold(true).Padding(0, 1).Border(lipgloss.NormalBorder()).BorderForeground(colorInsert)
	case ModeCommand:
		return lipgloss.NewStyle().Background(colorCommand).Foreground(colorBg).Bold(true).Padding(0, 1).Border(lipgloss.NormalBorder()).BorderForeground(colorCommand)
	default:
		return lipgloss.NewStyle().Background(colorNormal).Foreground(colorBg).Bold(true).Padding(0, 1).Border(lipgloss.NormalBorder()).BorderForeground(colorNormal)
	}
}

func ModeName(m Mode) string {
	switch m {
	case ModeInsert:
		return " INSERT "
	case ModeCommand:
		return " COMMAND "
	default:
		return " NORMAL "
	}
}
