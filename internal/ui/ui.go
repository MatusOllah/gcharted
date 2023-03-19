package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/MatusOllah/gcharted/internal/gcharted"
	"github.com/MatusOllah/gcharted/internal/notegrid"
)

func MakeUI(w fyne.Window) fyne.CanvasObject {
	rightAppTabs := container.NewAppTabs(
		makeSectionTab(),
		makeSongTab(),
		makeTracksTab(),
	)

	ng := notegrid.New()
	gcharted.NoteGrid = ng

	left := container.NewPadded(container.NewMax(ng))

	return container.NewBorder(
		makeToolbar(w),
		makeStatusBar(),
		nil,
		nil,
		container.NewHSplit(left, rightAppTabs),
	)
}
