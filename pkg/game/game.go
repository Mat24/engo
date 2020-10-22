package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

var (
	model  = "motw.png"
	width  = 52
	height = 73
)

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}
