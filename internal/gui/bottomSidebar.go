package gui

import (
	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/internal/i18n"
)

func bottomSidebar() giu.Widget {
	return giu.TabBar().Flags(giu.TabBarFlagsReorderable).TabItems(
		giu.TabItem(i18n.L("AudioMixer")).Layout(audioMixer()),
		giu.TabItem(i18n.L("ProjectDescription")).Layout(projDescription()),
	)
}

func audioMixer() giu.Widget {
	return giu.Label("TODO: audio mixer")
}

func projDescription() giu.Widget {
	return giu.Label("TODO: project description")
}
