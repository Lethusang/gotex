package editor

import (
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func (e *Editor) createMainMenu() *fyne.MainMenu {
    fileMenu := fyne.NewMenu("File",
        fyne.NewMenuItem("New", e.handleNew),
        fyne.NewMenuItem("Open...", e.handleOpen),
        fyne.NewMenuItem("Save", e.handleSave),
        fyne.NewMenuItem("Save As...", e.handleSaveAs),
        fyne.NewMenuItemSeparator(),
        fyne.NewMenuItem("Exit", e.handleExit),
    )

    editMenu := fyne.NewMenu("Edit",
        fyne.NewMenuItem("Cut", e.handleCut),
        fyne.NewMenuItem("Copy", e.handleCopy),
        fyne.NewMenuItem("Paste", e.handlePaste),
    )

    viewMenu := fyne.NewMenu("View",
        fyne.NewMenuItem("Toggle Theme", e.toggleTheme),
    )

    return fyne.NewMainMenu(
        fileMenu,
        editMenu,
        viewMenu,
    )
}

// File menu handlers
func (e *Editor) handleNew() {
    if e.hasUnsavedChanges() {
        dialog.ShowConfirm("New Document",
            "Do you want to save changes?",
            func(save bool) {
                if save {
                    e.handleSave()
                }
                e.newDocument()
            }, e.mainWindow)
    } else {
        e.newDocument()
    }
}

func (e *Editor) handleOpen() {
    dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
        if err != nil || reader == nil {
            return
        }
        defer reader.Close()

        content, err := io.ReadAll(reader)
        if err != nil {
            dialog.ShowError(err, e.mainWindow)
            return
        }

        e.filename = reader.URI().Path()
        e.textArea.SetText(string(content))
        e.updateStatus("File opened")
    }, e.mainWindow)
}

func (e *Editor) handleSave() {
    if e.filename == "" {
        e.handleSaveAs()
        return
    }
    e.saveFile(e.filename)
}

func (e *Editor) handleSaveAs() {
    dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
        if err != nil || writer == nil {
            return
        }
        defer writer.Close()

        _, err = writer.Write([]byte(e.textArea.Text))
        if err != nil {
            dialog.ShowError(err, e.mainWindow)
            return
        }

        e.filename = writer.URI().Path()
        e.updateStatus("File saved")
    }, e.mainWindow)
}

func (e *Editor) handleExit() {
    if e.hasUnsavedChanges() {
        dialog.ShowConfirm("Exit",
            "Do you want to save changes before exiting?",
            func(save bool) {
                if save {
                    e.handleSave()
                }
                e.mainWindow.Close()
            }, e.mainWindow)
    } else {
        e.mainWindow.Close()
    }
}

// Edit menu handlers
func (e *Editor) handleCut() {
    e.textArea.TypedShortcut(&fyne.ShortcutCut{
        Clipboard: e.mainWindow.Clipboard(),
    })
}

func (e *Editor) handleCopy() {
    e.textArea.TypedShortcut(&fyne.ShortcutCopy{
        Clipboard: e.mainWindow.Clipboard(),
    })
}

func (e *Editor) handlePaste() {
    e.textArea.TypedShortcut(&fyne.ShortcutPaste{
        Clipboard: e.mainWindow.Clipboard(),
    })
}

// Helper functions
func (e *Editor) newDocument() {
    e.filename = ""
    e.textArea.SetText("")
    e.updateStatus("New document")
}

func (e *Editor) saveFile(filename string) {
    err := os.WriteFile(filename, []byte(e.textArea.Text), 0644)
    if err != nil {
        dialog.ShowError(err, e.mainWindow)
        return
    }
    e.filename = filename
    e.updateStatus("File saved")
}

func (e *Editor) hasUnsavedChanges() bool {
    // TODO: Implement real change tracking
    return e.textArea.Text != ""
}
