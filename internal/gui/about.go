package gui

import (
	"image"
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

func makeAboutWindowLoop() (func(), error) {
	iconImg, err := getIcon()
	if err != nil {
		return nil, err
	}

	return func() {
		if showAboutWindow {
			giu.Window(i18n.L("About")).IsOpen(&showAboutWindow).Flags(giu.WindowFlagsNoResize | giu.WindowFlagsNoDocking).Layout(giu.CSSTag("window").To(
				giu.Align(giu.AlignCenter).To(giu.ImageWithRgba(iconImg).Size(64, 64)),
				giu.Label(i18n.LT("GChartedVersion", map[string]any{"Version": version.Version})),
				giu.Label(i18n.LT("GoVersion", map[string]any{"Version": runtime.Version(), "GOOS": runtime.GOOS, "GOARCH": runtime.GOARCH})),
				giu.Label(""),
				giu.Label(i18n.L("MadeWithLove")),
				giu.Label("Copyright (c) 2025 Matúš Ollah"),
				giu.Label(i18n.L("License")),
				giu.Link("https://github.com/MatusOllah/gcharted").OnClick(func() {
					giu.OpenURL("https://github.com/MatusOllah/gcharted")
				}),
			))
		}
	}, nil
}
