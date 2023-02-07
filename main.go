package main

import (
	"os"
	"runtime"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
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
	w := a.NewWindow("gcharted")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
