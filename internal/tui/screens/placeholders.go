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
  1-5       `) + glowStyle.Render("Switch tabs") + dimStyle.Render(`          Tab     Next panel
  `) + glowStyle.Render("q") + dimStyle.Render(`         Quit                 `) + glowStyle.Render("Esc") + dimStyle.Render(`     Back to Normal
  `) + glowStyle.Render(":") + dimStyle.Render(`         Command mode         `) + glowStyle.Render("i") + dimStyle.Render(`       Insert mode (Launcher)

  TIER 2 — BUILD CONTROL (Normal, Dashboard)
  `) + glowStyle.Render("r") + dimStyle.Render(`  Run build    `) + glowStyle.Render("x") + dimStyle.Render(`  Kill    `) + glowStyle.Render("f") + dimStyle.Render(`  Focus log    `) + glowStyle.Render("o") + dimStyle.Render(`  Open in editor

  TIER 3 — LOG NAVIGATION (Normal, log panel focused)
  `) + glowStyle.Render("j/k") + dimStyle.Render(` Scroll    `) + glowStyle.Render("g") + dimStyle.Render(` Top    `) + glowStyle.Render("G") + dimStyle.Render(` Bottom+follow    `) + glowStyle.Render("/") + dimStyle.Render(` Search    `) + glowStyle.Render("n/N") + dimStyle.Render(` Match

  TIER 4 — COMMAND MODE
  `) + glowStyle.Render(":run <cmd>") + dimStyle.Render(`          run a build
  `) + glowStyle.Render(":kill <id>") + dimStyle.Render(`          kill a build
  `) + glowStyle.Render(":profile save <n>") + dimStyle.Render(`   save a profile
  `) + glowStyle.Render(":q") + dimStyle.Render(`                  quit

  MODES:  `) + blueStyle.Render("NORMAL") + dimStyle.Render(`=blue   `) + accentStyle.Render("INSERT") + dimStyle.Render(`=teal   `) + mutedStyle.Render("COMMAND") + dimStyle.Render(`=amber
  `) + glowStyle.Render("Esc") + dimStyle.Render(` always returns to Normal from any mode.`)
	return panelElevated.Width(w - 2).Height(h).Render(content)
}

func placeholder(title, hint string, w, h int) string {
	return panelElevated.Width(w - 2).Height(h).Render(
		"\n" + superHeader(strings.TrimPrefix(title, "▸ ")) +
			"\n\n" + dimStyle.Render("  "+hint) +
			"\n\n" + accentStyle.Render("  ✓ Enhanced interactivity in progress") +
			"\n" + mutedStyle.Render("  Tip: use :help to see key commands to explore the interface."),
	)
}
