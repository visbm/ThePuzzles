package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"thegame/game"
	"thegame/logger"
)

const (
	screenWidth  uint = 550
	screenHeight uint = 550

	rows = 2
	cols = 2
)

func main() {
	log := logger.GetLogger()
	defer recoverPanic(log)

	//img, err := image.LoadImage("bmw.png", screenWidth, screenHeight)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//ctx := game.NewContext(img, rows, cols, screenWidth, screenHeight)

	theGame := game.NewGame(log,
		rows, cols,
		screenWidth, screenHeight,
	)

	err := theGame.Init()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))

	ebiten.SetWindowTitle("TheGame")
	if err := ebiten.RunGame(theGame); err != nil {
		log.Fatal(err)
	}
}

func recoverPanic(log logger.Logger) {
	if r := recover(); r != nil {
		log.Panicf("panic: %v", r)
	}
}
