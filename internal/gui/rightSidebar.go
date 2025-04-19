package gui

import "github.com/AllenDang/giu"

var showRightSiderbar bool = true

func rightSidebar() giu.Widget {
	if !showRightSiderbar {
		return nil
	}

	return giu.Label("TODO: right sidebar")
}
