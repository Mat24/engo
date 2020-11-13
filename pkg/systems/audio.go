package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type BackgroundAudioSystem struct {
	ecs.BasicEntity
	common.AudioComponent
	playList    []string
	player      *common.Player
	playerIndex int
}

func NewBackgroundAudioSystem(audio ...string) *BackgroundAudioSystem {
	return &BackgroundAudioSystem{playList: append(audio)}
}

func (w *BackgroundAudioSystem) Add(audio *common.AudioComponent) {

}

func (w *BackgroundAudioSystem) Remove(basic ecs.BasicEntity) {

}

func (w *BackgroundAudioSystem) Update(dt float32) {
	if w.player == nil {
		if w.playerIndex >= len(w.playList) {
			w.playerIndex = 0
		}
		var err error
		w.player, err = common.LoadedPlayer(w.playList[w.playerIndex])
		w.playerIndex++
		if err != nil {
			log.Fatalln(err)
		}
		w.AudioComponent.Player = w.player
		w.player.Play()
		return
	}
	// if w.player.IsPlaying() == false {
	// 	w.player = nil
	// }
}
