package notegrid

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/MatusOllah/gcharted/assets"
	"github.com/MatusOllah/gcharted/internal/funkin"
)

type NoteGrid struct {
	widget.Table

	song         *funkin.Song
	sectionIndex int
}

func New() *NoteGrid {
	ng := new(NoteGrid)
	ng.ExtendBaseWidget(ng)

	ng.sectionIndex = 0

	ng.Length = func() (int, int) {
		return 0, 4
	}

	ng.CreateCell = func() fyne.CanvasObject {
		return widget.NewLabel("")
	}

	ng.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
	}

	return ng
}

func (ng *NoteGrid) SetSong(song *funkin.Song) {
	ng.song = song

	ng.Length = func() (int, int) {
		return ng.song.Notes[ng.sectionIndex].LengthInSteps, 4
	}

	ng.CreateCell = func() fyne.CanvasObject {
		image := canvas.NewImageFromResource(assets.ArrowNull)
		image.FillMode = canvas.ImageFillOriginal
		image.SetMinSize(fyne.NewSize(64, 64))

		return image
	}

	ng.UpdateCell = func(id widget.TableCellID, template fyne.CanvasObject) {
		template.(*canvas.Image).Resize(fyne.NewSize(64, 64))
	}
}

func (ng *NoteGrid) SetSection(index int) {
	ng.sectionIndex = index
}
