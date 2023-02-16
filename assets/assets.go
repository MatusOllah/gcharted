package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed icons/fnf-arrow.svg
var fnfArrowSvg []byte
var FNFArrowIcon = &fyne.StaticResource{
	StaticName:    "fnf-arrow.svg",
	StaticContent: fnfArrowSvg,
}
