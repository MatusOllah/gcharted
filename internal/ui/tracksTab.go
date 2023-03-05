package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MatusOllah/gcharted/internal/gcharted"
	"github.com/rs/zerolog/log"
)

func makeTracksTab() *container.TabItem {
	instMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped inst mute button")

		gcharted.IsInstMuted = !gcharted.IsInstMuted
		log.Info().Bool("IsInstMuted", gcharted.IsInstMuted).Msg("")

		gcharted.SetInstMuted(gcharted.IsInstMuted)
	})
	instMuteButton.SetIcon(theme.VolumeMuteIcon())

	instVolumeSlider := widget.NewSlider(-100, 0)
	instVolumeSlider.Bind(gcharted.InstVolumeBinding)

	instVolumeEntry := NewNumEntry()
	instVolumeEntry.Bind(binding.FloatToStringWithFormat(gcharted.InstVolumeBinding, "%.2f"))

	instTrackCard := widget.NewCard("Instrumental", "", container.NewBorder(
		nil,
		widget.NewLabelWithData(gcharted.InstPathBinding),
		instMuteButton,
		instVolumeEntry,
		instVolumeSlider,
	))

	vocalsMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped vocals mute button")

		gcharted.IsVocalsMuted = !gcharted.IsVocalsMuted
		log.Info().Bool("IsVocalsMuted", gcharted.IsVocalsMuted).Msg("")

		gcharted.SetVocalsMuted(gcharted.IsVocalsMuted)
	})
	vocalsMuteButton.SetIcon(theme.VolumeMuteIcon())

	vocalsVolumeSlider := widget.NewSlider(-100, 0)
	vocalsVolumeSlider.Bind(gcharted.VocalsVolumeBinding)

	vocalsVolumeEntry := NewNumEntry()
	vocalsVolumeEntry.Bind(binding.FloatToStringWithFormat(gcharted.VocalsVolumeBinding, "%.2f"))

	vocalsTrackCard := widget.NewCard("Vocals", "", container.NewBorder(
		nil,
		widget.NewLabelWithData(gcharted.VocalsPathBinding),
		vocalsMuteButton,
		vocalsVolumeEntry,
		vocalsVolumeSlider,
	))

	return container.NewTabItem("Tracks", container.NewVBox(
		instTrackCard,
		vocalsTrackCard,
	))
}
