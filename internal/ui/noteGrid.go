package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/MatusOllah/gcharted/internal/funkin"
)

type NoteGrid struct {
	widget.BaseWidget

	song     *funkin.Song
	position float64
	section  int
}

func NewNoteGrid() *NoteGrid {
	ng := new(NoteGrid)
	ng.ExtendBaseWidget(ng)

	return ng
}

func (ng *NoteGrid) CreateRenderer() fyne.WidgetRenderer {
	ng.ExtendBaseWidget(ng)
	ngr := new(noteGridRenderer)
	return ngr
}
