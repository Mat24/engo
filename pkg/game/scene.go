package game

import (
	"image/color"
	"log"

	"game/pkg/characters"
	"game/pkg/components"
	"game/pkg/systems"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

var (
	actions []*common.Animation
)

type DefaultScene struct{}

func (*DefaultScene) Preload() {

	// Load character model
	engo.Files.Load(model)

	// Load TileMap
	if err := engo.Files.Load("example.tmx"); err != nil {
		panic(err)
	}

	err := engo.Files.Load("326488.wav")
	if err != nil {
		log.Println(err)
	}
	// err = engo.Files.Load("forest.mp3")
	// if err != nil {
	// 	log.Println(err)
	// }
	err = engo.Files.Load("TownTheme.mp3")
	if err != nil {
		log.Println(err)
	}

	components.StopUpAction = &common.Animation{
		Name:   "upstop",
		Frames: []int{37},
	}

	components.StopDownAction = &common.Animation{
		Name:   "downstop",
		Frames: []int{1},
	}

	components.StopLeftAction = &common.Animation{
		Name:   "leftstop",
		Frames: []int{13},
	}

	components.StopRightAction = &common.Animation{
		Name:   "rightstop",
		Frames: []int{25},
	}

	components.WalkUpAction = &common.Animation{
		Name:   "up",
		Frames: []int{36, 37, 38},
		Loop:   true,
	}

	components.WalkDownAction = &common.Animation{
		Name:   "down",
		Frames: []int{0, 1, 2},
		Loop:   true,
	}

	components.WalkLeftAction = &common.Animation{
		Name:   "left",
		Frames: []int{12, 13, 14},
		Loop:   true,
	}

	components.WalkRightAction = &common.Animation{
		Name:   "right",
		Frames: []int{24, 25, 26},
		Loop:   true,
	}

	actions = []*common.Animation{
		components.StopUpAction,
		components.StopDownAction,
		components.StopLeftAction,
		components.StopRightAction,
		components.WalkUpAction,
		components.WalkDownAction,
		components.WalkLeftAction,
		components.WalkRightAction,
	}

	engo.Input.RegisterButton(components.UpButton, engo.KeyW, engo.KeyArrowUp)
	engo.Input.RegisterButton(components.LeftButton, engo.KeyA, engo.KeyArrowLeft)
	engo.Input.RegisterButton(components.RightButton, engo.KeyD, engo.KeyArrowRight)
	engo.Input.RegisterButton(components.DownButton, engo.KeyS, engo.KeyArrowDown)
	engo.Input.RegisterButton(components.PauseButton, engo.KeySpace)
}

func (scene *DefaultScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	common.SetBackground(color.White)

	speedSystem := &systems.SpeedSystem{}
	controlSystem := &systems.ControlSystem{}
	audioSystem := systems.NewBackgroundAudioSystem("TownTheme.mp3")

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.AnimationSystem{})
	w.AddSystem(speedSystem)
	w.AddSystem(controlSystem)
	w.AddSystem(&systems.PauseSystem{})
	w.AddSystem(audioSystem)
	w.AddSystem(&common.AudioSystem{})

	w.AddSystem(&common.CollisionSystem{Solids: 1})
	w.AddSystem(&systems.ZControlSystem{})

	// Setup TileMap
	resource, err := engo.Files.Resource("example.tmx")
	if err != nil {
		panic(err)
	}
	tmxResource := resource.(common.TMXResource)
	levelData := tmxResource.Level

	// Extract Map Size
	levelWidth := levelData.Bounds().Max.X
	levelHeight := levelData.Bounds().Max.Y

	speedSystem.SetLevelArea(levelWidth, levelHeight)

	// Create Hero
	spriteSheet := common.NewSpritesheetFromFile(model, heroWidth, heroHeight)

	hero := scene.CreateHero(
		engo.Point{engo.GameWidth() / 2, engo.GameHeight() / 2},
		spriteSheet,
	)
	hero.ControlComponent = components.ControlComponent{
		SchemeHoriz: "horizontal",
		SchemeVert:  "vertical",
	}

	hero.RenderComponent.SetZIndex(10)
	hero.CollisionComponent = common.CollisionComponent{
		Main: 1,
	}

	enemy := scene.CreateEnemy(
		engo.Point{engo.GameWidth()/2 + 120, engo.GameHeight()/2 + 120},
		spriteSheet,
	)
	enemy.CollisionComponent = common.CollisionComponent{
		Group: 1,
	}

	enemy.RenderComponent.SetZIndex(9)

	// Add our hero to the appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.RenderComponent,
				&hero.SpaceComponent)

			sys.Add(&enemy.BasicEntity,
				&enemy.RenderComponent,
				&enemy.SpaceComponent)

		case *common.AnimationSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.AnimationComponent,
				&hero.RenderComponent)

		case *systems.ControlSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.AnimationComponent,
				&hero.ControlComponent,
				&hero.SpaceComponent)

		case *systems.SpeedSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.SpeedComponent,
				&hero.SpaceComponent)

		case *systems.PauseSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.AnimationComponent,
				&hero.SpaceComponent,
				&hero.RenderComponent,
				&hero.ControlComponent,
				&hero.SpeedComponent)

		case *common.CollisionSystem:
			sys.Add(&hero.BasicEntity,
				&hero.CollisionComponent,
				&hero.SpaceComponent)
			sys.Add(&enemy.BasicEntity,
				&enemy.CollisionComponent,
				&enemy.SpaceComponent)

		case *systems.ZControlSystem:
			sys.Add(&components.ZControlComponent{
				Space:    &enemy.SpaceComponent,
				Renderer: &enemy.RenderComponent,
			})
		case *common.AudioSystem:
			sys.Add(&audioSystem.BasicEntity, &audioSystem.AudioComponent)
		}

	}

	// Create render and space components for each of the tiles
	tileComponents := make([]*Tile, 0)

	for _, tileLayer := range levelData.TileLayers {
		for _, tileElement := range tileLayer.Tiles {

			if tileElement.Image != nil {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: tileElement,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: tileElement.Point,
					Width:    0,
					Height:   0,
				}

				if tileLayer.Name == "grass" {
					tile.RenderComponent.SetZIndex(0)
				}

				if tileLayer.Name == "trees" {
					tile.RenderComponent.SetZIndex(2)
				}

				tileComponents = append(tileComponents, tile)
			}
		}
	}

	for _, imageLayer := range levelData.ImageLayers {
		for _, imageElement := range imageLayer.Images {

			if imageElement.Image != nil {
				tile := &Tile{BasicEntity: ecs.NewBasic()}
				tile.RenderComponent = common.RenderComponent{
					Drawable: imageElement,
					Scale:    engo.Point{1, 1},
				}
				tile.SpaceComponent = common.SpaceComponent{
					Position: imageElement.Point,
					Width:    0,
					Height:   0,
				}

				if imageLayer.Name == "clouds" {
					tile.RenderComponent.SetZIndex(3)
				}

				tileComponents = append(tileComponents, tile)
			}
		}
	}

	// Add each of the tiles entities and its components to the render system
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, v := range tileComponents {
				sys.Add(&v.BasicEntity, &v.RenderComponent, &v.SpaceComponent)
			}
		case *systems.PauseSystem:
			for _, v := range tileComponents {
				sys.Add(
					&v.BasicEntity,
					nil,
					&v.SpaceComponent,
					&v.RenderComponent,
					nil,
					nil,
				)
			}
		}
	}

	// Access Object Layers
	for _, objectLayer := range levelData.ObjectLayers {
		log.Println("This object layer is called " + objectLayer.Name)
		// Do something with every regular Object
		for _, object := range objectLayer.Objects {
			log.Println("This object is called " + object.Name)
		}
	}

	// Setup character and movement
	engo.Input.RegisterAxis(
		"vertical",
		engo.AxisKeyPair{engo.KeyArrowUp, engo.KeyArrowDown},
		engo.AxisKeyPair{engo.KeyW, engo.KeyS},
	)

	engo.Input.RegisterAxis(
		"horizontal",
		engo.AxisKeyPair{engo.KeyArrowLeft, engo.KeyArrowRight},
		engo.AxisKeyPair{engo.KeyA, engo.KeyD},
	)

	// Add EntityScroller System
	w.AddSystem(&common.EntityScroller{
		SpaceComponent: &hero.SpaceComponent,
		TrackingBounds: levelData.Bounds(),
	})
}

func (*DefaultScene) Type() string { return "DefaultScene" }

func (*DefaultScene) CreateHero(point engo.Point, spriteSheet *common.Spritesheet) *characters.Hero {
	hero := &characters.Hero{BasicEntity: ecs.NewBasic()}

	hero.SpaceComponent = common.SpaceComponent{
		Position: point,
		Width:    float32(heroWidth) / 2,
		Height:   float32(heroHeight / 2),
	}
	hero.RenderComponent = common.RenderComponent{
		Drawable: spriteSheet.Cell(0),
		Scale:    engo.Point{1, 1},
	}

	hero.SpeedComponent = components.SpeedComponent{}
	hero.AnimationComponent = common.NewAnimationComponent(spriteSheet.Drawables(), 0.1)

	hero.AnimationComponent.AddAnimations(actions)
	hero.AnimationComponent.SelectAnimationByName("downstop")

	return hero
}

func (*DefaultScene) CreateEnemy(point engo.Point, spriteSheet *common.Spritesheet) *characters.Enemy {
	enemy := &characters.Enemy{BasicEntity: ecs.NewBasic()}

	enemy.SpaceComponent = common.SpaceComponent{
		Position: point,
		Width:    float32(heroWidth / 2),
		Height:   float32(heroHeight / 2),
	}
	enemy.RenderComponent = common.RenderComponent{
		Drawable: spriteSheet.Cell(10),
		Scale:    engo.Point{1, 1},
	}

	return enemy
}
