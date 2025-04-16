package gui

import (
	"github.com/AllenDang/giu"
)

func MakeWindowLoop() (func(), error) {
	aboutWnd, err := makeAboutWindowLoop()
	if err != nil {
		return nil, err
	}

	return func() {
		giu.SingleWindowWithMenuBar().Layout(
			menuBar(),
			giu.Label("TODO: Main content goes here."),
			statusBar(),
		)

		aboutWnd()
		convertVorbisWindowLoop()
	}, nil
}
