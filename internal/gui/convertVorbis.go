package gui

import (
	"context"
	"log/slog"

	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/internal/i18n"
	"github.com/ncruces/zenity"
	"github.com/xfrr/goffmpeg/v2/ffmpeg"
	"github.com/xfrr/goffmpeg/v2/ffprobe"
)

var showConvertVorbisWindow bool
var showConvertVorbisProgressWindow bool

var convertVorbisInputFilePath string
var convertVorbisOutputFilePath string
var convertVorbisQuality int32 = 8

var convertVorbisProgress float32

func convertVorbisWindowLoop() {
	if showConvertVorbisWindow {
		giu.Window(i18n.L("ConvertVorbis")).IsOpen(&showConvertVorbisWindow).Flags(giu.WindowFlagsNoResize|giu.WindowFlagsNoDocking).Layout(
			giu.Row(
				giu.Label(i18n.L("InputFile")),
				FileLabel(&convertVorbisInputFilePath, FileLabelTypeOpen).Size(400).FileFilters(zenity.FileFilters{zenity.FileFilter{Name: i18n.L("AnyFFmpegCompatible"), Patterns: []string{"*.*"}}}),
			),
			giu.Row(
				giu.Label(i18n.L("Quality")),
				giu.SliderInt(&convertVorbisQuality, 0, 10).Size(400),
			),
			giu.Row(
				giu.Label(i18n.L("OutputFile")),
				FileLabel(&convertVorbisOutputFilePath, FileLabelTypeSave).Size(400).FileFilters(zenity.FileFilters{zenity.FileFilter{Name: "Ogg Vorbis", Patterns: []string{"*.ogg"}}}),
			),
			giu.Label(""),
			giu.Row(
				giu.Button(i18n.L("Cancel")).OnClick(func() {
					slog.Debug("[convertVorbis] clicked cancel button")
					showConvertVorbisWindow = false // hide window
				}),
				giu.Button(i18n.L("Convert")).OnClick(func() {
					slog.Debug("[convertVorbis] clicked convert button")
					go convertVorbis(convertVorbisInputFilePath, convertVorbisOutputFilePath, int(convertVorbisQuality))
				}),
			),
		)
	}

	if showConvertVorbisProgressWindow {
		giu.Window(i18n.L("Converting")).IsOpen(&showConvertVorbisProgressWindow).Size(300, 100).Flags(giu.WindowFlagsNoResize | giu.WindowFlagsNoNav | giu.WindowFlagsNoDocking).Layout(
			giu.Column(
				giu.Label(i18n.L("Converting")),
				giu.ProgressBar(convertVorbisProgress).Overlayf("%.2f%%", convertVorbisProgress*100),
			),
		)
	}
}

func convertVorbis(input, output string, quality int) {
	if input == "" || output == "" {
		return
	}

	slog.Info("[convertVorbis] converting to ogg vorbis", "input", input, "output", output, "quality", quality)
	showConvertVorbisProgressWindow = true
	convertVorbisProgress = 0

	ctx, cancel := context.WithCancel(context.Background())

	// FFmpeg command: ffmpeg -i [input] -vn -dn -codec:a libvorbis -qscale:a [quality (range: -1 - 10)] -ac 2 -ar 44100 [output]
	cmd := ffmpeg.NewCommand().WithInputPath(input).WithDisableVideo().WithDisableData().WithAudioCodec("libvorbis").WithAudioQuality(quality).WithAudioChannels(2).WithAudioRate(44100).WithOutputFormat("ogg").WithOutputPath(output)

	mediafile, err := ffprobe.NewCommand().WithInputPath(input).Run(ctx)
	if err != nil {
		zenity.Error("Failed to determine length: " + err.Error())
		return
	}

	slog.Debug("[convertVorbis] running FFmpeg command", "args", cmd.Args())
	progress, err := cmd.Start(ctx)
	if err != nil {
		zenity.Error("Failed to run FFmpeg: " + err.Error())
		return
	}

	go func() {
		for msg := range progress {
			slog.Debug("[convertVorbis] converting", "msg", msg)
			convertVorbisProgress = float32(msg.Duration().Milliseconds()) / float32(mediafile.Duration().Milliseconds())
			giu.Update()
		}
		cancel()
		showConvertVorbisProgressWindow = false // hide window
	}()
}
