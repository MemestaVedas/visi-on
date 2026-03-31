package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/visi-on/visi-on/internal/tui/colors"
)

// AnimationFrame holds animation state for spinners and effects
type AnimationFrame struct {
	tick int
}

// SpinnerFrames defines ASCII spinner animation frames
var SpinnerFrames = []string{"◐", "◓", "◑", "◒"}

// PulseFrames defines pulse animation for activity indicators
var PulseFrames = []string{"●", "◐", "◑", "◒", "◑", "◐"}

// ProgressBarFrames defines flowing progress bar animation
var ProgressBarFrames = []string{"▉", "█", "▊", "█"}

// AnimationState maintains global animation tick
type AnimationState struct {
	Tick int
}

// NextFrame advances animation frame
func (a *AnimationState) NextFrame() {
	a.Tick = (a.Tick + 1) % 60 // 60-frame cycle
}

// GetSpinner returns current spinner frame
func (a *AnimationState) GetSpinner() string {
	return SpinnerFrames[a.Tick%len(SpinnerFrames)]
}

// GetPulse returns current pulse frame with optional coloring
func (a *AnimationState) GetPulse(color lipgloss.Color) string {
	frame := PulseFrames[a.Tick%len(PulseFrames)]
	return lipgloss.NewStyle().Foreground(color).Bold(true).Render(frame)
}

// GetProgressBar returns flowing progress bar fill
func (a *AnimationState) GetProgressBar(width int, percentage float64, color lipgloss.Color) string {
	if width < 4 {
		width = 4
	}

	filledWidth := int(float64(width) * percentage / 100.0)
	if filledWidth > width {
		filledWidth = width
	}
	if filledWidth < 1 && percentage > 0 {
		filledWidth = 1
	}

	// Create flowing animation within filled portion
	filled := ""
	for i := 0; i < filledWidth; i++ {
		// Animate only the fill, let tail position shift with tick
		offset := (i + a.Tick/2) % len(ProgressBarFrames)
		if i == filledWidth-1 {
			filled += lipgloss.NewStyle().Foreground(color).Bold(true).Render(ProgressBarFrames[offset])
		} else {
			filled += lipgloss.NewStyle().Foreground(color).Render("█")
		}
	}

	emptyWidth := width - filledWidth
	empty := lipgloss.NewStyle().Foreground(colors.Border).Render("░" + " "[0:1])
	if emptyWidth > 0 {
		empty = lipgloss.NewStyle().Foreground(colors.Border).Render(" "[0:1])
		for i := 0; i < emptyWidth; i++ {
			empty += "░"
		}
	}

	return filled + empty
}
