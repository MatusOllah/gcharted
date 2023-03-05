package gcharted

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/rs/zerolog/log"
)

func OpenInst(w fyne.Window) {
	openDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.NewError(err, w).Show()
		}

		log.Info().Msgf("inst path: %s", uc.URI().Path())

		if err := InstPathBinding.Set(uc.URI().Name()); err != nil {
			dialog.NewError(err, w).Show()
		}

		log.Info().Msgf("opening inst file %s", uc.URI().Path())
		file, err := os.Open(uc.URI().Path())
		if err != nil {
			dialog.NewError(err, w).Show()
		}

		log.Info().Msg("decoding inst")
		streamer, format, err := vorbis.Decode(file)
		if err != nil {
			dialog.NewError(err, w).Show()
		}

		InstStreamer = streamer
		InstCtrl = &beep.Ctrl{
			Streamer: InstStreamer,
			Paused:   false,
		}
		InstVolume = &effects.Volume{
			Streamer: InstCtrl,
			Base:     2,
			Volume:   0,
			Silent:   false,
		}
		InstFormat = format

		go func() {
			for {
				UpdatePosition()
			}
		}()
	}, w)
	openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".ogg"}))
	openDialog.Show()
}

func OpenVocals(w fyne.Window) {
	openDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.NewError(err, w).Show()
		}

		log.Info().Msgf("vocals path: %s", uc.URI().Path())

		if err := VocalsPathBinding.Set(uc.URI().Name()); err != nil {
			dialog.NewError(err, w).Show()
		}

		log.Info().Msgf("opening vocals file %s", uc.URI().Path())
		file, err := os.Open(uc.URI().Path())
		if err != nil {
			dialog.NewError(err, w).Show()
		}

		log.Info().Msg("decoding vocals")
		streamer, _, err := vorbis.Decode(file)
		if err != nil {
			dialog.NewError(err, w).Show()
		}

		VocalsStreamer = streamer
		VocalsCtrl = &beep.Ctrl{
			Streamer: VocalsStreamer,
			Paused:   false,
		}
		VocalsVolume = &effects.Volume{
			Streamer: VocalsCtrl,
			Base:     2,
			Volume:   0,
			Silent:   false,
		}
	}, w)
	openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".ogg"}))
	openDialog.Show()
}

func SetInstMuted(muted bool) {
	speaker.Lock()
	InstVolume.Silent = muted
	speaker.Unlock()
}

func SetVocalsMuted(muted bool) {
	speaker.Lock()
	VocalsVolume.Silent = muted
	speaker.Unlock()
}
