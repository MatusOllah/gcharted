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

/*
type fileLabelWidgetState struct {
	path string
}

func (s *fileLabelWidgetState) Dispose() {}

func (w *FileLabelWidget) getState() *fileLabelWidgetState {
	if s := giu.GetState[fileLabelWidgetState](giu.Context, w.stateID()); s != nil {
		return s
	}

	newState := w.newState()
	giu.Context.SetState(w.stateID(), newState)

	return w.getState()
}

func (w *FileLabelWidget) newState() *fileLabelWidgetState {
	return &fileLabelWidgetState{path: ""}
}

func (w *FileLabelWidget) stateID() giu.ID {
	return w.id
}
*/

type FileLabelWidget struct {
	id          giu.ID
	path        *string
	typ         FileLabelType
	dialogTitle string
	showHidden  bool
	filters     zenity.FileFilters
}

func FileLabel(path *string, typ FileLabelType) *FileLabelWidget {
	return &FileLabelWidget{
		id:   giu.GenAutoID("file-label"),
		path: path,
		typ:  typ,
	}
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
		giu.InputText(w.path),
		giu.Button(i18n.L("Browse")).OnClick(func() {
			slog.Debug("[FileLabelWidget] clicked browse button", "id", w.id, "typ", w.typ)
			switch w.typ {
			case FileLabelTypeOpen:
				opts := []zenity.Option{zenity.Title(w.dialogTitle), zenity.Filename(*w.path), w.filters}
				if w.showHidden {
					opts = append(opts, zenity.ShowHidden())
				}
				path, _ := zenity.SelectFile(opts...)
				if path != "" {
					*w.path = path
				}
			}
		}),
	).Build()
}
