package gui

import (
	"log/slog"

	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/internal/i18n"
	"github.com/ncruces/zenity"
)

var showConvertVorbisWindow bool = false

var convertVorbisInputFilePath string
var convertVorbisOutputFilePath string
var convertVorbisQuality int32 = 8

func convertVorbisWindowLoop() {
	if showConvertVorbisWindow {
		giu.Window(i18n.L("ConvertVorbis")).IsOpen(&showConvertVorbisWindow).Size(640, 480).Flags(giu.WindowFlagsNoResize).Layout(
			giu.Row(
				giu.Label(i18n.L("InputFile")),
				FileLabel(&convertVorbisInputFilePath, FileLabelTypeOpen).Size(400).FileFilters(zenity.FileFilters{zenity.FileFilter{Name: i18n.L("AnyFFmpegCompatible"), Patterns: []string{"*.*"}}}),
			),
			giu.Row(
				giu.Label(i18n.L("Quality")),
				giu.SliderInt(&convertVorbisQuality, -1, 10).Size(400),
			),
			giu.Row(
				giu.Label(i18n.L("OutputFile")),
				FileLabel(&convertVorbisOutputFilePath, FileLabelTypeSave).Size(400).FileFilters(zenity.FileFilters{zenity.FileFilter{Name: "Ogg Vorbis", Patterns: []string{"*.ogg"}}}),
			),
			giu.Row(
				giu.Button(i18n.L("Cancel")).OnClick(func() {
					slog.Debug("[convertVorbis] clicked cancel button")
					showConvertVorbisWindow = false // hide window
				}),
				giu.Button(i18n.L("Convert")).OnClick(func() {
					slog.Debug("[convertVorbis] clicked convert button")
					slog.Warn("TODO: convert vorbis using FFmpeg", "input", convertVorbisInputFilePath, "quality", convertVorbisQuality, "output", convertVorbisOutputFilePath)
					showConvertVorbisWindow = false // hide window
				}),
			),
			// FFmpeg command: ffmpeg -i [input] -vn -codec:a libvorbis -qscale:a [quality (range: -1 - 10)] -ac 2 -ar 44100 [output]
		)
	}
}
