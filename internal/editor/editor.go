package editor

import (
	"fmt"

	"gotex/internal/buffer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Editor struct {
	app        fyne.App
	mainWindow fyne.Window
	textArea   *widget.Entry
	statusBar  *widget.Label
	buffer     *buffer.Buffer
	filename   string
}

func NewEditor() *Editor {
	e := &Editor{
		app:    app.New(),
		buffer: buffer.New(),
	}

	e.applyTheme()
	e.mainWindow = e.app.NewWindow("GoTex Editor")
	e.initUI()

	return e
}

func (e *Editor) Run() error {
	e.mainWindow.ShowAndRun()
	return nil
}

func (e *Editor) initUI() {
	// Initialize text area with handlers
	e.textArea = widget.NewMultiLineEntry()
	e.textArea.OnChanged = e.handleTextChanged
	e.textArea.Wrapping = fyne.TextWrapWord
	e.textArea.MultiLine = true
	e.textArea.SetPlaceHolder("Type your text here...") // Add this to see something initially

	// Initialize status bar
	e.statusBar = widget.NewLabel("Ready")

	// Create toolbar
	toolbar := e.createToolbar()

	// Create main container
	content := container.NewBorder(
		toolbar,                         // top
		e.createStatusBar(),             // bottom
		nil,                             // left
		nil,                             // right
		container.NewScroll(e.textArea), // center
	)

	// Set window content
	e.mainWindow.SetContent(content)
	e.mainWindow.SetMainMenu(e.createMainMenu())

	// Set a reasonable default size
	e.mainWindow.Resize(fyne.NewSize(800, 600))

	// Set window title
	e.mainWindow.SetTitle("GoTex Editor")

	// Center the window on screen
	e.mainWindow.CenterOnScreen()
}

func (e *Editor) handleTextChanged(content string) {
	// Update status bar with character count
	e.updateStatus(fmt.Sprintf("Characters: %d", len(content)))

	// Update buffer
	e.buffer = buffer.New()
	for _, r := range content {
		e.buffer.InsertRune(r)
	}
}

func (e *Editor) updateStatus(status string) {
	if e.filename != "" {
		status += " | " + e.filename
	}
	e.statusBar.SetText(status)
}

func (e *Editor) setupShortcuts() {
	// Add common keyboard shortcuts
	e.mainWindow.Canvas().AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyS,
		Modifier: fyne.KeyModifierControl, // Updated modifier
	}, func(shortcut fyne.Shortcut) {
		e.handleSave()
	})

	e.mainWindow.Canvas().AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyO,
		Modifier: fyne.KeyModifierControl, // Updated modifier
	}, func(shortcut fyne.Shortcut) {
		e.handleOpen()
	})
}

func (e *Editor) createStatusBar() fyne.CanvasObject {
	return container.NewHBox(e.statusBar)
}

func (e *Editor) applyTheme() {
	e.app.Settings().SetTheme(NewEditorTheme())
}
