package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/MatusOllah/gcharted/internal/gcharted"
	"github.com/rs/zerolog/log"
)

func makeTracksTab() *container.TabItem {
	masterMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped master mute button")
		gcharted.IsMasterMuted = !gcharted.IsMasterMuted
		log.Info().Bool("IsMasterMuted", gcharted.IsMasterMuted).Msg("")
	})
	masterMuteButton.SetIcon(theme.VolumeMuteIcon())

	masterVolumeSlider := widget.NewSlider(-100, 0)

	masterVolumeLabel := widget.NewLabel("0")

	masterTrackCard := widget.NewCard("Master", "", container.NewBorder(nil, nil, masterMuteButton, masterVolumeLabel, masterVolumeSlider))

	instMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped inst mute button")
		gcharted.IsInstMuted = !gcharted.IsInstMuted
		log.Info().Bool("IsInstMuted", gcharted.IsInstMuted).Msg("")
	})
	instMuteButton.SetIcon(theme.VolumeMuteIcon())

	instVolumeSlider := widget.NewSlider(-100, 0)

	instVolumeLabel := widget.NewLabel("0")

	instTrackCard := widget.NewCard("Instrumental", "", container.NewBorder(nil, nil, instMuteButton, instVolumeLabel, instVolumeSlider))

	vocalsMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped vocals mute button")
		gcharted.IsVocalsMuted = !gcharted.IsVocalsMuted
		log.Info().Bool("IsVocalsMuted", gcharted.IsVocalsMuted).Msg("")
	})
	vocalsMuteButton.SetIcon(theme.VolumeMuteIcon())

	vocalsVolumeSlider := widget.NewSlider(-100, 0)

	vocalsVolumeLabel := widget.NewLabel("0")

	vocalsTrackCard := widget.NewCard("Vocals", "", container.NewBorder(nil, widget.NewLabel("path/to/vocals.ogg"), vocalsMuteButton, vocalsVolumeLabel, vocalsVolumeSlider))

	return container.NewTabItem("Tracks", container.NewVBox(
		masterTrackCard,
		instTrackCard,
		vocalsTrackCard,
	))
}
