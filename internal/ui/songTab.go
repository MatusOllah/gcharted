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

	charList := []string{
		"bf",
		"dad",
		"gf",
		"spooky",
		"pico",
		"mom",
		"mom-car",
		"bf-car",
		"parents-christmas",
		"monster-christmas",
		"bf-christmas",
		"gf-christmas",
		"monster",
		"bf-pixel",
		"senpai",
		"senpai-angry",
		"spirit",
	}

	playerEntry := widget.NewSelectEntry(charList)
	playerEntry.SetPlaceHolder("Player")

	opponentEntry := widget.NewSelectEntry(charList)
	opponentEntry.SetPlaceHolder("Opponent")

	return container.NewTabItem("Song", container.NewVBox(
		container.NewBorder(nil, nil, widget.NewLabel("Name:"), nil, nameEntry),
		container.NewBorder(nil, nil, widget.NewLabel("BPM:"), nil, bpmEntry),
		container.NewBorder(nil, nil, widget.NewLabel("Speed:"), nil, speedEntry),
		container.NewBorder(nil, nil, widget.NewLabel("Player:"), nil, playerEntry),
		container.NewBorder(nil, nil, widget.NewLabel("Opponent:"), nil, opponentEntry),
		hasVoiceTrackCheckbox,
	))
}
