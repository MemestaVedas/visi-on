package colors

import "github.com/charmbracelet/lipgloss"

const (
	NightInk      = "#0b1220"
	NightDeep     = "#121a2b"
	NightPanel    = "#161f33"
	NightPanelAlt = "#101827"
	NightBorder   = "#2c3d5d"
	NightBorderHi = "#4f7cff"
	NightText     = "#d7e2ff"
	NightSubtext  = "#a9b7d9"
	NightMuted    = "#6f819f"
	NightBlue     = "#4f7cff"
	NightTeal     = "#30d0b7"
	NightAmber    = "#ffb454"
	NightRed      = "#ff6b81"
	NightGreen    = "#3dd68c"
	// Focus & glow accents
	NightGlowBlue    = "#6d95ff"
	NightGlowTeal    = "#52dfc9"
	NightAccentFaint = "#1f2e47"
)

var (
	Bg          = lipgloss.Color(NightInk)
	BgPanel     = lipgloss.Color(NightPanel)
	BgPanelAlt  = lipgloss.Color(NightPanelAlt)
	BgAccent    = lipgloss.Color(NightAccentFaint)
	Border      = lipgloss.Color(NightBorder)
	BorderFocus = lipgloss.Color(NightBorderHi)
	GlowBlue    = lipgloss.Color(NightGlowBlue)
	GlowTeal    = lipgloss.Color(NightGlowTeal)

	Normal  = lipgloss.Color(NightBlue)
	Insert  = lipgloss.Color(NightTeal)
	Command = lipgloss.Color(NightAmber)

	TextPrimary   = lipgloss.Color(NightText)
	TextSecondary = lipgloss.Color(NightSubtext)
	TextMuted     = lipgloss.Color(NightMuted)

	Success = lipgloss.Color(NightGreen)
	Error   = lipgloss.Color(NightRed)
	Running = lipgloss.Color(NightBlue)
)
