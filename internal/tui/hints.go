package tui

import "strings"

var tabHints = map[Tab]string{
	TabDashboard: "r Run   x Kill   f Focus Log   o Open Editor   / Search",
	TabLauncher:  "Tab Complete   Esc Normal   Ctrl+W Delete Word   Up/Down Move",
	TabHistory:   "Enter View Logs   o Open Project   r Re-run   / Search",
	TabPlugins:   "Enter Configure   e Enable/Disable   :plugin install <n>",
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
	for _, p := range parts {
		kv := strings.SplitN(p, " ", 2)
		if len(kv) == 2 {
			out = append(out, HintKeyStyle.Render(kv[0])+" "+TextMuted.Render(kv[1]))
		} else {
			out = append(out, TextMuted.Render(p))
		}
	}
	return HintsBarStyle.Width(width).Render(strings.Join(out, TextMuted.Render("   ")))
}
