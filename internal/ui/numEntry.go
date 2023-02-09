package ui

import (
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type NumEntry struct {
	widget.Entry
}

func (n *NumEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func NewNumEntry() *NumEntry {
	e := &NumEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`\d`, "Must contain a number")
	return e
}
