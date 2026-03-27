package colors

import "github.com/charmbracelet/lipgloss"

const (
	MochaCrust     = "#11111b"
	MochaMantle    = "#181825"
	MochaBase      = "#1e1e2e"
	MochaSurface0  = "#313244"
	MochaSurface1  = "#45475a"
	MochaOverlay0  = "#6c7086"
	MochaText      = "#cdd6f4"
	MochaSubtext0  = "#a6adc8"
	MochaLavender  = "#b4befe"
	MochaBlue      = "#89b4fa"
	MochaSapphire  = "#74c7ec"
	MochaSky       = "#89dceb"
	MochaTeal      = "#94e2d5"
	MochaGreen     = "#a6e3a1"
	MochaYellow    = "#f9e2af"
	MochaPeach     = "#fab387"
	MochaMaroon    = "#eba0ac"
	MochaRed       = "#f38ba8"
	MochaMauve     = "#cba6f7"
	MochaPink      = "#f5c2e7"
	MochaFlamingo  = "#f2cdcd"
	MochaRosewater = "#f5e0dc"
)

var (
	Bg          = lipgloss.Color(MochaBase)
	BgPanel     = lipgloss.Color(MochaMantle)
	BgPanelAlt  = lipgloss.Color(MochaCrust)
	Border      = lipgloss.Color(MochaSurface0)
	BorderFocus = lipgloss.Color(MochaBlue)

	Normal  = lipgloss.Color(MochaBlue)
	Insert  = lipgloss.Color(MochaGreen)
	Command = lipgloss.Color(MochaPeach)

	TextPrimary   = lipgloss.Color(MochaText)
	TextSecondary = lipgloss.Color(MochaSubtext0)
	TextMuted     = lipgloss.Color(MochaOverlay0)

	Success = lipgloss.Color(MochaGreen)
	Error   = lipgloss.Color(MochaRed)
	Running = lipgloss.Color(MochaBlue)
)
