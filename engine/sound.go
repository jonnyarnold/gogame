package engine

import (
	"github.com/banthar/Go-SDL/mixer"
	"path"
)

// SoundPath is the folder
// that sounds will be pulled from.
var SoundPath string

type sound struct {
	music   mixer.Music
	playing bool
}

func LoadSound(soundName string) *sound {
	soundPath := path.Join(SoundPath, soundName)
	mus := mixer.LoadMUS(soundPath)
	localSound := sound{music: *mus}
	return &localSound
}

func (s sound) Play() {
	if mixer.PausedMusic() == 1 {
		mixer.ResumeMusic()
	}

	if s.playing == false {
		s.music.PlayMusic(1)
		s.playing = true
	}
}

func PauseAllMusic() {
	if mixer.PlayingMusic() == 1 {
		mixer.PauseMusic()
	}
}
