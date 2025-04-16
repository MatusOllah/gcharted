package gui

import (
	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/version"
)

func MakeWindowLoop() (func(), error) {
	aboutWnd, err := makeAboutWindowLoop()
	if err != nil {
		return nil, err
	}

	return func() {
		giu.SingleWindowWithMenuBar().Layout(
			menuBar(),
			giu.Label("Main content goes here."),
			giu.Custom(func() {
				// Push status bar to bottom
				w, h := giu.GetAvailableRegion()
				giu.Dummy(w, h-25).Build() // Spacer to push the footer down

				// Draw status bar
				giu.Separator().Build()
				giu.Row(
					giu.Label(version.Version),
				).Build()
			}),
		)

		aboutWnd()
		convertVorbisWindowLoop()
	}, nil
}
