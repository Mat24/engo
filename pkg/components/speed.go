package components

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type SpeedEntity struct {
	*ecs.BasicEntity
	*SpeedComponent
	*common.SpaceComponent
}

type SpeedComponent struct {
	engo.Point
}
