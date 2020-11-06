package components

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type PauseEntity struct {
	*ecs.BasicEntity
	*common.AnimationComponent
	*common.SpaceComponent
	*common.RenderComponent
	*ControlComponent
	*SpeedComponent
}
