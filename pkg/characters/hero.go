package characters

import (
	"game/pkg/components"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type Hero struct {
	ecs.BasicEntity
	common.CollisionComponent
	common.AnimationComponent
	common.RenderComponent
	common.SpaceComponent
	components.ControlComponent
	components.SpeedComponent
}
