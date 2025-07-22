package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type Theme struct {
}

var _ fyne.Theme = (*Theme)(nil)

func (t Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{
			R: 63,
			G: 63,
			B: 63,
			A: 100,
		}
	case theme.ColorNameScrollBarBackground:
		return color.White
	case theme.ColorNameButton:
		return color.White
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 211, G: 211, B: 211, A: 255}
	case theme.ColorNameDisabled:
		return color.Black
	case theme.ColorNameSuccess:
		return color.White
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	orig := theme.DefaultTheme().Icon(name)
	return theme.NewThemedResource(orig)
}

func (t Theme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t Theme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
