package main

import (
	"context"
	"guessing-game/app"
)

func main() {
	appInstance := app.Init(app.CreateConfig())
	appInstance.Start(context.Background())
}
