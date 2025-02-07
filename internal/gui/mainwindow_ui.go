// Generated by miqt-uic. To update this file, edit the .ui file in
// Qt Designer, and then run 'go generate'.
//
//go:generate go run ../../scripts/uic/uic.go "-InFile" "mainwindow.ui" "-OutFile" "mainwindow_ui.go" "-Package" "gui"

package gui

import (
	qt "github.com/mappu/miqt/qt6"
)

type MainWindowUi struct {
	MainWindow     *qt.QMainWindow
	centralwidget  *qt.QWidget
	menubar        *qt.QMenuBar
	menuFile       *qt.QMenu
	menuHelp       *qt.QMenu
	toolBar        *qt.QToolBar
	statusbar      *qt.QStatusBar
	actionAbout    *qt.QAction
	actionAbout_Qt *qt.QAction
	actionExit     *qt.QAction
}

// NewMainWindowUi creates all Qt widget classes for MainWindow.
func NewMainWindowUi() *MainWindowUi {
	ui := &MainWindowUi{}

	ui.MainWindow = qt.NewQMainWindow(nil)
	ui.MainWindow.SetObjectName(*qt.NewQAnyStringView3("MainWindow"))
	ui.MainWindow.Resize(1280, 720)
	ui.MainWindow.SetWindowTitle("GCharted")
	icon0 := qt.NewQIcon()
	icon0.AddFile4(":/icons/gcharted.png", qt.NewQSize(), qt.QIcon__Normal, qt.QIcon__Off)
	ui.MainWindow.SetWindowIcon(icon0)

	ui.actionAbout = qt.NewQAction()
	ui.actionAbout.SetObjectName(*qt.NewQAnyStringView3("actionAbout"))

	ui.actionAbout_Qt = qt.NewQAction()
	ui.actionAbout_Qt.SetObjectName(*qt.NewQAnyStringView3("actionAbout_Qt"))

	ui.actionExit = qt.NewQAction()
	ui.actionExit.SetObjectName(*qt.NewQAnyStringView3("actionExit"))

	ui.centralwidget = qt.NewQWidget(ui.MainWindow.QWidget)
	ui.centralwidget.SetObjectName(*qt.NewQAnyStringView3("centralwidget"))
	ui.MainWindow.SetCentralWidget(ui.centralwidget) // Set central widget

	ui.menubar = qt.NewQMenuBar(ui.MainWindow.QWidget)
	ui.menubar.SetObjectName(*qt.NewQAnyStringView3("menubar"))
	ui.menubar.Resize(1280, 33)

	ui.menuFile = qt.NewQMenu(ui.menubar.QWidget)
	ui.menuFile.SetObjectName(*qt.NewQAnyStringView3("menuFile"))
	ui.menuFile.QWidget.AddAction(ui.actionExit)

	ui.menuHelp = qt.NewQMenu(ui.menubar.QWidget)
	ui.menuHelp.SetObjectName(*qt.NewQAnyStringView3("menuHelp"))
	ui.menuHelp.QWidget.AddAction(ui.actionAbout)
	ui.menuHelp.QWidget.AddAction(ui.actionAbout_Qt)
	ui.menubar.AddMenu(ui.menuFile)
	ui.menubar.AddMenu(ui.menuHelp)
	ui.MainWindow.SetMenuBar(ui.menubar)

	ui.toolBar = qt.NewQToolBar(ui.MainWindow.QWidget)
	ui.toolBar.SetObjectName(*qt.NewQAnyStringView3("toolBar"))
	ui.MainWindow.AddToolBar(qt.TopToolBarArea, ui.toolBar)
	/* miqt-uic: no handler for toolBar attribute 'toolBarBreak' */

	ui.statusbar = qt.NewQStatusBar(ui.MainWindow.QWidget)
	ui.statusbar.SetObjectName(*qt.NewQAnyStringView3("statusbar"))
	ui.MainWindow.SetStatusBar(ui.statusbar)

	ui.Retranslate()

	return ui
}

// Retranslate reapplies all text translations.
func (ui *MainWindowUi) Retranslate() {
	ui.actionAbout.SetText(qt.QMainWindow_Tr("About"))
	ui.actionAbout_Qt.SetText(qt.QMainWindow_Tr("About Qt"))
	ui.actionExit.SetText(qt.QMainWindow_Tr("Exit"))
	ui.menuFile.SetTitle(qt.QMenuBar_Tr("File"))
	ui.menuHelp.SetTitle(qt.QMenuBar_Tr("Help"))
	ui.toolBar.SetWindowTitle(qt.QMainWindow_Tr("toolBar"))
}

