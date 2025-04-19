package gui

import (
	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/version"
)

func statusBar() giu.Widget {
	return giu.Row(
		giu.Label(version.Version),
	)
}
