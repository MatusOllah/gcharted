package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"
)

func makeToolbar(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), func() {
			log.Info().Msg("selected toolbar item New")
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			log.Info().Msg("selected toolbar item Open")
			dialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {

			}, w)

			dialog.Show()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			log.Info().Msg("selected toolbar item Save")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), func() {}),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {}),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), func() {}),
	), widget.NewSeparator())
}
