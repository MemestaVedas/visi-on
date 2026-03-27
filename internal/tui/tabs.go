package tui

import (
	"fmt"
	"strings"
)

func TabBarView(active Tab, width int) string {
	var parts []string
	for i, label := range tabLabels {
		num := TabNumberStyle.Render(fmt.Sprintf("%d ", i+1))
		if Tab(i) == active {
			parts = append(parts, TabActiveStyle.Render(label))
		} else {
			parts = append(parts, num+TabInactiveStyle.Render(label))
		}
	}
	return TabBarStyle.Width(width).Render(" " + strings.Join(parts, " "))
}
