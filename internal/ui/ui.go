package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func MakeUI(w fyne.Window) fyne.CanvasObject {
	rightAppTabs := container.NewAppTabs(
		makeSectionTab(),
		makeSongTab(),
		makeTracksTab(),
	)

	return container.NewBorder(
		makeToolbar(w),
		makeStatusBar(),
		nil,
		nil,
		container.NewHSplit(container.NewDocTabs(), rightAppTabs),
	)
}
