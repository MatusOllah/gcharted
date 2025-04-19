package gui

import "github.com/AllenDang/giu"

var showBottomSiderbar bool = true

func bottomSidebar() giu.Widget {
	if !showBottomSiderbar {
		return nil
	}

	return giu.Label("TODO: bottom sidebar")
}
