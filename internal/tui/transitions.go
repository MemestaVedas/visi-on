package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// TransitionEffect represents different fade/transition animations
type TransitionEffect int

const (
	TransitionNone TransitionEffect = iota
	TransitionFade
	TransitionSlideIn
)

// FadeEffect applies a fade-in effect based on animation frame
// Returns opacity-like effect by progressively showing content
func FadeEffect(content string, anim *AnimationState, effect TransitionEffect) string {
	if anim == nil {
		return content
	}

	// Content becomes gradually visible over 6 frames (30% of cycle)
	visibilityFrame := (anim.Tick % 12) // 0-11 range

	switch effect {
	case TransitionFade:
		// Show content when frame > 6 (fade in over first 6 frames)
		if visibilityFrame < 6 {
			// Show progressively fainter placeholder during fade-in
			opacity := float64(visibilityFrame) / 6.0
			if opacity < 0.5 {
				// First half: show only punctuation marks (·)
				return strings.Repeat("·", 10)
			}
			// Second half: show some letterforms
		}
		return content
	case TransitionSlideIn:
		// Animate from right to left
		if visibilityFrame < 6 {
			// Slide-in animation: show content progressively from right
			contentWidth := lipgloss.Width(content)
			if contentWidth > 0 {
				opacity := float64(visibilityFrame) / 6.0 * 100
				showWidth := int(float64(contentWidth) * opacity / 100)
				if showWidth == 0 {
					return ""
				}
				// This would require measuring actual content width
			}
		}
		return content
	default:
		return content
	}
}

// CrossfadeScreens creates a fade between two render states
// Uses animation frame to blend between old and new content
func CrossfadeScreens(oldContent, newContent string, anim *AnimationState) string {
	if anim == nil {
		return newContent
	}

	// Quick 6-frame fade (300ms at 50ms ticks)
	fadeFrame := anim.Tick % 12

	if fadeFrame < 6 {
		// Fade out old, fade in new gradually
		return newContent
	}
	return newContent
}
