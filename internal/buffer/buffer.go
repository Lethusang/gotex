package buffer

import (
	"strings"
)

// Selection represents a selected region of text
type Selection struct {
	startX, startY int  // Starting position
	endX, endY     int  // Ending position
	active         bool // Whether selection is currently active
}

// Buffer represents the text content and cursor position
type Buffer struct {
	lines     []string
	cursorX   int
	cursorY   int
	filename  string
	selection Selection
}

func NewBuffer() *Buffer {
	return &Buffer{
		lines:    []string{""},
		cursorX:  0,
		cursorY:  0,
		filename: "",
		selection: Selection{
			active: false,
		},
	}
}

// GetLines returns all lines in the buffer
func (b *Buffer) GetLines() []string {
	return b.lines
}

// Other required methods
func (b *Buffer) GetCursor() (int, int) {
	return b.cursorX, b.cursorY
}

// StartSelection begins a new text selection at the current cursor position
func (b *Buffer) StartSelection() {
	b.selection.startX = b.cursorX
	b.selection.startY = b.cursorY
	b.selection.endX = b.cursorX
	b.selection.endY = b.cursorY
	b.selection.active = true
}

// EndSelection ends the current selection
func (b *Buffer) EndSelection() {
	b.selection.active = false
}

// UpdateSelection updates the selection end point to the current cursor position
func (b *Buffer) UpdateSelection() {
	if b.selection.active {
		b.selection.endX = b.cursorX
		b.selection.endY = b.cursorY
	}
}

// HasSelection returns whether there is an active selection
func (b *Buffer) HasSelection() bool {
	return b.selection.active
}

// GetSelection returns the selected text
func (b *Buffer) GetSelection() string {
	if !b.selection.active {
		return ""
	}

	// Normalize selection coordinates
	startX, startY, endX, endY := b.getNormalizedSelection()

	// Build selected text
	var result strings.Builder

	// Single line selection
	if startY == endY {
		line := b.lines[startY]
		result.WriteString(line[startX:endX])
		return result.String()
	}

	// Multi-line selection
	// First line
	result.WriteString(b.lines[startY][startX:])
	result.WriteString("\n")

	// Middle lines
	for y := startY + 1; y < endY; y++ {
		result.WriteString(b.lines[y])
		result.WriteString("\n")
	}

	// Last line
	result.WriteString(b.lines[endY][:endX])

	return result.String()
}

// NewLine inserts a new line at the current cursor position
func (b *Buffer) NewLine() {
	currentLine := b.lines[b.cursorY]
	remainingText := currentLine[b.cursorX:]
	b.lines[b.cursorY] = currentLine[:b.cursorX]

	// Insert new line
	b.lines = append(b.lines[:b.cursorY+1], append([]string{remainingText}, b.lines[b.cursorY+1:]...)...)

	b.cursorY++
	b.cursorX = 0
}

// InsertRune inserts a single character at the current cursor position
func (b *Buffer) InsertRune(r rune) {
	currentLine := b.lines[b.cursorY]
	newLine := currentLine[:b.cursorX] + string(r) + currentLine[b.cursorX:]
	b.lines[b.cursorY] = newLine
	b.cursorX++
}

// DeleteSelection removes the selected text
func (b *Buffer) DeleteSelection() {
	if !b.selection.active {
		return
	}

	startX, startY, endX, endY := b.getNormalizedSelection()

	// Single line deletion
	if startY == endY {
		line := b.lines[startY]
		b.lines[startY] = line[:startX] + line[endX:]
		b.cursorX = startX
		b.cursorY = startY
		b.EndSelection()
		return
	}

	// Multi-line deletion
	newLine := b.lines[startY][:startX] + b.lines[endY][endX:]
	b.lines = append(b.lines[:startY], append([]string{newLine}, b.lines[endY+1:]...)...)

	b.cursorX = startX
	b.cursorY = startY
	b.EndSelection()
}

// ReplaceSelection replaces the selected text with new text
func (b *Buffer) ReplaceSelection(newText string) {
	if !b.selection.active {
		return
	}

	startX, startY, _, _ := b.getNormalizedSelection() // Use blank identifier _ for unused values

	// Delete existing selection
	b.DeleteSelection()

	// Insert new text
	b.cursorX = startX
	b.cursorY = startY

	lines := strings.Split(newText, "\n")
	for i, line := range lines {
		if i > 0 {
			b.NewLine()
		}
		for _, r := range line {
			b.InsertRune(r)
		}
	}
}

// IsPositionInSelection checks if a given position is within the selection
func (b *Buffer) IsPositionInSelection(x, y int) bool {
	if !b.selection.active {
		return false
	}

	startX, startY, endX, endY := b.getNormalizedSelection()

	// If y is outside the selection range
	if y < startY || y > endY {
		return false
	}

	// If we're on the start line, check if x is after startX
	if y == startY && x < startX {
		return false
	}

	// If we're on the end line, check if x is before endX
	if y == endY && x >= endX {
		return false
	}

	// Position is within selection
	return true
}

// getNormalizedSelection returns selection coordinates in correct order
func (b *Buffer) getNormalizedSelection() (startX, startY, endX, endY int) {
	startX = b.selection.startX
	startY = b.selection.startY
	endX = b.selection.endX
	endY = b.selection.endY

	// Swap if selection is backwards
	if startY > endY || (startY == endY && startX > endX) {
		startX, endX = endX, startX
		startY, endY = endY, startY
	}

	return
}

// MoveCursor with selection support
func (b *Buffer) MoveCursor(deltaX, deltaY int) {
	newY := b.cursorY + deltaY

	if newY < 0 {
		newY = 0
	} else if newY >= len(b.lines) {
		newY = len(b.lines) - 1
	}

	newX := b.cursorX + deltaX
	if newX < 0 {
		if newY > 0 {
			newY--
			newX = len(b.lines[newY])
		} else {
			newX = 0
		}
	} else if newX > len(b.lines[newY]) {
		if newY < len(b.lines)-1 {
			newY++
			newX = 0
		} else {
			newX = len(b.lines[newY])
		}
	}

	b.cursorX = newX
	b.cursorY = newY

	// Update selection if active
	b.UpdateSelection()
}

// GetSelectionCoordinates returns the current selection coordinates
func (b *Buffer) GetSelectionCoordinates() (startX, startY, endX, endY int, active bool) {
	return b.selection.startX, b.selection.startY, b.selection.endX, b.selection.endY, b.selection.active
}
