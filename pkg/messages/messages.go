package messages

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

const (
	SPEED_MESSAGE         = "SpeedMessage"
	POSITION_HERO_MESSAGE = "PositionHero"
	SPEED_SCALE           = 16
)

type SpeedMessage struct {
	*ecs.BasicEntity
	engo.Point
}

func (SpeedMessage) Type() string {
	return SPEED_MESSAGE
}

type PositionHeroMessage struct {
	engo.Point
	ZPosition float32
}

func (PositionHeroMessage) Type() string {
	return POSITION_HERO_MESSAGE
}
