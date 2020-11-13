package characters

import (
	"game/pkg/components"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type Enemy struct {
	ecs.BasicEntity
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent
	components.ZControlComponent
}
