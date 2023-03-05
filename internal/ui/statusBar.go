package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/MatusOllah/gcharted/internal/gcharted"
)

func makeStatusBar() fyne.CanvasObject {
	gcharted.PositionBinding.Set("Position: ")

	return container.NewVBox(widget.NewSeparator(), container.NewHBox(
		widget.NewLabel(gcharted.Version),
		widget.NewSeparator(),
		widget.NewLabelWithData(gcharted.PositionBinding),
		widget.NewSeparator(),
		widget.NewLabel("Section: 0"),
		widget.NewSeparator(),
		widget.NewLabel("CurStep: 0"),
	))
}
