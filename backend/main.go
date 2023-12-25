package main

import (
	"context"
	"guessing-game/app"
	"guessing-game/config"
)

func main() {
	game := app.Init(config.CreateConfig())
	game.Start(context.Background())
}
