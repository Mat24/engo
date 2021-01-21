package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

//BackgroundAudioSystem represents the audio system that will play sounds in the bg
type BackgroundAudioSystem struct {
	ecs.BasicEntity
	common.AudioComponent
	playList    []string
	player      *common.Player
	playerIndex int
}

//NewBackgroundAudioSystem creates a new instance of *BackgroundAudioSystem
func NewBackgroundAudioSystem(audio ...string) *BackgroundAudioSystem {
	return &BackgroundAudioSystem{
		AudioComponent: common.AudioComponent{
			Player: &common.Player{},
		},
		playList: append(audio),
	}
}

//Add _
func (w *BackgroundAudioSystem) Add(audio *common.AudioComponent) {

}

//Remove _
func (w *BackgroundAudioSystem) Remove(basic ecs.BasicEntity) {

}

func (w *BackgroundAudioSystem) setNextPlayItem() error {
	if w.playerIndex >= len(w.playList) {
		w.playerIndex = 0
	}
	var err error
	w.player, err = common.LoadedPlayer(w.playList[w.playerIndex])
	w.playerIndex++
	if err != nil {
		return err
	}
	w.AudioComponent.Player = w.player
	return nil
}

//Update _
func (w *BackgroundAudioSystem) Update(dt float32) {
	if w.player != nil && w.player.IsPlaying() {
		return
	}

	if err := w.setNextPlayItem(); err != nil {
		log.Fatal(err)
		return
	}

	w.player.Rewind()
	w.player.Play()
	w.player.SetVolume(0.8)
}
