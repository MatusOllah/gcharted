package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/MatusOllah/slogcolor"
	qt "github.com/mappu/miqt/qt6"
)

func main() {
	// Logger
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))

	qt.NewQApplication(os.Args)

	btn := qt.NewQPushButton3("Hello world!")
	btn.SetFixedWidth(320)

	var counter int = 0

	btn.OnPressed(func() {
		counter++
		btn.SetText(fmt.Sprintf("You have clicked the button %d time(s)", counter))
	})

	btn.Show()

	qt.QApplication_Exec()
}
