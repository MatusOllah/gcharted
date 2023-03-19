package gcharted

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/MatusOllah/gcharted/internal/funkin"
	"github.com/rs/zerolog/log"
	"github.com/ztrue/tracerr"
)

func OpenChart(w fyne.Window) {
	openDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil {
			tracerr.Print(err)
			dialog.NewError(err, w).Show()
			return
		}

		log.Info().Msgf("chart path: %s", uc.URI().Path())

		log.Info().Msgf("reading chart file %s", uc.URI().Path())
		content, err := os.ReadFile(uc.URI().Path())
		if err != nil {
			tracerr.Print(err)
			dialog.NewError(err, w).Show()
			return
		}

		log.Info().Msg("decoding chart file")
		song, err := funkin.LoadSongFromJSON(content)
		if err != nil {
			tracerr.Print(err)
			dialog.NewError(err, w).Show()
			return
		}

		Song = song
		NoteGrid.SetSong(Song)
	}, w)
	openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
	openDialog.Show()
}
