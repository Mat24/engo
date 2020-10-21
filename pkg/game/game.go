package game

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

var (
	WalkUpAction    *common.Animation
	WalkDownAction  *common.Animation
	WalkLeftAction  *common.Animation
	WalkRightAction *common.Animation
	StopUpAction    *common.Animation
	StopDownAction  *common.Animation
	StopLeftAction  *common.Animation
	StopRightAction *common.Animation
	SkillAction     *common.Animation
	actions         []*common.Animation

	upButton    = "up"
	downButton  = "down"
	leftButton  = "left"
	rightButton = "right"
	pauseButton = "pause"
	model       = "motw.png"
	width       = 52
	height      = 73
	levelWidth  float32
	levelHeight float32
)

const (
	SPEED_MESSAGE = "SpeedMessage"
	SPEED_SCALE   = 64
)

type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}

type SpeedMessage struct {
	*ecs.BasicEntity
	engo.Point
}

func (SpeedMessage) Type() string {
	return SPEED_MESSAGE
}

func setAnimation(e ControlEntity) {
	if engo.Input.Button(upButton).JustPressed() {
		e.AnimationComponent.SelectAnimationByAction(WalkUpAction)
	} else if engo.Input.Button(downButton).JustPressed() {
		e.AnimationComponent.SelectAnimationByAction(WalkDownAction)
	} else if engo.Input.Button(leftButton).JustPressed() {
		e.AnimationComponent.SelectAnimationByAction(WalkLeftAction)
	} else if engo.Input.Button(rightButton).JustPressed() {
		e.AnimationComponent.SelectAnimationByAction(WalkRightAction)
	}

	if engo.Input.Button(upButton).JustReleased() {
		e.AnimationComponent.SelectAnimationByAction(StopUpAction)
		if engo.Input.Button(leftButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkLeftAction)
		} else if engo.Input.Button(rightButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkRightAction)
		} else if engo.Input.Button(upButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkUpAction)
		} else if engo.Input.Button(downButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkDownAction)
		}
	} else if engo.Input.Button(downButton).JustReleased() {
		e.AnimationComponent.SelectAnimationByAction(StopDownAction)
		if engo.Input.Button(leftButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkLeftAction)
		} else if engo.Input.Button(rightButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkRightAction)
		} else if engo.Input.Button(upButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkUpAction)
		} else if engo.Input.Button(downButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkDownAction)
		}
	} else if engo.Input.Button(leftButton).JustReleased() {
		e.AnimationComponent.SelectAnimationByAction(StopLeftAction)
		if engo.Input.Button(leftButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkLeftAction)
		} else if engo.Input.Button(rightButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkRightAction)
		} else if engo.Input.Button(upButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkUpAction)
		} else if engo.Input.Button(downButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkDownAction)
		}
	} else if engo.Input.Button(rightButton).JustReleased() {
		e.AnimationComponent.SelectAnimationByAction(StopRightAction)
		if engo.Input.Button(leftButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkLeftAction)
		} else if engo.Input.Button(rightButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkRightAction)
		} else if engo.Input.Button(upButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkUpAction)
		} else if engo.Input.Button(downButton).Down() {
			e.AnimationComponent.SelectAnimationByAction(WalkDownAction)
		}
	}
}

func GetSpeed(e ControlEntity) (p engo.Point, changed bool) {
	p.X = engo.Input.Axis(e.ControlComponent.SchemeHoriz).Value()
	p.Y = engo.Input.Axis(e.ControlComponent.SchemeVert).Value()
	origX, origY := p.X, p.Y

	if engo.Input.Button(upButton).JustPressed() {
		p.Y = -1
	} else if engo.Input.Button(downButton).JustPressed() {
		p.Y = 1
	}
	if engo.Input.Button(leftButton).JustPressed() {
		p.X = -1
	} else if engo.Input.Button(rightButton).JustPressed() {
		p.X = 1
	}

	if engo.Input.Button(upButton).JustReleased() || engo.Input.Button(downButton).JustReleased() {
		p.Y = 0
		changed = true
		if engo.Input.Button(upButton).Down() {
			p.Y = -1
		} else if engo.Input.Button(downButton).Down() {
			p.Y = 1
		} else if engo.Input.Button(leftButton).Down() {
			p.X = -1
		} else if engo.Input.Button(rightButton).Down() {
			p.X = 1
		}
	}
	if engo.Input.Button(leftButton).JustReleased() || engo.Input.Button(rightButton).JustReleased() {
		p.X = 0
		changed = true
		if engo.Input.Button(leftButton).Down() {
			p.X = -1
		} else if engo.Input.Button(rightButton).Down() {
			p.X = 1
		} else if engo.Input.Button(upButton).Down() {
			p.Y = -1
		} else if engo.Input.Button(downButton).Down() {
			p.Y = 1
		}
	}
	changed = changed || p.X != origX || p.Y != origY
	return
}
