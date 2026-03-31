package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func TabBarView(active Tab, width int) string {
	var parts []string
	parts = append(parts, BrandStyle.Render("VISI-ON"))
	for i, label := range tabLabels {
		num := TabNumberStyle.Render(fmt.Sprintf("%d", i+1))
		tabText := num + " " + label
		if Tab(i) == active {
			// Enhanced active tab with glow effect
			styled := TabActiveStyle.Render(tabText)
			parts = append(parts, lipgloss.NewStyle().Bold(true).Render(styled))
		} else {
			parts = append(parts, TabInactiveStyle.Render(tabText))
		}
	}
	return TabBarStyle.Width(width).Render(strings.Join(parts, "  "))
}
