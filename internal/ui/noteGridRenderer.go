package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type noteGridRenderer struct {
	arrowLeft  *canvas.Image
	arrowDown  *canvas.Image
	arrowUp    *canvas.Image
	arrowRight *canvas.Image

	bg *canvas.Rectangle

	noteGrid *NoteGrid
}

func (ngr *noteGridRenderer) Destroy() {
	return
}

func (ngr *noteGridRenderer) Layout(fyne.Size) {
	return
}

func (ngr *noteGridRenderer) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (ngr *noteGridRenderer) Objects() []fyne.CanvasObject {
	return nil
}

func (ngr *noteGridRenderer) Refresh() {
	return
}
