package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type StatusBarModel struct {
	Connected bool
	CPUPct    float64
	RAMUsedGB float64
	NetUpKB   float64
	NetDownKB float64
}

func (s StatusBarModel) View(m Mode, width int) string {
	sep := SepStyle.Render()
	title := AppTitleStyle.Render("build orchestration")
	pill := ModeStyle(m).Render(ModeName(m))

	// Enhanced connection indicator
	var conn string
	if s.Connected {
		conn = lipgloss.NewStyle().Foreground(colorSuccess).Bold(true).Render("●") + " CONNECTED"
	} else {
		conn = lipgloss.NewStyle().Foreground(colorError).Bold(true).Render("●") + " OFFLINE"
	}

	left := strings.Join([]string{title, sep, pill, sep, conn}, " ")
	right := strings.Join([]string{
		MetricWarnStyle.Render(fmt.Sprintf("CPU %.0f%%", s.CPUPct)),
		MetricOKStyle.Render(fmt.Sprintf("RAM %.1fG", s.RAMUsedGB)),
		TextSecondary.Render(fmt.Sprintf("↑ %.0fKB", s.NetUpKB)),
		TextSecondary.Render(fmt.Sprintf("↓ %.0fKB", s.NetDownKB)),
	}, " ")
	gap := width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if gap < 1 {
		gap = 1
	}
	return StatusBarStyle.Width(width).Render(left + strings.Repeat(" ", gap) + right)
}
