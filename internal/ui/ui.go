package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"
)

func MakeUI(w fyne.Window) fyne.CanvasObject {
	sectionTabItem := container.NewTabItem("Section", widget.NewLabel("horalky"))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Name")

	hasVoiceTrackCheckbox := widget.NewCheck("Has voice track", func(b bool) {
		if b {
			log.Info().Msgf("checked has voice track checkbox")
		} else if !b {
			log.Info().Msgf("unchecked has voice track checkbox")
		}
	})

	bpmEntry := NewNumEntry()
	bpmEntry.SetPlaceHolder("BPM")

	speedEntry := NewNumEntry()
	speedEntry.SetPlaceHolder("Speed")

	playerEntry := widget.NewEntry()
	playerEntry.SetPlaceHolder("Player")

	opponentEntry := widget.NewEntry()
	opponentEntry.SetPlaceHolder("Opponent")

	songTabItem := container.NewTabItem("Song", container.NewVBox(
		nameEntry,
		hasVoiceTrackCheckbox,
		bpmEntry,
		speedEntry,
		playerEntry,
		opponentEntry,
	))

	masterMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped master mute button")
	})
	masterMuteButton.SetIcon(theme.VolumeMuteIcon())

	masterVolumeSlider := widget.NewSlider(-100, 0)

	masterVolumeLabel := widget.NewLabel("0")

	masterTrackCard := widget.NewCard("Master", "", container.NewBorder(nil, nil, masterMuteButton, masterVolumeLabel, masterVolumeSlider))

	instMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped inst mute button")
	})
	instMuteButton.SetIcon(theme.VolumeMuteIcon())

	instVolumeSlider := widget.NewSlider(-100, 0)

	instVolumeLabel := widget.NewLabel("0")

	instTrackCard := widget.NewCard("Instrumental", "", container.NewBorder(nil, nil, instMuteButton, instVolumeLabel, instVolumeSlider))

	vocalsMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped vocals mute button")
	})
	vocalsMuteButton.SetIcon(theme.VolumeMuteIcon())

	vocalsVolumeSlider := widget.NewSlider(-100, 0)

	vocalsVolumeLabel := widget.NewLabel("0")

	vocalsTrackCard := widget.NewCard("Vocals", "", container.NewBorder(nil, widget.NewLabel("path/to/vocals.ogg"), vocalsMuteButton, vocalsVolumeLabel, vocalsVolumeSlider))

	tracksTabItem := container.NewTabItem("Tracks", container.NewVBox(
		masterTrackCard,
		instTrackCard,
		vocalsTrackCard,
	))

	rightAppTabs := container.NewAppTabs(
		sectionTabItem,
		songTabItem,
		tracksTabItem,
	)

	return container.NewBorder(
		makeToolbar(w),
		makeStatusBar(),
		nil,
		nil,
		container.NewHSplit(container.NewDocTabs(), rightAppTabs),
	)
}
