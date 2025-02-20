package editor

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// EditorTheme implements a custom theme for our editor
type EditorTheme struct {
	defaultTheme fyne.Theme
	isDark       bool
}

// NewEditorTheme creates a new instance of our custom theme
func NewEditorTheme() *EditorTheme { // Changed return type to *EditorTheme
	return &EditorTheme{
		defaultTheme: theme.DefaultTheme(),
		isDark:       false,
	}
}

// Add ThemeVariant method to implement the Theme interface
func (t *EditorTheme) ThemeVariant() fyne.ThemeVariant {
	if t.isDark {
		return theme.VariantDark
	}
	return theme.VariantLight
}

// Color overrides specific colors from the default theme
func (t *EditorTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	// Use the theme's own dark/light state instead of the variant parameter
	switch name {
	case theme.ColorNameBackground:
		if t.isDark {
			return colorBackgroundDark
		}
		return colorBackground

	case theme.ColorNameForeground:
		if t.isDark {
			return colorForegroundDark
		}
		return colorForeground

	case theme.ColorNamePrimary:
		return colorPrimary

	case theme.ColorNameSelection:
		return colorSelection

	case theme.ColorNameFocus:
		return colorPrimary
	}

	// Use the theme's state for the default theme colors
	variant := theme.VariantLight
	if t.isDark {
		variant = theme.VariantDark
	}
	return t.defaultTheme.Color(name, variant)
}

// Icon overrides specific icons from the default theme
func (t *EditorTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return t.defaultTheme.Icon(name)
}

// Font overrides the default font
func (t *EditorTheme) Font(style fyne.TextStyle) fyne.Resource {
	return t.defaultTheme.Font(style)
}

// Size overrides specific sizes from the default theme
func (t *EditorTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 12 // Default text size
	case theme.SizeNamePadding:
		return 4 // Default padding
	case theme.SizeNameInnerPadding:
		return 2 // Inner padding
	case theme.SizeNameScrollBar:
		return 12 // Scrollbar width
	case theme.SizeNameScrollBarSmall:
		return 3 // Small scrollbar width
	}
	return t.defaultTheme.Size(name)
}

// Constants for theme colors
var (
	// Light theme colors
	colorBackground = color.NRGBA{R: 252, G: 252, B: 252, A: 255}
	colorForeground = color.NRGBA{R: 32, G: 32, B: 32, A: 255}
	colorPrimary    = color.NRGBA{R: 51, G: 153, B: 255, A: 255}
	colorSelection  = color.NRGBA{R: 51, G: 153, B: 255, A: 100}

	// Dark theme colors
	colorBackgroundDark = color.NRGBA{R: 32, G: 32, B: 32, A: 255}
	colorForegroundDark = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
)

func (e *Editor) toggleTheme() {
	currentTheme, ok := e.app.Settings().Theme().(*EditorTheme)
	if !ok {
		// If not using our theme, create new one
		currentTheme = NewEditorTheme()
	}

	// Toggle the theme
	currentTheme.isDark = !currentTheme.isDark

	// Apply the theme
	e.app.Settings().SetTheme(currentTheme)

	// Update status
	if currentTheme.isDark {
		e.updateStatus("Switched to Dark Theme")
	} else {
		e.updateStatus("Switched to Light Theme")
	}
}
