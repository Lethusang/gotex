package main

import (
	"fmt"
	"gotex/internal/buffer"
)

func main() {
    buf := buffer.NewBuffer()

    fmt.Print("\033[?1049h") // Use alternate screen buffer
    defer fmt.Print("\033[?1049l") // Restore main screen buffer

    for {
        // Clear screen
        fmt.Print("\033[2J")
        fmt.Print("\033[H")

        // Display buffer content
        content := buf.GetLines()
        for i, line := range content {
            fmt.Printf("%3d: %s\n", i+1, line)
        }

        // Display cursor position
        x, y := buf.GetCursor()
        fmt.Printf("\nCursor position: (%d, %d)\n", x, y)

        if buf.HasSelection() {
            startX, startY, endX, endY, _ := buf.GetSelectionCoordinates()
            fmt.Printf("Selection: (%d,%d) -> (%d,%d)\n", startX, startY, endX, endY)
            fmt.Printf("Selected text: %s\n", buf.GetSelection())
        }

        // Display menu
        fmt.Println("\nCommands:")
        fmt.Println("1: Insert text")
        fmt.Println("2: Move cursor")
        fmt.Println("3: Start selection")
        fmt.Println("4: End selection")
        fmt.Println("5: Delete selection")
        fmt.Println("q: Quit")

        // Get command
        var cmd string
        fmt.Print("\nEnter command: ")
        fmt.Scan(&cmd)

        switch cmd {
        case "1":
            fmt.Print("Enter text: ")
            var text string
            fmt.Scan(&text)
            for _, r := range text {
                buf.InsertRune(r)
            }

        case "2":
            fmt.Print("Enter dx dy: ")
            var dx, dy int
            fmt.Scan(&dx, &dy)
            buf.MoveCursor(dx, dy)

        case "3":
            buf.StartSelection()
            fmt.Println("Selection started")

        case "4":
            buf.EndSelection()
            fmt.Println("Selection ended")

        case "5":
            buf.DeleteSelection()
            fmt.Println("Selection deleted")

        case "q":
            return
        }
    }
}
