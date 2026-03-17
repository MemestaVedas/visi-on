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
	title := AppTitleStyle.Render("VISI-ON v1.0")
	pill := ModeStyle(m).Render(ModeName(m))
	conn := TextMuted.Render("○ Offline")
	if s.Connected {
		conn = TextSuccess.Render("● Connected")
	}
	left := strings.Join([]string{title, sep, pill, sep, conn}, " ")
	right := TextSecondary.Render(fmt.Sprintf("CPU %.0f%%  RAM %.1fG  ↑%.0f ↓%.0fKB",
		s.CPUPct, s.RAMUsedGB, s.NetUpKB, s.NetDownKB))
	gap := width - lipgloss.Width(left) - lipgloss.Width(right) - 2
	if gap < 1 {
		gap = 1
	}
	return StatusBarStyle.Width(width).Render(left + strings.Repeat(" ", gap) + right)
}
