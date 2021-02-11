package characters

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type DecorationItem struct {
	ecs.BasicEntity
	common.CollisionComponent
	common.SpaceComponent
}
