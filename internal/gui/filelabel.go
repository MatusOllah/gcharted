package gui

import (
	"log/slog"

	"github.com/AllenDang/giu"
	"github.com/MatusOllah/gcharted/internal/i18n"
	"github.com/ncruces/zenity"
)

type FileLabelType int

const (
	FileLabelTypeOpen FileLabelType = iota
	FileLabelTypeSave
)

type FileLabelWidget struct {
	id          giu.ID
	path        *string
	typ         FileLabelType
	dialogTitle string
	showHidden  bool
	width       float32
	filters     zenity.FileFilters
}

func FileLabel(path *string, typ FileLabelType) *FileLabelWidget {
	return &FileLabelWidget{
		id:   giu.GenAutoID("file-label"),
		path: path,
		typ:  typ,
	}
}

func (w *FileLabelWidget) Size(width float32) *FileLabelWidget {
	w.width = width
	return w
}

func (w *FileLabelWidget) DialogTitle(s string) *FileLabelWidget {
	w.dialogTitle = s
	return w
}

func (w *FileLabelWidget) ShowHidden(show bool) *FileLabelWidget {
	w.showHidden = show
	return w
}

func (w *FileLabelWidget) FileFilters(filters zenity.FileFilters) *FileLabelWidget {
	w.filters = filters
	return w
}

func (w *FileLabelWidget) ID(id giu.ID) *FileLabelWidget {
	w.id = id
	return w
}

func (w *FileLabelWidget) Build() {
	giu.Row(
		giu.InputText(w.path).Size(w.width-100),
		giu.Button(i18n.L("Browse")).Size(100, 0).OnClick(func() {
			slog.Debug("[FileLabelWidget] clicked browse button", "id", w.id, "typ", w.typ)

			opts := []zenity.Option{zenity.Title(w.dialogTitle), zenity.Filename(*w.path), w.filters}
			if w.showHidden {
				opts = append(opts, zenity.ShowHidden())
			}

			switch w.typ {
			case FileLabelTypeOpen:
				path, _ := zenity.SelectFile(opts...)
				if path != "" {
					*w.path = path
				}
			case FileLabelTypeSave:
				opts = append(opts, zenity.ConfirmOverwrite())
				path, _ := zenity.SelectFileSave(opts...)
				if path != "" {
					*w.path = path
				}
			}
		}),
	).Build()
}
