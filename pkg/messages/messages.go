package messages

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

const (
	SPEED_MESSAGE = "SpeedMessage"
	SPEED_SCALE   = 64
)

type SpeedMessage struct {
	*ecs.BasicEntity
	engo.Point
}

func (SpeedMessage) Type() string {
	return SPEED_MESSAGE
}
