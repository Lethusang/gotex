// Package main provides a simple text editor demonstration using the buffer package for now
package main

import (
	"fmt"
	"log"

	"gotex/internal/buffer"
)

func main() {
	// Create a new buffer
	buf := buffer.New()

	// Example of basic text insertion
	insertText(buf, "Hello, World!")

	// Example of selection and manipulation
	demonstrateSelection(buf)

	// Print final buffer contents
	printBufferContents(buf)
}

// insertText inserts a string into the buffer character by character.
func insertText(buf *buffer.Buffer, text string) {
	for _, r := range text {
		if err := buf.InsertRune(r); err != nil {
			log.Printf("Error inserting rune: %v", err)
		}
	}
}

// demonstrateSelection shows selection functionality.
func demonstrateSelection(buf *buffer.Buffer) {
	// Move cursor to beginning
	if err := buf.MoveCursor(-buf.GetCursor().X, 0); err != nil {
		log.Printf("Error moving cursor to start: %v", err)
		return
	}

	// Start selection
	buf.StartSelection()

	// Select "Hello"
	if err := buf.MoveCursor(5, 0); err != nil {
		log.Printf("Error moving cursor during selection: %v", err)
		return
	}

	// Get selected text
	selected, err := buf.GetSelection()
	if err != nil {
		log.Printf("Error getting selection: %v", err)
		return
	}

	fmt.Printf("Selected text: %q\n", selected)

	// Replace selection with new text
	if err := buf.DeleteSelection(); err != nil {
		log.Printf("Error deleting selection: %v", err)
		return
	}

	insertText(buf, "Greetings")
}

// printBufferContents prints the current state of the buffer.
func printBufferContents(buf *buffer.Buffer) {
	fmt.Println("\nBuffer contents:")
	for i, line := range buf.GetLines() {
		fmt.Printf("Line %d: %q\n", i+1, line)
	}

	cursor := buf.GetCursor()
	fmt.Printf("\nCursor position: (%d, %d)\n", cursor.X, cursor.Y)
}

// Example usage of other buffer features
func additionalExamples(buf *buffer.Buffer) error {
	// Insert a new line
	if err := buf.NewLine(); err != nil {
		return fmt.Errorf("error inserting new line: %w", err)
	}

	// Insert some text on the new line
	insertText(buf, "This is a new line")

	// Move cursor around
	movements := [][2]int{
		{-1, 0}, // left
		{1, 0},  // right
		{0, -1}, // up
		{0, 1},  // down
	}

	for _, move := range movements {
		if err := buf.MoveCursor(move[0], move[1]); err != nil {
			log.Printf("Warning: couldn't move cursor (%d, %d): %v", move[0], move[1], err)
		}
	}

	return nil
}

// handleError is a helper function for error handling
func handleError(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}
