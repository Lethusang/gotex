// Package statusbar provides status bar functionality
package editor

import (
	"fyne.io/fyne/v2/widget"
)

// StatusBar represents the editor's status bar
type StatusBar struct {
	label *widget.Label
}

// New creates a new status bar
func NewStatusBar() *StatusBar {
	return &StatusBar{
		label: widget.NewLabel("Ready"),
	}
}

// Update updates the status bar text
func (s *StatusBar) Update(text string) {
	s.label.SetText(text)
}
