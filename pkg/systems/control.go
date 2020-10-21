package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/Mat24/engo/pkg/game"
)

type ControlSystem struct {
	entities []ControlEntity
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, anim *common.AnimationComponent, control *game.ControlComponent, space *common.SpaceComponent) {
	c.entities = append(c.entities, ControlEntity{basic, anim, control, space})
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
		setAnimation(e)

		if vector, changed := getSpeed(e); changed {
			speed := dt * SPEED_SCALE
			vector, _ = vector.Normalize()
			vector.MultiplyScalar(speed)
			engo.Mailbox.Dispatch(SpeedMessage{e.BasicEntity, vector})
		}
	}
}
