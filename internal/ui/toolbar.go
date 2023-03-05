package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MatusOllah/gcharted/internal/gcharted"
	"github.com/rs/zerolog/log"
)

func makeToolbar(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), func() {
			log.Info().Msg("selected toolbar item New")
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			log.Info().Msg("selected toolbar item Open")
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			log.Info().Msg("selected toolbar item Save")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			log.Info().Msg("selected toolbar item Play")
			gcharted.Play()
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			log.Info().Msg("selected toolbar item Pause")

			gcharted.IsPaused = !gcharted.IsPaused
			log.Info().Bool("IsPaused", gcharted.IsPaused).Msg("")

			gcharted.SetPaused(gcharted.IsPaused)
		}),
		widget.NewToolbarAction(theme.MediaFastRewindIcon(), func() {
			log.Info().Msg("selected toolbar item Rewind")
			if err := gcharted.Rewind(); err != nil {
				dialog.NewError(err, w).Show()
			}
		}),
		widget.NewToolbarAction(theme.MediaFastForwardIcon(), func() {
			log.Info().Msg("selected toolbar item Forward")
			if err := gcharted.Forward(); err != nil {
				dialog.NewError(err, w).Show()
			}
		}),
	), widget.NewSeparator())
}
