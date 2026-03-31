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
	blue        = colors.Normal
	muted       = colors.TextMuted
	dim         = colors.TextSecondary
	body        = colors.TextPrimary

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

func superHeader(title string) string {
	return blueStyle.Render("[ ") + titleStyle.Render(strings.ToUpper(title)) + blueStyle.Render(" ]")
}

func (d DashboardModel) View() string {
	if d.Width < 20 || d.Height < 10 {
		return ""
	}

	sideW := 40
	mainW := d.Width - sideW - 3

	leftTop := panelFoc.Width(sideW).Height(d.Height / 2).Render(
		superHeader("Active Builds") + "\n\n" +
			dimStyle.Render("  Queue is idle. No active pipelines.") + "\n\n" +
			mutedStyle.Render("  Start a run: ") + blueStyle.Render(":run <command>") + "\n" +
			mutedStyle.Render("  Quick key: ") + blueStyle.Render("r") + mutedStyle.Render(" in Launcher") + "\n\n" +
			mutedStyle.Render("  Last result") + "\n" +
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
			mutedStyle.Render("  Controls: ") + blueStyle.Render("G") + mutedStyle.Render(" follow   ") +
			blueStyle.Render("/") + mutedStyle.Render(" search   ") +
			blueStyle.Render("f") + mutedStyle.Render(" focus panel"),
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
	for i, s := range stages {
		fill := i * bw / 6
		if fill < 1 {
			fill = 1
		}
		bar := blueStyle.Render(strings.Repeat("█", fill)) + mutedStyle.Render(strings.Repeat("░", bw-fill))
		lines = append(lines, fmt.Sprintf("  %s %s  %d%%", s, bar, i*15))
	}
	return strings.Join(lines, "\n")
}
