package editor

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// EditorTheme implements a custom theme for our editor
type EditorTheme struct {
	defaultTheme fyne.Theme
}

// NewEditorTheme creates a new instance of our custom theme
func NewEditorTheme() fyne.Theme {
	return &EditorTheme{
		defaultTheme: theme.DefaultTheme(),
	}
}

// Color overrides specific colors from the default theme
func (t *EditorTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		if variant == theme.VariantDark {
			return colorBackgroundDark
		}
		return colorBackground

	case theme.ColorNameForeground:
		if variant == theme.VariantDark {
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
	if e.app.Settings().ThemeVariant() == theme.VariantDark {
		e.app.Settings().SetTheme(theme.LightTheme())
	} else {
		e.app.Settings().SetTheme(theme.DarkTheme())
	}
}
