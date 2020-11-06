package systems

import (
	"log"

	"game/pkg/components"
	"game/pkg/messages"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type SpeedSystem struct {
	entities    []components.SpeedEntity
	levelWidth  float32
	levelHeight float32
}

func (s *SpeedSystem) SetLevelArea(width, height float32) {
	s.levelHeight = height
	s.levelWidth = width
}

func (s *SpeedSystem) New(*ecs.World) {
	engo.Mailbox.Listen(messages.SPEED_MESSAGE, func(message engo.Message) {
		speed, isSpeed := message.(messages.SpeedMessage)
		if isSpeed {
			log.Printf("%#v\n", speed.Point)
			for _, e := range s.entities {
				if e.ID() == speed.BasicEntity.ID() {
					e.SpeedComponent.Point = speed.Point
				}
			}
		}
	})

	engo.Mailbox.Listen(common.CollisionMessage{}.Type(), func(message engo.Message) {
		_, isCollision := message.(common.CollisionMessage)
		if isCollision {
			// See if we also have that Entity, and if so, change the speed
			// for _, e := range s.entities {
			// 	if e.ID() == collision.Entity.BasicEntity.ID() {
			// 		e.SpeedComponent.X *= -1
			// 	}
			// }

			log.Printf("do u wanna a pice of me? \n")
		}
	})
}

func (s *SpeedSystem) Add(basic *ecs.BasicEntity, speed *components.SpeedComponent, space *common.SpaceComponent) {
	s.entities = append(s.entities, components.SpeedEntity{basic, speed, space})
}

func (s *SpeedSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range s.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *SpeedSystem) Update(dt float32) {

	for _, e := range s.entities {
		speed := engo.GameWidth() * dt
		e.SpaceComponent.Position.X = e.SpaceComponent.Position.X + speed*e.SpeedComponent.Point.X
		e.SpaceComponent.Position.Y = e.SpaceComponent.Position.Y + speed*e.SpeedComponent.Point.Y

		// Add Game Border Limits
		var heightLimit float32 = s.levelHeight - e.SpaceComponent.Height
		if e.SpaceComponent.Position.Y < 0 {
			e.SpaceComponent.Position.Y = 0
		} else if e.SpaceComponent.Position.Y > heightLimit {
			e.SpaceComponent.Position.Y = heightLimit
		}

		var widthLimit float32 = s.levelWidth - e.SpaceComponent.Width
		if e.SpaceComponent.Position.X < 0 {
			e.SpaceComponent.Position.X = 0
		} else if e.SpaceComponent.Position.X > widthLimit {
			e.SpaceComponent.Position.X = widthLimit
		}
	}

}
