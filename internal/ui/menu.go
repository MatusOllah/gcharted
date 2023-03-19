package ui

import (
	"fmt"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"github.com/rs/zerolog/log"

	"github.com/MatusOllah/gcharted/assets"
	"github.com/MatusOllah/gcharted/internal/gcharted"
)

func MakeMenu(w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", func() {
		log.Info().Msg("selected menu item File>New")
	})
	newItem.Icon = theme.FolderNewIcon()

	chartItem := fyne.NewMenuItem("Chart", func() {
		log.Info().Msg("selected menu item Open>Chart")
		gcharted.OpenChart(w)
	})
	chartItem.Icon = assets.FNFArrowIcon

	instItem := fyne.NewMenuItem("Instrumental", func() {
		log.Info().Msg("selected menu item Open>Audio>Instrumental")
		gcharted.OpenInst(w)
	})

	vocalsItem := fyne.NewMenuItem("Vocals", func() {
		log.Info().Msg("selected menu item Open>Audio>Vocals")
		gcharted.OpenVocals(w)
	})

	audioItem := fyne.NewMenuItem("Audio", nil)
	audioItem.Icon = theme.MediaMusicIcon()
	audioItem.ChildMenu = fyne.NewMenu("",
		instItem,
		vocalsItem,
	)

	openItem := fyne.NewMenuItem("Open", nil)
	openItem.Icon = theme.FolderOpenIcon()
	openItem.ChildMenu = fyne.NewMenu("",
		chartItem,
		audioItem,
	)

	file := fyne.NewMenu(
		"File",
		newItem,
		fyne.NewMenuItemSeparator(),
		openItem,
	)

	help := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			log.Info().Msg("selected menu item Help>About")
			dialog.NewInformation("About gcharted", fmt.Sprintf(
				"gcharted version %s\nGo version %s",
				gcharted.Version,
				runtime.Version(),
			), w).Show()
		}),
	)

	main := fyne.NewMainMenu(
		file,
		help,
	)
	return main
}
