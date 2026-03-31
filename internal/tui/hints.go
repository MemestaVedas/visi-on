package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	HintButtonStyle = lipgloss.NewStyle().
			Foreground(colorTextPrimary).
			Background(colorBgPanelAlt).
			Padding(0, 1).
			Bold(true)
	HintButtonActiveStyle = lipgloss.NewStyle().
				Foreground(colorBg).
				Background(colorGlowBlue).
				Padding(0, 1).
				Bold(true)
)

var tabHints = map[Tab]string{
	TabDashboard: "r Run build   x Kill build   f Focus logs   o Open editor   / Search logs",
	TabLauncher:  "Tab Complete command   Esc Normal mode   Ctrl+W Delete word   Up/Down History",
	TabHistory:   "Enter Open logs   o Open project   r Re-run build   / Search entries",
	TabPlugins:   "Enter Configure plugin   e Toggle plugin   :plugin install <name>",
	TabHelp:      "j/k Scroll   q Quit",
}

func HintsBarView(activeTab Tab, mode Mode, commandBuf string, width int) string {
	if mode == ModeCommand {
		return HintsBarStyle.Width(width).Render(
			TextWarning.Render(":") + TextPrimary.Render(commandBuf) + TextMuted.Render("█"),
		)
	}
	hint := tabHints[activeTab]
	parts := strings.Split(hint, "   ")
	var out []string
	separator := TextMuted.Render(" • ")
	for i, p := range parts {
		kv := strings.SplitN(p, " ", 2)
		if len(kv) == 2 {
			// Alternate button styles for visual rhythm
			var keyStyle lipgloss.Style
			if i%2 == 0 {
				keyStyle = HintButtonStyle
			} else {
				keyStyle = lipgloss.NewStyle().Foreground(colorGlowBlue).Bold(true)
			}
			out = append(out, keyStyle.Render(kv[0])+" "+TextSecondary.Render(kv[1]))
		} else {
			out = append(out, TextMuted.Render(p))
		}
	}
	return HintsBarStyle.Width(width).Render(strings.Join(out, separator))
}
