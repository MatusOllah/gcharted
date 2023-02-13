package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ztrue/tracerr"

	"github.com/MatusOllah/gcharted/internal/funkin"
	"github.com/MatusOllah/gcharted/internal/ui"
)

const (
	version = "1.0.0"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}).With().Caller().Logger()

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}

	log.Info().Msgf("gcharted version %s", version)
	log.Info().Msgf("Go version %s", runtime.Version())
	log.Info().Msg("ahoj!")

	a := app.New()

	a.Lifecycle().SetOnStarted(func() {
		log.Info().Msg("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Info().Msg("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Info().Msg("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Info().Msg("Lifecycle: Exited Foreground")
	})

	w := a.NewWindow("gcharted")
	w.SetMaster()
	w.Resize(fyne.NewSize(1280, 720))
	w.SetMainMenu(makeMenu(a, w))
	w.SetContent(container.NewBorder(makeToolbar(w), makeStatusBar(), nil, nil, makeUI()))

	w.ShowAndRun()

	log.Info().Msg("exiting")
	runtime.GC()
	os.Exit(0)
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	chartItem := fyne.NewMenuItem("Chart", func() {
		log.Info().Msg("selected menu item Open>Chart")
	})
	chartItem.Icon = fnfArrowIcon

	instItem := fyne.NewMenuItem("Instrumental", func() {
		log.Info().Msg("selected menu item Open>Audio>Instrumental")
	})

	vocalsItem := fyne.NewMenuItem("Vocals", func() {
		log.Info().Msg("selected menu item Open>Audio>Vocals")
	})

	audioItem := fyne.NewMenuItem("Audio", nil)
	audioItem.Icon = theme.VolumeUpIcon()
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
		fyne.NewMenuItem("New", func() {
			log.Info().Msg("selected menu item File>New")
		}),
		fyne.NewMenuItemSeparator(),
		openItem,
	)

	cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		log.Info().Msg("selected menu item Edit>Cut")
		shortcutFocused(cutShortcut, w)
	})
	cutItem.Shortcut = cutShortcut

	copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	copyItem := fyne.NewMenuItem("Copy", func() {
		log.Info().Msg("selected menu item Edit>Copy")
		shortcutFocused(copyShortcut, w)
	})
	copyItem.Shortcut = copyShortcut

	pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	pasteItem := fyne.NewMenuItem("Paste", func() {
		log.Info().Msg("selected menu item Edit>Paste")
		shortcutFocused(pasteShortcut, w)
	})
	pasteItem.Shortcut = pasteShortcut

	edit := fyne.NewMenu("Edit",
		cutItem,
		copyItem,
		pasteItem,
	)

	help := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			log.Info().Msg("selected menu item Help>About")
			dialog.NewInformation("About gcharted", fmt.Sprintf(
				"gcharted version %s\nGo version %s",
				version,
				runtime.Version(),
			), w).Show()
		}),
	)

	main := fyne.NewMainMenu(
		file,
		edit,
		help,
	)
	return main
}

func makeToolbar(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), func() {
			log.Info().Msg("selected toolbar item New")
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			log.Info().Msg("selected toolbar item Open")
			dialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {

			}, w)

			dialog.Show()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			log.Info().Msg("selected toolbar item Save")
		}),
	), widget.NewSeparator())
}

func makeStatusBar() fyne.CanvasObject {
	return container.NewVBox(widget.NewSeparator(), container.NewHBox(
		widget.NewLabel(version),
		widget.NewSeparator(),
		widget.NewLabel("Position: 0 / 0 s"),
		widget.NewSeparator(),
		widget.NewLabel("Section: 0"),
		widget.NewSeparator(),
		widget.NewLabel("CurStep: 0"),
	))
}

func makeUI() fyne.CanvasObject {
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

	bpmEntry := ui.NewNumEntry()
	bpmEntry.SetPlaceHolder("BPM")

	speedEntry := ui.NewNumEntry()
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

	masterTrackCard := widget.NewCard("Master", "Volume", container.NewBorder(nil, nil, masterMuteButton, masterVolumeLabel, masterVolumeSlider))

	instMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped inst mute button")
	})
	instMuteButton.SetIcon(theme.VolumeMuteIcon())

	instVolumeSlider := widget.NewSlider(-100, 0)

	instVolumeLabel := widget.NewLabel("0")

	instTrackCard := widget.NewCard("Instrumental", "Volume", container.NewBorder(nil, nil, instMuteButton, instVolumeLabel, instVolumeSlider))

	vocalsMuteButton := widget.NewButton("", func() {
		log.Info().Msg("tapped vocals mute button")
	})
	vocalsMuteButton.SetIcon(theme.VolumeMuteIcon())

	vocalsVolumeSlider := widget.NewSlider(-100, 0)

	vocalsVolumeLabel := widget.NewLabel("0")

	vocalsTrackCard := widget.NewCard("Vocals", "Volume", container.NewBorder(nil, nil, vocalsMuteButton, vocalsVolumeLabel, vocalsVolumeSlider))

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

	return container.NewHSplit(container.NewDocTabs(), rightAppTabs)
}

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}

func openFile(uri fyne.URI) error {
	content, err := os.ReadFile(uri.Path())
	if err != nil {
		return tracerr.Wrap(err)
	}

	song, err := funkin.LoadSongFromJSON(content)
	if err != nil {
		return tracerr.Wrap(err)
	}

	log.Info().Msgf("%v", song)

	return nil
}
