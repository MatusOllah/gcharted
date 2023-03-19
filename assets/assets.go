package assets

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed icons/fnf-arrow.svg
var fnfArrowSVG []byte
var FNFArrowIcon = &fyne.StaticResource{
	StaticName:    "fnf-arrow.svg",
	StaticContent: fnfArrowSVG,
}

//go:embed icons/arrows/null.png
var arrowNullPNG []byte
var ArrowNull = &fyne.StaticResource{
	StaticName:    "null.png",
	StaticContent: arrowNullPNG,
}

//go:embed icons/arrows/left.png
var arrowLeftPNG []byte
var ArrowLeft = &fyne.StaticResource{
	StaticName:    "left.png",
	StaticContent: arrowLeftPNG,
}

//go:embed icons/arrows/down.png
var arrowDownPNG []byte
var ArrowDown = &fyne.StaticResource{
	StaticName:    "down.png",
	StaticContent: arrowDownPNG,
}

//go:embed icons/arrows/up.png
var arrowUpPNG []byte
var ArrowUp = &fyne.StaticResource{
	StaticName:    "up.png",
	StaticContent: arrowUpPNG,
}

//go:embed icons/arrows/right.png
var arrowRightPNG []byte
var ArrowRight = &fyne.StaticResource{
	StaticName:    "right.png",
	StaticContent: arrowRightPNG,
}
