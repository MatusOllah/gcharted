package main

import (
	"log/slog"
	"os"

	"github.com/MatusOllah/gcharted/internal/gui"
	"github.com/MatusOllah/slogcolor"
	qt "github.com/mappu/miqt/qt6"
)

func main() {
	// Logger
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))

	qt.NewQApplication(os.Args)

	gui.NewMainWindowUi().MainWindow.Show()

	qt.QApplication_Exec()
}
