package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed assets/icons/fnf-arrow.svg
var fnfArrowSvg []byte
var fnfArrowIcon = &fyne.StaticResource{
	StaticName:    "fnf-arrow.svg",
	StaticContent: fnfArrowSvg,
}
