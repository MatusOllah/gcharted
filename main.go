package main

import (
	"os"
	"runtime"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/MatusOllah/gcharted/internal/gcharted"
	"github.com/MatusOllah/gcharted/internal/ui"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).With().Caller().Logger()

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	log.Info().Msgf("gcharted version %s", gcharted.Version)
	log.Info().Msgf("Go version %s", runtime.Version())
	log.Info().Msg("ahoj!")

	a := app.New()

	a.Lifecycle().SetOnStarted(func() {
		log.Info().Msg("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Info().Msg("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Info().Msg("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Info().Msg("Lifecycle: Exited Foreground")
	})

	w := a.NewWindow("gcharted")
	w.SetMaster()
	w.Resize(fyne.NewSize(1280, 720))
	w.SetMainMenu(ui.MakeMenu(w))
	w.SetContent(ui.MakeUI(w))

	w.ShowAndRun()

	log.Info().Msg("exiting")
	runtime.GC()
	os.Exit(0)
}
