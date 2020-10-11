package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/uTa3/Breakout/breakout"
)

func main() {
	game, err := breakout.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(breakout.ScreenWidth*2, breakout.ScreenHeight*2)
	ebiten.SetWindowTitle("Breakout")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}