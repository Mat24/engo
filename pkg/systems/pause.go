package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type PauseSystem struct {
	entities []pauseEntity
	world    *ecs.World
	paused   bool
}

func (p *PauseSystem) New(w *ecs.World) {
	p.world = w
}

func (p *PauseSystem) Add(basic *ecs.BasicEntity, animation *common.AnimationComponent, space *common.SpaceComponent, render *common.RenderComponent, control *ControlComponent, speed *SpeedComponent) {
	p.entities = append(p.entities, pauseEntity{basic, animation, space, render, control, speed})
}

func (p *PauseSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range p.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		p.entities = append(p.entities[:delete], p.entities[delete+1:]...)
	}
}

func (p *PauseSystem) Update(dt float32) {
	if engo.Input.Button(pauseButton).JustPressed() {
		if !p.paused {
			for _, system := range p.world.Systems() {
				switch sys := system.(type) {
				case *common.AnimationSystem:
					for _, ent := range p.entities {
						sys.Remove(*ent.BasicEntity)
					}
				case *SpeedSystem:
					for _, ent := range p.entities {
						sys.Remove(*ent.BasicEntity)
					}
				case *ControlSystem:
					for _, ent := range p.entities {
						sys.Remove(*ent.BasicEntity)
					}
				}
			}
		} else {
			for _, system := range p.world.Systems() {
				switch sys := system.(type) {
				case *common.AnimationSystem:
					for _, ent := range p.entities {
						if ent.AnimationComponent != nil {
							sys.Add(
								ent.BasicEntity,
								ent.AnimationComponent,
								ent.RenderComponent,
							)
						}
					}
				case *SpeedSystem:
					for _, ent := range p.entities {
						if ent.SpeedComponent != nil {
							sys.Add(
								ent.BasicEntity,
								ent.SpeedComponent,
								ent.SpaceComponent,
							)
						}
					}
				case *ControlSystem:
					for _, ent := range p.entities {
						if ent.ControlComponent != nil {
							sys.Add(
								ent.BasicEntity,
								ent.AnimationComponent,
								ent.ControlComponent,
								ent.SpaceComponent,
							)
						}
					}
				}
			}
		}
		p.paused = !p.paused
	}
}
