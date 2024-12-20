package main

import (
	"log/slog"
	"os"
	"runtime"

	_ "github.com/MatusOllah/gcharted/assets"
	"github.com/MatusOllah/gcharted/internal/gui"
	"github.com/MatusOllah/slogcolor"
	qt "github.com/mappu/miqt/qt6"
)

func main() {
	// Logger
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))

	slog.Info("GCharted version " + Version)
	slog.Info("Go version " + runtime.Version())
	slog.Info("Qt version " + qt.QLibraryInfo_Version().ToString())

	// Qt Application
	slog.Info("initializing QApplication")
	qt.NewQApplication(os.Args)

	qt.QGuiApplication_SetApplicationDisplayName("GCharted")
	qt.QCoreApplication_SetApplicationName("GCharted")
	qt.QCoreApplication_SetApplicationVersion(Version)

	slog.Info("showing GUI")
	gui.NewMainWindow().Ui().MainWindow.Show()

	slog.Info("executing QApplication")
	code := qt.QApplication_Exec()

	os.Exit(code)
}
