package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeSectionTab() *container.TabItem {
	return container.NewTabItem("Section", widget.NewLabel("horalky"))
}
