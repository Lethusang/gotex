// cmd/gotex/main.go
package main

import (
	"gotex/internal/editor"
	"log"
)

func main() {
    ed := editor.NewEditor()
    if err := ed.Run(); err != nil {
        log.Fatalf("Error running editor: %v", err)
    }
}
