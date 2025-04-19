package gui

import (
	"github.com/AllenDang/giu"
)

var sashPosRightSidebar float32 = 1000
var sashPosBottomSidebar float32 = 500

var showRightSiderbar bool = true
var showBottomSiderbar bool = true

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
				giu.Custom(func() {
					main := giu.Label("TODO: main content")

					var widget giu.Widget = main

					if showBottomSiderbar {
						widget = giu.SplitLayout(giu.DirectionHorizontal, &sashPosBottomSidebar, main, bottomSidebar())
					}
					if showRightSiderbar {
						widget = giu.SplitLayout(giu.DirectionVertical, &sashPosRightSidebar, widget, rightSidebar())
					}

					widget.Build()
				}),
			),
			giu.Child().ID("statusbar").Size(-1, -1).Flags(giu.WindowFlagsNoDecoration).Layout(giu.CSSTag("statusbar").To(statusBar())),
		))
		giu.PopStyle()

		aboutWnd()
		convertVorbisWindowLoop()
	}, nil
}
