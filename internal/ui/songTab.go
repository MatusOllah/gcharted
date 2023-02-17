package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"
)

func makeSongTab() *container.TabItem {
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

	return container.NewTabItem("Song", container.NewVBox(
		nameEntry,
		hasVoiceTrackCheckbox,
		bpmEntry,
		speedEntry,
		playerEntry,
		opponentEntry,
	))
}
