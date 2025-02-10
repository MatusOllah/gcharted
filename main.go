package main

import (
	"log/slog"
	"os"
	"runtime"

	_ "github.com/MatusOllah/gcharted/assets"
	"github.com/MatusOllah/gcharted/internal/gui"
	"github.com/MatusOllah/gcharted/internal/i18n"
	"github.com/MatusOllah/slogcolor"
	qt "github.com/mappu/miqt/qt6"
)

func main() {
	// Logger
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, nil)))

	slog.Info("GCharted version " + Version)
	slog.Info("Go version " + runtime.Version())
	slog.Info("Qt version " + qt.QLibraryInfo_Version().ToString())

	slog.Info("loading i18n")
	if err := i18n.Init("sk"); err != nil { //TODO: get system locale or from some config
		panic(err)
	}

	// Qt Application
	slog.Info("initializing QApplication")
	qt.NewQApplication(os.Args)

	qt.QGuiApplication_SetApplicationDisplayName("GCharted")
	qt.QCoreApplication_SetApplicationName("GCharted")
	qt.QCoreApplication_SetApplicationVersion(Version)

	slog.Info("showing GUI")
	gui.NewMainWindow().Ui().MainWindow.Show()

	slog.Info("executing QApplication")
	qt.QApplication_Exec()
}
