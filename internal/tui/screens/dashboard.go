package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/visi-on/visi-on/internal/tui/colors"
)

var (
	border      = colors.Border
	borderFocus = colors.BorderFocus
	glowBlue    = colors.GlowBlue
	glowTeal    = colors.GlowTeal
	accentBg    = colors.BgAccent
	blue        = colors.Normal
	muted       = colors.TextMuted
	dim         = colors.TextSecondary
	body        = colors.TextPrimary

	panelStyle    = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(border).Padding(0, 1)
	panelFoc      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(borderFocus).Padding(0, 1)
	panelElevated = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(glowBlue).Padding(0, 1)
	titleStyle    = lipgloss.NewStyle().Foreground(body).Bold(true)
	dimStyle      = lipgloss.NewStyle().Foreground(dim)
	mutedStyle    = lipgloss.NewStyle().Foreground(muted)
	blueStyle     = lipgloss.NewStyle().Foreground(blue).Bold(true)
	glowStyle     = lipgloss.NewStyle().Foreground(glowBlue).Bold(true)
	accentStyle   = lipgloss.NewStyle().Foreground(body).Background(accentBg).Padding(0, 1)
)

type DashboardModel struct {
	Width, Height int
	Animation     interface{} // AnimationState pointer
}

func NewDashboard() DashboardModel { return DashboardModel{} }

func superHeader(title string) string {
	return blueStyle.Render("[ ") + titleStyle.Render(strings.ToUpper(title)) + blueStyle.Render(" ]")
}

func (d DashboardModel) View() string {
	if d.Width < 20 || d.Height < 10 {
		return ""
	}

	sideW := 40
	mainW := d.Width - sideW - 3

	leftTop := panelElevated.Width(sideW).Height(d.Height / 2).Render(
		superHeader("Active Builds") + "\n\n" +
			dimStyle.Render("  Queue is idle. No active pipelines.") + "\n\n" +
			mutedStyle.Render("  Start a run: ") + glowStyle.Render(":run <command>") + "\n" +
			mutedStyle.Render("  Quick key: ") + glowStyle.Render("r") + mutedStyle.Render(" in Launcher") + "\n\n" +
			accentStyle.Render("  Last result") + "\n" +
			dimStyle.Render("  duration  0.0s   errors  0   warnings  0"),
	)

	leftBot := panelStyle.Width(sideW).Height(d.Height - d.Height/2 - 1).Render(
		superHeader("Error Analysis") + "\n\n" +
			dimStyle.Render("  Workspace health is clean.") + "\n\n" +
			mutedStyle.Render("  Most common failures") + "\n" +
			dimStyle.Render("  1) none") + "\n" +
			dimStyle.Render("  2) none") + "\n" +
			dimStyle.Render("  3) none"),
	)

	rightTop := panelStyle.Width(mainW).Height(d.Height / 2).Render(
		superHeader("Pipeline Timeline") + "\n\n" +
			renderFlame(mainW),
	)

	rightBot := panelStyle.Width(mainW).Height(d.Height - d.Height/2 - 1).Render(
		superHeader("Log Output") + "\n\n" +
			dimStyle.Render("  Live stream appears when a build starts.") + "\n\n" +
			mutedStyle.Render("  Controls: ") + glowStyle.Render("G") + mutedStyle.Render(" follow   ") +
			glowStyle.Render("/") + mutedStyle.Render(" search   ") +
			glowStyle.Render("f") + mutedStyle.Render(" focus panel"),
	)

	left := lipgloss.JoinVertical(lipgloss.Left, leftTop, leftBot)
	right := lipgloss.JoinVertical(lipgloss.Left, rightTop, rightBot)
	return lipgloss.JoinHorizontal(lipgloss.Top, left, " ", right)
}

func renderFlame(width int) string {
	return renderFlameWithAnimation(width, nil)
}

func renderFlameWithAnimation(width int, anim interface{}) string {
	bw := width - 30
	if bw < 4 {
		bw = 4
	}
	stages := []string{"  Compile", "  Link   ", "  Bundle "}
	stageColors := []lipgloss.Color{colors.GlowBlue, colors.GlowTeal, colors.Normal}
	percentages := []float64{20, 35, 50}
	var lines []string
	for i, s := range stages {
		pct := percentages[i]
		fill := int(float64(bw) * pct / 100)
		if fill < 1 {
			fill = 1
		}
		colorStyle := lipgloss.NewStyle().Foreground(stageColors[i])
		bar := colorStyle.Render(strings.Repeat("█", fill)) + dimStyle.Render(strings.Repeat("░", bw-fill))
		lines = append(lines, fmt.Sprintf("%s %s  %.0f%%", s, bar, pct))
	}
	return strings.Join(lines, "\n")
}
