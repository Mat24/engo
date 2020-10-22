package main

import (
	"github.com/EngoEngine/engo"
	"github.com/Mat24/engo/pkg/game"
)

/*
1. Refactor, show some order!
2. Explore collision using engo
3. Explore background sound
4. HID on mouse click
5. Explore animmation
6. Enemy?
7. Handle live, game over?
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
