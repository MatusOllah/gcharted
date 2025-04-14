package gui

import (
	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/internal/i18n"
	"github.com/ncruces/zenity"
)

var showConvertVorbisWindow bool = false

var convertVorbisInputFilePath string
var convertVorbisOutputFilePath string

func convertVorbisWindowLoop() {
	if showConvertVorbisWindow {
		giu.Window(i18n.L("ConvertVorbis")).IsOpen(&showConvertVorbisWindow).Size(640, 480).Flags(giu.WindowFlagsNoResize).Layout(
			giu.Row(
				giu.Label(i18n.L("InputFile")+": "),
				FileLabel(&convertVorbisInputFilePath, FileLabelTypeOpen).FileFilters(zenity.FileFilters{zenity.FileFilter{Name: i18n.L("AnyFFmpegCompatible"), Patterns: []string{"*.*"}}}),
			),
			// TODO: ui
			// FFmpeg command: ffmpeg -i [input] -vn -codec:a libvorbis -qscale:a [quality (range: -1 - 10)] -ac 2 -ar 44100 [output]
		)
	}
}
