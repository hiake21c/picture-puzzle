package main

import (
	"log"
	"picture-puzzle/global"
	"picture-puzzle/scenemanager"
	"picture-puzzle/scenes"

	"github.com/hajimehoshi/ebiten"
)

func main() {

	scenemanager.SetScene(&scenes.StartScene{})
	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "Jongmin Kim's Puzzle")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
