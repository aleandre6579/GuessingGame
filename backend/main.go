package main

import (
	"context"
	"guessing-game/app"
)

func main() {
	game := app.Init(app.CreateConfig())
	game.Start(context.Background())
}
