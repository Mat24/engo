package main

import (
	"github.com/EngoEngine/engo"
	"github.com/Mat24/engo/pkg/game"
)

func main() {
	opts := engo.RunOptions{
		Title:      "My Little Adventure",
		AssetsRoot: "../assets",
		Width:      500,
		Height:     500,
	}
	engo.Run(opts, &game.DefaultScene{})
}
