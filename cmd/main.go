package main

import (
	"game/pkg/game"

	"github.com/EngoEngine/engo"
)

/*
[1]1. Refactor, show some order!
[0]2. Explore collision using engo
[0]3. Explore background sound
[0]4. HID on mouse click
[0]5. Explore animmation
[0]6. Enemy?
[0]7. Handle live, game over?
*/

func main() {
	opts := engo.RunOptions{
		Title:      "My Little Adventure",
		AssetsRoot: "../assets",
		Width:      500,
		Height:     500,
	}
	engo.Run(opts, &game.DefaultScene{})
}
