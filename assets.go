package main

//go:generate go run scripts/rcc/rcc.go "-Input" "assets.qrc"

import (
	_ "embed"

	qt "github.com/mappu/miqt/qt6"
)

//go:embed assets.rcc
var _resourceRcc []byte

func init() {
	qt.QResource_RegisterResourceWithRccData(&_resourceRcc[0])
}
