package assets

//go:generate go run ../scripts/rcc/rcc.go "-Input" "assets.qrc" "-Package" "assets"

import (
	_ "embed"

	qt "github.com/mappu/miqt/qt6"
)

//go:embed assets.rcc
var _resourceRcc []byte

func init() {
	qt.QResource_RegisterResourceWithRccData(&_resourceRcc[0])
}
