package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"math/rand"
	"picture-puzzle/global"
	"picture-puzzle/scenemanager"
	"picture-puzzle/scenes"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	scenemanager.SetScene(&scenes.StartScene{})
	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "시애틀 2030 Puzzle")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
