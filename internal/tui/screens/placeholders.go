package screens

import (
	"strings"
)

func LauncherView(w, h int) string {
	return placeholder("Build Launcher",
		"Create and run build profiles from a single command lane.", w, h)
}

func HistoryView(w, h int) string {
	return placeholder("Build History",
		"Explore previous runs with status, duration, and error density.", w, h)
}

func PluginsView(w, h int) string {
	return placeholder("Plugins",
		"Manage integrations like Slack, Discord, and repository analyzers.", w, h)
}

func HelpView(w, h int) string {
	content := superHeader("Keybindings") + "\n\n" + dimStyle.Render(`
  1-5       Switch tabs          Tab     Next panel
  q         Quit                 Esc     Back to Normal
  :         Command mode         i       Insert mode (Launcher)

  TIER 2 — BUILD CONTROL (Normal, Dashboard)
  r  Run build    x  Kill    f  Focus log    o  Open in editor

  TIER 3 — LOG NAVIGATION (Normal, log panel focused)
  j/k Scroll    g Top    G Bottom+follow    / Search    n/N Match

  TIER 4 — COMMAND MODE
  :run <cmd>          run a build
  :kill <id>          kill a build
  :profile save <n>   save a profile
  :q                  quit

  MODES:  NORMAL=blue   INSERT=teal   COMMAND=amber
  Esc always returns to Normal from any mode.`)
	return panelFoc.Width(w - 2).Height(h).Render(content)
}

func placeholder(title, hint string, w, h int) string {
	return panelFoc.Width(w - 2).Height(h).Render(
		"\n" + superHeader(strings.TrimPrefix(title, "▸ ")) +
			"\n\n" + dimStyle.Render("  "+hint) +
			"\n\n" + mutedStyle.Render("  This workspace will become interactive in the next phase.") +
			"\n" + mutedStyle.Render("  Tip: use :help to see key commands while this screen evolves."),
	)
}
