package systems

import (
	"game/pkg/components"
	"game/pkg/messages"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type ControlSystem struct {
	entities []components.ControlEntity
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, anim *common.AnimationComponent, control *components.ControlComponent, space *common.SpaceComponent) {
	c.entities = append(c.entities, components.ControlEntity{basic, anim, control, space})
}

func (c *ControlSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

func (c *ControlSystem) Update(dt float32) {
	for _, e := range c.entities {
		components.SetAnimation(e)

		if vector, changed := components.GetSpeed(e); changed {
			speed := dt * messages.SPEED_SCALE
			vector, _ = vector.Normalize()
			vector.MultiplyScalar(speed)
			engo.Mailbox.Dispatch(messages.SpeedMessage{e.BasicEntity, vector})
			engo.Mailbox.Dispatch(messages.PositionHeroMessage{
				e.Position,
				10,
			})
		}
	}
}
