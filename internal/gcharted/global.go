package gcharted

import "fyne.io/fyne/v2/data/binding"

var (
	Version = "1.0.0"

	IsMasterMuted       bool = false
	IsInstMuted         bool = false
	IsVocalsMuted       bool = false
	MasterVolumeBinding binding.Float
	InstVolumeBinding   binding.Float
	VocalsVolumeBinding binding.Float
	InstPathBinding     binding.String
	VocalsPathBinding   binding.String
)
