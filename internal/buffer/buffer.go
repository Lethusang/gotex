// Package buffer provides text buffer manipulation capabilities with support
// for cursor positioning, text selection, and basic editing operations.
package buffer

import (
	"errors"
	"strings"
)

// Position represents a location in the buffer defined by X (column) and Y (row).
type Position struct {
    X, Y int
}

// Selection represents a selected region of text defined by start and end positions.
type Selection struct {
    start, end Position
}

// Buffer represents an editable text buffer with cursor positioning and selection support.
// It maintains the text content as a slice of lines and provides methods for text manipulation.
type Buffer struct {
    lines     []string   // Text content split by lines
    cursor    Position   // Current cursor position
    selection *Selection // Current selection, nil if none
    filename  string     // Associated filename, empty if new buffer
}

// Common errors that can occur during buffer operations.
var (
    ErrInvalidPosition  = errors.New("invalid cursor position")
    ErrInvalidSelection = errors.New("invalid selection bounds")
    ErrNoSelection     = errors.New("no active selection")
)

// New creates a new empty buffer.
// It initializes the buffer with a single empty line and cursor at position (0,0).
func New() *Buffer {
    return &Buffer{
        lines:  []string{""},
        cursor: Position{},
    }
}

// GetLines returns all lines in the buffer.
func (b *Buffer) GetLines() []string {
    return b.lines
}

// GetCursor returns the current cursor position.
func (b *Buffer) GetCursor() Position {
    return b.cursor
}

// StartSelection begins a new text selection at the current cursor position.
func (b *Buffer) StartSelection() {
    b.selection = &Selection{
        start: b.cursor,
        end:   b.cursor,
    }
}

// EndSelection clears the current selection.
func (b *Buffer) EndSelection() {
    b.selection = nil
}

// HasSelection returns whether there is an active selection.
func (b *Buffer) HasSelection() bool {
    return b.selection != nil
}

// UpdateSelection updates the selection end point to the current cursor position.
func (b *Buffer) UpdateSelection() {
    if b.HasSelection() {
        b.selection.end = b.cursor
    }
}

// GetSelection returns the selected text.
// Returns an empty string and error if no selection exists.
func (b *Buffer) GetSelection() (string, error) {
    if !b.HasSelection() {
        return "", ErrNoSelection
    }

    start, end := b.getNormalizedSelectionBounds()
    return b.getTextBetween(start, end), nil
}

// NewLine inserts a new line at the current cursor position.
func (b *Buffer) NewLine() error {
    if err := b.validateCursor(); err != nil {
        return err
    }

    currentLine := b.lines[b.cursor.Y]
    remainingText := currentLine[b.cursor.X:]
    b.lines[b.cursor.Y] = currentLine[:b.cursor.X]

    // Insert new line
    b.lines = append(
        b.lines[:b.cursor.Y+1],
        append([]string{remainingText}, b.lines[b.cursor.Y+1:]...)...,
    )

    b.cursor.Y++
    b.cursor.X = 0
    return nil
}

// InsertRune inserts a single character at the current cursor position.
func (b *Buffer) InsertRune(r rune) error {
    if err := b.validateCursor(); err != nil {
        return err
    }

    currentLine := b.lines[b.cursor.Y]
    if b.cursor.X > len(currentLine) {
        return ErrInvalidPosition
    }

    newLine := currentLine[:b.cursor.X] + string(r) + currentLine[b.cursor.X:]
    b.lines[b.cursor.Y] = newLine
    b.cursor.X++
    return nil
}

// DeleteSelection removes the selected text.
// Returns an error if no selection exists.
func (b *Buffer) DeleteSelection() error {
    if !b.HasSelection() {
        return ErrNoSelection
    }

    start, end := b.getNormalizedSelectionBounds()

    // Single line deletion
    if start.Y == end.Y {
        line := b.lines[start.Y]
        b.lines[start.Y] = line[:start.X] + line[end.X:]
        b.cursor = start
        b.EndSelection()
        return nil
    }

    // Multi-line deletion
    newLine := b.lines[start.Y][:start.X] + b.lines[end.Y][end.X:]
    b.lines = append(b.lines[:start.Y], append([]string{newLine}, b.lines[end.Y+1:]...)...)

    b.cursor = start
    b.EndSelection()
    return nil
}

// MoveCursor moves the cursor by the specified delta.
// Returns an error if the resulting position would be invalid.
func (b *Buffer) MoveCursor(deltaX, deltaY int) error {
    newPos := Position{
        X: b.cursor.X + deltaX,
        Y: b.cursor.Y + deltaY,
    }

    if err := b.validatePosition(newPos); err != nil {
        return err
    }

    b.cursor = newPos
    b.UpdateSelection()
    return nil
}

// IsPositionInSelection checks if a given position is within the selection.
func (b *Buffer) IsPositionInSelection(pos Position) bool {
    if !b.HasSelection() {
        return false
    }

    start, end := b.getNormalizedSelectionBounds()
    return isPositionBetween(pos, start, end)
}

// Private helper methods

// validateCursor checks if the current cursor position is valid.
func (b *Buffer) validateCursor() error {
    return b.validatePosition(b.cursor)
}

// validatePosition checks if the given position is valid within the buffer.
func (b *Buffer) validatePosition(pos Position) error {
    if pos.Y < 0 || pos.Y >= len(b.lines) {
        return ErrInvalidPosition
    }
    if pos.X < 0 || pos.X > len(b.lines[pos.Y]) {
        return ErrInvalidPosition
    }
    return nil
}

// getNormalizedSelectionBounds returns selection bounds in correct order.
func (b *Buffer) getNormalizedSelectionBounds() (start, end Position) {
    start = b.selection.start
    end = b.selection.end

    // Swap if selection is backwards
    if start.Y > end.Y || (start.Y == end.Y && start.X > end.X) {
        start, end = end, start
    }

    return start, end
}

// getTextBetween returns the text between two positions.
func (b *Buffer) getTextBetween(start, end Position) string {
    var result strings.Builder

    // Single line selection
    if start.Y == end.Y {
        result.WriteString(b.lines[start.Y][start.X:end.X])
        return result.String()
    }

    // Multi-line selection
    result.WriteString(b.lines[start.Y][start.X:])
    result.WriteString("\n")

    // Middle lines
    for y := start.Y + 1; y < end.Y; y++ {
        result.WriteString(b.lines[y])
        result.WriteString("\n")
    }

    // Last line
    result.WriteString(b.lines[end.Y][:end.X])

    return result.String()
}

// isPositionBetween checks if a position is between start and end positions.
func isPositionBetween(pos, start, end Position) bool {
    if pos.Y < start.Y || pos.Y > end.Y {
        return false
    }
    if pos.Y == start.Y && pos.X < start.X {
        return false
    }
    if pos.Y == end.Y && pos.X >= end.X {
        return false
    }
    return true
}
