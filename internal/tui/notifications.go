package tui

import (
	"fmt"
	"time"
)

// NotificationType defines the type of notification
type NotificationType int

const (
	NotifySuccess NotificationType = iota
	NotifyError
	NotifyInfo
	NotifyWarning
)

// Notification represents a transient toast message
type Notification struct {
	Type      NotificationType
	Message   string
	CreatedAt time.Time
	Duration  time.Duration
}

// NotificationQueue manages active notifications
type NotificationQueue struct {
	items    []Notification
	maxItems int
}

// NewNotificationQueue creates a new notification queue
func NewNotificationQueue(maxItems int) *NotificationQueue {
	if maxItems <= 0 {
		maxItems = 5
	}
	return &NotificationQueue{
		items:    make([]Notification, 0, maxItems),
		maxItems: maxItems,
	}
}

// Add adds a notification to the queue
func (nq *NotificationQueue) Add(notifType NotificationType, message string) {
	notification := Notification{
		Type:      notifType,
		Message:   message,
		CreatedAt: time.Now(),
		Duration:  3 * time.Second,
	}

	nq.items = append(nq.items, notification)

	// Keep only the latest maxItems
	if len(nq.items) > nq.maxItems {
		nq.items = nq.items[len(nq.items)-nq.maxItems:]
	}
}

// Success adds a success notification
func (nq *NotificationQueue) Success(message string) {
	nq.Add(NotifySuccess, message)
}

// Error adds an error notification
func (nq *NotificationQueue) Error(message string) {
	nq.Add(NotifyError, message)
}

// Info adds an info notification
func (nq *NotificationQueue) Info(message string) {
	nq.Add(NotifyInfo, message)
}

// Warning adds a warning notification
func (nq *NotificationQueue) Warning(message string) {
	nq.Add(NotifyWarning, message)
}

// GetActive returns active notifications that haven't expired
func (nq *NotificationQueue) GetActive() []Notification {
	var active []Notification
	now := time.Now()

	for _, notif := range nq.items {
		elapsed := now.Sub(notif.CreatedAt)
		if elapsed < notif.Duration {
			active = append(active, notif)
		}
	}

	return active
}

// Clear removes all notifications
func (nq *NotificationQueue) Clear() {
	nq.items = nq.items[:0]
}

// Render renders a notification as a styled string
func (n Notification) Render(anim *AnimationState) string {
	icon := "✓"
	style := ToastSuccessStyle

	switch n.Type {
	case NotifyError:
		icon = "✗"
		style = ToastErrorStyle
	case NotifyWarning:
		icon = "⚠"
		style = ToastInfoStyle
	case NotifyInfo:
		icon = "ℹ"
		style = ToastInfoStyle
	}

	// Add fade-out effect as notification ages
	elapsed := time.Since(n.CreatedAt)
	if elapsed > n.Duration*4/5 {
		// Last 20% of duration: show fading
		remaining := n.Duration - elapsed
		pct := int(100 * remaining / (n.Duration / 5))
		if pct > 0 && pct < 100 {
			// Could add opacity here if lipgloss supports it
			// For now, just show the notification normally
		}
	}

	message := fmt.Sprintf(" %s %s", icon, n.Message)
	return style.Render(message)
}
