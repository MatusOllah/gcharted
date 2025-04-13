package gui

import (
	"image"
	"log/slog"
	"os"
	"runtime"

	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/assets"
	"github.com/MatusOllah/gcharted/internal/i18n"
	"github.com/MatusOllah/gcharted/version"
)

var showAboutWindow bool = false

func getIcon() (*image.RGBA, error) {
	f, err := assets.FS.Open("icon.png")
	if err != nil {
		return nil, err
	}

	return giu.PNGToRgba(f)
}

func MakeWindowLoop() (func(), error) {
	iconImg, err := getIcon()
	if err != nil {
		return nil, err
	}

	return func() {
		giu.MainMenuBar().Layout(
			giu.Menu(i18n.L("File")).Layout(
				giu.MenuItem(i18n.L("Exit")).OnClick(func() {
					slog.Info("clicked exit button, exiting")
					os.Exit(0)
				}),
			),
			giu.Menu(i18n.L("Help")).Layout(
				giu.MenuItem(i18n.L("About")).OnClick(func() {
					slog.Debug("clicked about button")
					showAboutWindow = true
				}),
			),
		).Build()

		if showAboutWindow {
			giu.Window(i18n.L("About")).IsOpen(&showAboutWindow).Size(320, 240).Flags(giu.WindowFlagsNoResize).Layout(
				giu.Align(giu.AlignCenter).To(giu.ImageWithRgba(iconImg).Size(64, 64)),
				giu.Label(i18n.LT("GChartedVersion", map[string]any{"Version": version.Version})),
				giu.Label(i18n.LT("GoVersion", map[string]any{"Version": runtime.Version(), "GOOS": runtime.GOOS, "GOARCH": runtime.GOARCH})),
				giu.Label(""),
				giu.Label(i18n.L("MadeWithLove")),
				giu.Label("Copyright (c) 2025 Matúš Ollah"),
				giu.Label(i18n.L("License")),
			)
		}
	}, nil
}
