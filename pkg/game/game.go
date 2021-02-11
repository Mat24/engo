package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

var (
	model      = "motw.png"
	heroWidth  = 52
	heroHeight = 73
)

type Tile struct {
	Name string
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}
