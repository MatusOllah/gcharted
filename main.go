package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	version = "1.0.0"
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

	log.Info().Msgf("gcharted version %s", version)
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

	w.SetMainMenu(makeMenu(a, w))

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()

	log.Info().Msg("exiting")
	runtime.GC()
	os.Exit(0)
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	file := fyne.NewMenu(
		"File",
		fyne.NewMenuItem("New", func() {
			log.Debug().Msg("selected menu item File>New")
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Open", func() {
			log.Debug().Msg("selected menu item File>Open")
		}),
	)

	help := fyne.NewMenu(
		"Help",
		fyne.NewMenuItem("About", func() {
			log.Debug().Msg("selected menu item Help>About")
			dialog.NewInformation("About gcharted", fmt.Sprintf("gcharted version %s\nGo version %s", version, runtime.Version()), w).Show()
		}),
	)

	main := fyne.NewMainMenu(
		file,
		help,
	)
	return main
}
