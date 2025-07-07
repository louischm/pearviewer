package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/louischm/pkg/logger"
)

var log = logger.NewLog()

func Run() {
	a := app.New()
	w := a.NewWindow("Pearviewer")
	w.Resize(fyne.NewSize(800, 600))
	a.Settings().SetTheme(&Theme{})

	login(w)

	w.ShowAndRun()
}
