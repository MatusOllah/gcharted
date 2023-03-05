package gcharted

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
)

var (
	Version = "1.0.0"

	PositionBinding     binding.String = binding.NewString()
	IsPaused            bool
	IsInstMuted         bool           = false
	IsVocalsMuted       bool           = false
	InstVolumeBinding   binding.Float  = binding.NewFloat()
	VocalsVolumeBinding binding.Float  = binding.NewFloat()
	InstPathBinding     binding.String = binding.NewString()
	VocalsPathBinding   binding.String = binding.NewString()
	InstStreamer        beep.StreamSeeker
	VocalsStreamer      beep.StreamSeeker
	InstCtrl            *beep.Ctrl
	VocalsCtrl          *beep.Ctrl
	InstVolume          *effects.Volume
	VocalsVolume        *effects.Volume
	InstFormat          beep.Format
)
