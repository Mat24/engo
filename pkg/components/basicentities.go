package components

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type SpeedComponent struct {
	engo.Point
}

type ControlComponent struct {
	SchemeVert  string
	SchemeHoriz string
}

type ControlEntity struct {
	*ecs.BasicEntity
	*common.AnimationComponent
	*ControlComponent
	*common.SpaceComponent
}

type SpeedEntity struct {
	*ecs.BasicEntity
	*SpeedComponent
	*common.SpaceComponent
}

type PauseEntity struct {
	*ecs.BasicEntity
	*common.AnimationComponent
	*common.SpaceComponent
	*common.RenderComponent
	*ControlComponent
	*SpeedComponent
}
