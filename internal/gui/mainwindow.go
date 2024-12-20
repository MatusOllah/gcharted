package gui

import (
	"log/slog"

	qt "github.com/mappu/miqt/qt6"
)

type MainWindow struct {
	ui *MainWindowUi
}

func NewMainWindow() *MainWindow {
	w := &MainWindow{ui: NewMainWindowUi()}

	// Connections
	w.ui.actionExit.OnTriggered(w.on_actionExit_triggered)
	w.ui.actionAbout.OnTriggered(w.on_actionAbout_triggered)
	w.ui.actionAbout_Qt.OnTriggered(w.on_actionAbout_Qt_triggered)

	return w
}

func (w *MainWindow) Ui() *MainWindowUi {
	return w.ui
}

func (w *MainWindow) on_actionExit_triggered() {
	slog.Info("triggered exit button, exiting")
	qt.QCoreApplication_Exit()
}

func (w *MainWindow) on_actionAbout_triggered() {
	slog.Info("triggered about button")
	qt.QMessageBox_About(
		w.ui.centralwidget,
		qt.QGuiApplication_ApplicationDisplayName(),
		"<b> GCharted "+qt.QCoreApplication_ApplicationVersion()+"</b><br>The Go + Qt based FNF Chart Editor<br><br><a href='https://github.com/MatusOllah/gcharted'>GCharted GitHub repository</a>",
	)
}

func (w *MainWindow) on_actionAbout_Qt_triggered() {
	slog.Info("triggered about Qt button")
	qt.QMessageBox_AboutQt(w.ui.centralwidget)
}
