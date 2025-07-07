package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// EntryWrapper impose une taille minimale au widget enfant
type EntryWrapper struct {
	widget.BaseWidget
	entry *widget.Entry
	size  fyne.Size
}

func NewEntryWrapper(e *widget.Entry, size fyne.Size) *EntryWrapper {
	w := &EntryWrapper{
		entry: e,
		size:  size,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *EntryWrapper) CreateRenderer() fyne.WidgetRenderer {
	w.entry.Resize(w.size)
	return widget.NewSimpleRenderer(w.entry)
}

func (w *EntryWrapper) MinSize() fyne.Size {
	return w.size
}
