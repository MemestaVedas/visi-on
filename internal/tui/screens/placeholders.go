package screens

import "strings"

func LauncherView(w, h int) string {
	return placeholder("▸ NEW BUILD",
		"Build launcher — type a directory and command.", w, h)
}

func HistoryView(w, h int) string {
	return placeholder("▸ BUILD HISTORY",
		"Past builds — status, duration, error counts.", w, h)
}

func PluginsView(w, h int) string {
	return placeholder("▸ PLUGINS",
		"Slack, Discord, git-detect — manage plugins here.", w, h)
}

func HelpView(w, h int) string {
	content := blueStyle.Render("  ▸ KEYBINDINGS\n\n") + dimStyle.Render(`
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

  MODES:  NORMAL=blue   INSERT=green   COMMAND=orange
  Esc always returns to Normal from any mode.`)
	return panelFoc.Width(w - 2).Height(h).Render(content)
}

func placeholder(title, hint string, w, h int) string {
	return panelFoc.Width(w - 2).Height(h).Render(
		"\n\n" + blueStyle.Render("  "+strings.ToUpper(title)) +
			"\n\n" + dimStyle.Render("  "+hint) +
			"\n\n" + mutedStyle.Render("  Coming in the next phase."),
	)
}
