package tui

import (
	"fmt"
	"strings"
)

func TabBarView(active Tab, width int) string {
	var parts []string
	parts = append(parts, BrandStyle.Render("VISI-ON"))
	for i, label := range tabLabels {
		num := TabNumberStyle.Render(fmt.Sprintf("%d", i+1))
		if Tab(i) == active {
			parts = append(parts, TabActiveStyle.Render(num+" "+label))
		} else {
			parts = append(parts, TabInactiveStyle.Render(num+" "+label))
		}
	}
	return TabBarStyle.Width(width).Render(strings.Join(parts, " "))
}
