package systems

import (
	"game/pkg/components"
	"game/pkg/messages"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

type ZControlSystem struct {
	entities []components.ZControlComponent
}

func (c *ZControlSystem) New(*ecs.World) {
	engo.Mailbox.Listen(messages.POSITION_HERO_MESSAGE, func(message engo.Message) {
		msg, isSpeed := message.(messages.PositionHeroMessage)
		if isSpeed {
			log.Printf("ZControl %#v\n", msg.Point)
			for _, e := range c.entities {
				e.Renderer.SetZIndex(msg.ZPosition - 1)
				if e.Space.Position.Y > msg.Point.Y {
					e.Renderer.SetZIndex(msg.ZPosition + 1)
				}
			}
		}
	})
}

//Add _
func (c *ZControlSystem) Add(control *components.ZControlComponent) {
	c.entities = append(c.entities, *control)
}

//Remove _
func (c *ZControlSystem) Remove(basic ecs.BasicEntity) {
}

//Update _
func (c *ZControlSystem) Update(dt float32) {
}
