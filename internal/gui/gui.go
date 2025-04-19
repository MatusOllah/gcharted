package gui

import (
	"github.com/AllenDang/giu"
)

var sashPosRightSidebar float32 = 1000
var sashPosBottomSidebar float32 = 500

func MakeWindowLoop() (func(), error) {
	aboutWnd, err := makeAboutWindowLoop()
	if err != nil {
		return nil, err
	}

	return func() {
		giu.PushWindowPadding(0, 0)
		giu.SingleWindowWithMenuBar().Layout(giu.CSSTag("masterwindow").To(
			menuBar(),
			giu.Child().ID("container").Size(-1, -25).Border(false).Layout(
				giu.SplitLayout(giu.DirectionVertical, &sashPosRightSidebar,
					giu.SplitLayout(giu.DirectionHorizontal, &sashPosBottomSidebar,
						giu.Label("TODO: main content"),
						bottomSidebar(),
					),
					rightSidebar(),
				)),
			giu.Child().ID("statusbar").Size(-1, -1).Flags(giu.WindowFlagsNoDecoration).Layout(giu.CSSTag("statusbar").To(statusBar())),
		))
		giu.PopStyle()

		aboutWnd()
		convertVorbisWindowLoop()
	}, nil
}
