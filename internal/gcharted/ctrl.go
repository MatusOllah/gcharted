package gcharted

import (
	"fmt"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/rs/zerolog/log"
)

func UpdatePosition() {
	PositionBinding.Set(fmt.Sprintf("Position: %v", InstFormat.SampleRate.D(InstStreamer.Position()).Round(time.Millisecond)))
}

func Play() {
	if InstStreamer == nil {
		log.Warn().Msg("audio playback not initialized, InstStreamer is nil")
	} else {
		log.Info().Msg("Initializing audio playback")
		speaker.Init(InstFormat.SampleRate, InstFormat.SampleRate.N(time.Second/10))
	}

	if InstStreamer == nil {
		log.Warn().Msg("skipping audio playback, InstStreamer is nil")
	} else {
		log.Info().Msg("Starting audio playback")
		if VocalsStreamer == nil {
			speaker.Play(InstVolume)
		} else {
			speaker.Play(beep.Mix(InstVolume, VocalsVolume))
		}
	}
}

func SetPaused(paused bool) {
	speaker.Lock()
	InstCtrl.Paused = paused
	VocalsCtrl.Paused = paused
	speaker.Unlock()
}

func Rewind() error {
	log.Info().Msg("rewinding")
	speaker.Lock()
	if err := InstStreamer.Seek(0); err != nil {
		return err
	}
	if err := VocalsStreamer.Seek(0); err != nil {
		return err
	}

	log.Info().Int("InstStreamer.Len()", InstStreamer.Len()).Msg("")
	log.Info().Int("VocalsStreamer.Len()", VocalsStreamer.Len()).Msg("")
	speaker.Unlock()

	return nil
}

func Forward() error {
	log.Info().Msg("forwarding")
	speaker.Lock()
	if err := InstStreamer.Seek(InstStreamer.Len()); err != nil {
		return err
	}
	if err := VocalsStreamer.Seek(VocalsStreamer.Len()); err != nil {
		return err
	}

	log.Info().Int("InstStreamer.Len()", InstStreamer.Len()).Msg("")
	log.Info().Int("VocalsStreamer.Len()", VocalsStreamer.Len()).Msg("")
	speaker.Unlock()

	return nil
}
