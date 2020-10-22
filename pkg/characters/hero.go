package characters

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/Mat24/engo/pkg/components"
)

type Hero struct {
	ecs.BasicEntity
	common.AnimationComponent
	common.RenderComponent
	common.SpaceComponent
	components.ControlComponent
	components.SpeedComponent
}
