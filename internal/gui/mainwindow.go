package gui

import (
	"log/slog"

	"github.com/MatusOllah/gcharted/i18n"
	qt "github.com/mappu/miqt/qt6"
)

type MainWindow struct {
	ui *MainWindowUi
}

func NewMainWindow() *MainWindow {
	w := &MainWindow{ui: NewMainWindowUi()}
	w.Retranslate()

	// Connections
	w.ui.actionExit.OnTriggered(w.on_actionExit_triggered)
	w.ui.actionAbout.OnTriggered(w.on_actionAbout_triggered)
	w.ui.actionAbout_Qt.OnTriggered(w.on_actionAbout_Qt_triggered)

	return w
}

func (w *MainWindow) Ui() *MainWindowUi {
	return w.ui
}

// Retranslate reapplies all text translations.
func (w *MainWindow) Retranslate() {
	w.ui.actionAbout.SetText(i18n.L("About"))
	w.ui.actionAbout_Qt.SetText(i18n.L("AboutQt"))
	w.ui.actionExit.SetText(i18n.L("Exit"))
	w.ui.menuFile.SetTitle(i18n.L("File"))
	w.ui.menuHelp.SetTitle(i18n.L("Help"))
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
	) // TODO: probably should translate this also ^
}

func (w *MainWindow) on_actionAbout_Qt_triggered() {
	slog.Info("triggered about Qt button")
	qt.QMessageBox_AboutQt2(w.ui.centralwidget, i18n.L("AboutQt"))
}
