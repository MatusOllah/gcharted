package gui

import (
	"log/slog"
	"os"

	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/internal/i18n"
)

func menuBar() *giu.MenuBarWidget {
	return giu.MenuBar().Layout(
		giu.Menu(i18n.L("File")).Layout(
			giu.MenuItem(i18n.L("Exit")).OnClick(func() {
				slog.Debug("clicked exit menu item, exiting")
				os.Exit(0)
			}),
		),
		giu.Menu(i18n.L("Tools")).Layout(
			giu.MenuItem(i18n.L("ConvertVorbis")).OnClick(func() {
				slog.Debug("clicked convert to vorbis menu item")
				showConvertVorbisWindow = true
			}),
		),
		giu.Menu(i18n.L("Help")).Layout(
			giu.MenuItem(i18n.L("About")).OnClick(func() {
				slog.Debug("clicked about menu item")
				showAboutWindow = true
			}),
		),
	)
}
