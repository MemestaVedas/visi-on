package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	border      = lipgloss.Color("#2a2f3d")
	borderFocus = lipgloss.Color("#4a90d9")
	blue        = lipgloss.Color("#4a90d9")
	muted       = lipgloss.Color("#3d4460")
	dim         = lipgloss.Color("#6b7594")
	body        = lipgloss.Color("#d4d8e8")

	panelStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(border).Padding(0, 1)
	panelFoc   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(borderFocus).Padding(0, 1)
	titleStyle = lipgloss.NewStyle().Foreground(body).Bold(true)
	dimStyle   = lipgloss.NewStyle().Foreground(dim)
	mutedStyle = lipgloss.NewStyle().Foreground(muted)
	blueStyle  = lipgloss.NewStyle().Foreground(blue).Bold(true)
)

type DashboardModel struct {
	Width, Height int
}

func NewDashboard() DashboardModel { return DashboardModel{} }

func (d DashboardModel) View() string {
	if d.Width < 20 || d.Height < 10 {
		return ""
	}

	sideW := 40
	mainW := d.Width - sideW - 3

	leftTop := panelFoc.Width(sideW).Height(d.Height / 2).Render(
		blueStyle.Render("▸ ACTIVE BUILDS (0)") + "\n\n" +
			dimStyle.Render("  Waiting for builds...\n\n") +
			mutedStyle.Render("  :run <command>\n") +
			mutedStyle.Render("  or press  r  in Launcher"),
	)

	leftBot := panelStyle.Width(sideW).Height(d.Height - d.Height/2 - 1).Render(
		titleStyle.Render("  ERROR ANALYSIS (0)") + "\n\n" +
			dimStyle.Render("  No errors — all clear ✓"),
	)

	rightTop := panelStyle.Width(mainW).Height(d.Height / 2).Render(
		titleStyle.Render("  FLAMECHART") + "\n\n" +
			renderFlame(mainW),
	)

	rightBot := panelStyle.Width(mainW).Height(d.Height - d.Height/2 - 1).Render(
		titleStyle.Render("  LOG OUTPUT") + "\n\n" +
			dimStyle.Render("  No active build — logs stream here\n") +
			dimStyle.Render("  G=follow  /=search  f=focus"),
	)

	left := lipgloss.JoinVertical(lipgloss.Left, leftTop, leftBot)
	right := lipgloss.JoinVertical(lipgloss.Left, rightTop, rightBot)
	return lipgloss.JoinHorizontal(lipgloss.Top, left, " ", right)
}

func renderFlame(width int) string {
	bw := width - 22
	if bw < 4 {
		bw = 4
	}
	stages := []string{"Compile", "Link   ", "Bundle "}
	var lines []string
	for _, s := range stages {
		bar := mutedStyle.Render(strings.Repeat("░", bw))
		lines = append(lines, fmt.Sprintf("  %s %s  0%%", s, bar))
	}
	return strings.Join(lines, "\n")
}
